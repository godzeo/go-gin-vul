# Go Gin vulnerability Example

An example of gin contains many useful vul

一个go写的WEB漏洞靶场，实际自己写一下，加固一下知识

GIN框架 整个web框架是go-gin-Example 上面改的，没有前端框架，只有一个swagger，直接发包吧




# 0x0 Vulnerability code analysis and fix 漏洞代码解析和修复

run
```
go mod tidy
go run main.go
```
conf/app.ini
```
[database]
Type = mysql
User = root
Password = 123456
Host = 127.0.0.1:3066
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
```


## 0x01 sqli

实际中最常见的一种编码问题 Order by 之后存在列和表的的时候，一般采用拼接的情况出现sql注入

由于表/列名无法使用参数化查询，所以推荐使用白名单或转义操作



### 0x011 常见错误拼接

主要是运用 fmt.Sprintf()、buffer.WriteString()等方式将字符串连接到一起。

简单就是先拼接，后查询都有问题

```
db.Select(xxx).First(&user) 

db.Where(fmt.Sprintf("name = '%s'", xxx)).Find(&user) 

db.Raw("select name from " + xxx).First(&user) 

db.Exec("select name from " + xxx).First(&user) 
```

### 0x012 业务中常见一定要拼接的地方

对于开发者来讲，SQL注入的修复主要有两种场景：

1. 常规value的拼接;
2. 表/列名的拼接。

原因可以看之前的文章，简单来说就是如果预编译会导致列名失效



实际中最常见的一种编码问题 Order by 之后存在列和表的的时候，一般采用拼接的情况出现sql注入

由于表/列名无法使用参数化查询，所以推荐使用白名单或转义操作

下面展开有代码


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
Host: 127.0.0.1:8080
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
Host: 127.0.0.1:8080
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

补充，最简单的拼接，用来测试waf
代码部分
```
	var qLogindata Logindata
	rawSQL := fmt.Sprintf("SELECT * FROM blog_auth WHERE id = '%s'", userID)

	if err := db.Raw(rawSQL).Scan(&qLogindata).Error; err != nil {
		return Sqliuserdata{RawSQL: rawSQL}, err
	}
```

````
POST /api/vul/sqli/byid HTTP/1.1
Host: 127.0.0.1:8080
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/109.0
Content-Type: application/x-www-form-urlencoded
Content-Length: 50

userid=1' UNION ALL SELECT 1111,22222,version()-- 
````
````
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 04 Mar 2023 02:24:06 GMT
Content-Length: 147

{"Logindata":{"ID":1111,"user":"","password":"8.0.32"},"RawSQL":"SELECT * FROM blog_auth WHERE id = '1' UNION ALL SELECT 1111,22222,version()-- '"}
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
Host: 127.0.0.1:8080
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
Host: 127.0.0.1:8080
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
修复：
建议直接写死，白名单



## 0x03 SSRF

SSRF攻击通常会导致组织内未经授权的操作或对数据的访问，
在某些情况下，SSRF漏洞可能允许攻击者执行任意命令执行。



常见函数
```go
http.Get(url)
http.Post(url, contentType, body)
http.Head(url)
http.PostForm(url, data)
http.NewRequest(method, url, body)
```
Full echo SSRF
routers/api/unAuth/ssrf.go
```go
        url := c.PostForm("q")
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("get image failed")
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	c.JSON(200, gin.H{
		"success": string(body),
	})
```

```go
POST /api/vul/ssrf HTTP/1.1
Host: 127.0.0.1:8080
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
Content-Length: 33

q=http://www.badiu.com/robots.txt
```
简单白修复
```go
    func GetImageSafe(c *gin.Context) {
	q := c.PostForm("q")
	url := "https://test.image.com/path/?q="
	res, err := http.Get(url + q)
	if err != nil {
		fmt.Println(err)
		fmt.Println("get image failed")
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	c.JSON(200, gin.H{
		"success": string(body),
	})
}
```


```go
[GIN] 2022/10/22 - 22:44:28 | 500 |   946.54142ms |       127.0.0.1 | POST     /api/safe/ssrf
Get "https://test.image.com/path/?q=@http://www.badiu.com/": x509: certificate is not valid for any names, but wanted to match test.image.com
get image failed

```
当然这个修复过于简单了，后面看一下成熟的修复方案










# 0x04 File Operation 文件操作

## 0x041 path traversing 路径穿越

在执行文件操作时，如果对从外部传入的文件名没有限制，则可能导致任意文件读取或任意文件写入，这可能严重导致代码执行。



### 0x0411 arbitrary file read 任意文件读



routers/api/unAuth/path.go

```go
func FileRead(c *gin.Context) {
	path := c.Query("filename")

	// Unfiltered file paths
	data, _ := ioutil.ReadFile(path)

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})

}

func Dirfile(c *gin.Context) {
	path := c.Query("filename")
	data, _ := ioutil.ReadFile(filepath.Join("/Users/zy", path))

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})

}
```



```
GET /api/vul/read?filename=/../../../../../../../../../../etc/passwd HTTP/1.1
Host: 127.0.0.1:8080
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

GET /api/vul/dir?filename=/../../etc/passwd HTTP/1.1
Host: 127.0.0.1:8080
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


```





### 0x0412 arbitrary file write 任意文件写

routers/api/unAuth/path.go

```go
func Unzip(c *gin.Context) {
	path := c.Query("filename")
	text := c.Query("text")
	file_path := filepath.Join("/Users/zy/", path)
	r, _ := zip.OpenReader(file_path)

	var abspath string
	for _, f := range r.File {
		abspath, _ = filepath.Abs(f.Name)
		ioutil.WriteFile(abspath, []byte(text), 0640)
	}

	data, _ := ioutil.ReadFile(abspath)

	c.JSON(200, gin.H{
		"success": "read: " + string(data),
	})
}
```



```raw
GET /api/vul/unzip?filename=radconfig.zip&text=Zeo666 HTTP/1.1
Host: 127.0.0.1:8080
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


```



### 0x0413 arbitrary file remove 任意文件删除

```go
//arbitrary file remove
func Fileremove(c *gin.Context) {
 	path := c.Query("path")
	os.Remove(path)
}
```



### 0x0414 FIX 修复

routers/api/unAuth/path.go

过滤 `..`

```go
func Unzipsafe(c *gin.Context) {
	path := c.Query("filename")
	file_path := filepath.Join("/Users/zy/", path)
	r, err := zip.OpenReader(file_path)
	if err != nil {
		fmt.Println("read zip file fail")
		c.JSON(500, gin.H{
			"success": "err: " + err.Error(),
		})
	}
	for _, f := range r.File {
		if !strings.Contains(f.Name, "..") {
			p, _ := filepath.Abs(f.Name)
			ioutil.WriteFile(p, []byte("present"), 0640)
		} else {
			c.JSON(500, gin.H{
				"success": "err: " + err.Error(),
			})
		}
	}
	c.JSON(200, gin.H{
		"success": "OK",
	})
}
```



```
GET /api/safe/unzip?filename=../../radconfig.zip&text=Zeo666 HTTP/1.1
Host: 127.0.0.1:8080
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

HTTP/1.1 500 Internal Server Error
Content-Type: application/json; charset=utf-8
Date: Sat, 05 Nov 2022 08:04:40 GMT
Content-Length: 65
Connection: close

{"success":"err: open /radconfig.zip: no such file or directory"}

```




## 0x042 File access permissions 文件权限

根据创建文件的敏感度设置不同级别的访问权限，以防止具有任意权限的用户读取敏感数据。例如，将文件权限设置为: -rw-r -----

```go
ioutil.WriteFile(p, []byte("present"), 0640)
```



```
-rw------- (600)    只有拥有者有读写权限。
-rw------- (640)    只有拥有者和属组用户有读写权限。
-rw-r--r-- (644)    只有拥有者有读写权限；而属组用户和其他用户只有读权限。
-rwx------ (700)    只有拥有者有读、写、执行权限。
-rwxr-xr-x (755)    拥有者有读、写、执行权限；而属组用户和其他用户只有读、执行权限。
-rwx--x--x (711)    拥有者有读、写、执行权限；而属组用户和其他用户只有执行权限。
-rw-rw-rw- (666)    所有用户都有文件读、写权限。
-rwxrwxrwx (777)    所有用户都有读、写、执行权限。
```




0x05 Open Redirect 重定向


不要直接重定向到用户可控制的地址。





```go
func redirect(c *gin.Context) {
    loc := c.Query("redirect")
    c.Redirect(302, loc)
}
```



```
GET /api/vul/redirect?redirect=https://www.qq.com HTTP/1.1
Host: 127.0.0.1:8080
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


```



### 0x42 FIX 修复

```go
func SafeRedirect(c *gin.Context) {
	baseUrl := "https://baidu.com/path?q="
	loc := c.Query("redirect")
	c.Redirect(302, baseUrl+loc)
}
```

```
GET /api/safe/redirect?redirect=https://www.qq.com HTTP/1.1
Host: 127.0.0.1:8080
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


```




# 0x05 CORS

CORS请求保护不当会导致敏感信息泄露，因此应严格设置Access-Control-Allow-Origin以使用同源策略进行保护。

任意源


```go
func Cors1(c *gin.Context) {

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

}

```

```
GET /api/vul/cors1 HTTP/1.1
Host: 127.0.0.1:8080
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

HTTP/1.1 200 OK
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: *
Date: Sat, 05 Nov 2022 09:05:18 GMT
Content-Length: 0
Connection: close


```

任意添加源

```go
func Cors2(c *gin.Context) {

	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

}
```

```
GET /api/vul/cors2 HTTP/1.1
Host: 127.0.0.1:8080
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
Origin: zeo.cool

HTTP/1.1 200 OK
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: zeo.cool
Date: Sat, 05 Nov 2022 09:06:33 GMT
Content-Length: 0
Connection: close

```



### FIX 修复

直接白名单

code:

```go
func corsDemo2(c *gin.Context) {
    allowedOrigin := "https://test.com"
 
    c.Header("Access-Control-Allow-Origin", allowedOrigin)
    c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}
```



```
GET /api/safe/cors HTTP/1.1
Host: 127.0.0.1:8080
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
Origin: zeo.cool

HTTP/1.1 200 OK
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: https://test.com
Date: Sat, 05 Nov 2022 09:09:53 GMT
Content-Length: 0
Connection: close
```
