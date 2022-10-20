
# 常见SQL注入

若用户输入的字符未经任何处理直接拼接到执行SQL操作的方法，将会导致SQL注入。

拼接是指StringConcat操作，比如通过+、fmt.Sprintf()、buffer.WriteString()等方式将字符串连接到一起。

常见错误用法如下：

```
// 约定：taint为用户可控的数据，比如可以是req.user的值
db.Select(taint).First(&user) // taint="name; drop table users;"

db.Where(fmt.Sprintf("name = '%s'", taint)).Find(&user) // taint="name; drop table users;" db.Model(&user).Pluck(taint, &names) // taint="name; drop table users;"

db.Group(taint).First(&user) // taint="name; drop table users;"

db.Group("name").Having(taint).First(&user) // taint="1 = 1;drop table users;"

db.Raw("select name from " + taint).First(&user) // taint="users; drop table users;"

db.Exec("select name from " + taint).First(&user) // taint="users; drop table users;"

db.Order(taint).First(&user) // taint="name; drop table users;"

if strings.Contains(taint, "tom") {
db.Order(req.loc)
}
```


对于开发者来讲，SQL注入的修复主要有两种场景：1. 常规value的拼接; 2. 表/列名的拼接。

一个大的原则是，优先使用预编译，无法使用预编译的，使用白名单/转义操作进行防御。

对于常规value的拼接
对于除了表/列名的常规值，使用参数化查询对语句进行预编译，可以完全防御SQL注入，并且会提升查询效率，推荐使用。
表/列名的拼接
对于表/列名，无法使用预编译，可使用白名单机制，或者对特殊字符进行转义来防御注入。


方案一：常规value
````
// 对于常规的value
db.Where("name = ?", taint).First(&user)
// 对于like
db.Like("name like %?%", taint).First(&users)
// 对于exec
db.Exec("select name from users where name = ?", taint).First(&user)
// 对于raw
db.Raw("select name from users where name = ?", taint).First(&user)
// 对于in查询
db.Where("name in (?)", []string{taint1, taint2}).Find(&users)

````


方案二：表/列名
````
由于表/列名无法使用参数化查询，所以推荐使用白名单或转义操作。转义最基本的要对single quote、double quotes、back quote、slash、backslash进行转义。

// 对于列名
validCols := map[string]bool{"col1": true, "col2":true}

if _, ok := validCols[taint]; !ok {
fmt.Println("illegal column")
return
}
db.Order(taint)

// 对于表名
ValidTables := map[string]bool{"table1": true, "table2": true}
if _, ok := validTables[taint]; !ok {
fmt.Println("illegal table")
return
}
sql := fmt.Sprintf("select name from %s", taint)
db.Exec(sql).First(&user)
````

