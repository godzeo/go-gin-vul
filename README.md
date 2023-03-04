# Go Gin vulnerability Range 

An example of gin contains many useful vul

一个go写的WEB漏洞靶场，实际自己写一下，加固一下知识，也偶尔测试waf的时候，用一下自己的靶场

The vulnerability websit with Go/GIN , the actual write their own, to strengthen the knowledge, but also occasionally test the WAF when you use their own range

整体的构建是最常见的 GIN 框架，加入了自己写的几个接口

GIN框架 整个web框架是go-gin-Example 上面改的，没有前端框架，只有一个swagger，直接发包吧

后期有时间再加前端吧，对于安全工程师前端是真的烦。


# 0x00 How to Code Run 安装部署🚀

## 0x01 docker 一键启动
首先刚刚加入了docker 可以一键启动，自行安装docker 和 docker-compose

```
cd docker
docker-compose up -d

```
访问 http://127.0.0.1:8080/  

漏洞接口就参考doc https://github.com/godzeo/go-gin-vul/blob/main/doc/vul.md


## 0x02 手动搭建
1 需要准备mysql和redis

2 修改配置conf/app.ini
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

3 自己创建对应的数据库 运行 conf/blog.sql 构建数据库

4 运行主程序
```
go mod tidy
go run main.go
```

# 漏洞文档
漏洞接口和代码分析在
https://github.com/godzeo/go-gin-vul/blob/main/doc/vul.md
