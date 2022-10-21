# Go Gin vulnerability Example 

An example of gin contains many useful vul

一个go写的WEB漏洞靶场，实际自己写一下，加固一下知识

GIN框架 整个web框架是go-gin-Example 上面改的，没有前端框架，只有一个swagger，直接发包吧






## 0x01 sqli

实际中最常见的一种编码问题 Order by 之后存在列和表的的时候，一般采用拼接的情况出现sql注入

由于表/列名无法使用参数化查询，所以推荐使用白名单或转义操作
````
db.Order(xxxx).First(&user)
````
// 对于列名的修复，稳妥的是白名单
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
```

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
```
