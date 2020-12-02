1、问题:
=====
我们在数据库操作的时候，比如dao层中当遇到一个sql.ErrNoRows时，是否应该Wrap这个error，抛给上层。为什么，应该怎么做请写出代码。

2、ErrNoRows的理解
=====
## （1）根据官方文档库

ErrNoRows is returned by Scan when QueryRow doesn't return a row. In such a case, QueryRow returns a placeholder Row value that defers this error until a Scan.

当执行QueryRow查询时，如果没法返回一行，则返回ErrNoRows。
## （2）当执行类似于
err := Db.QueryRow("SELECT content, author FROM posts where id=$1", id).Scan(&content, &author)    

的语句时，会有三种结果：ErrNoRows、err != nil、nil三种结果。    

ErrNoRows实际是是数据表中无记录；nil表示可以查询出结果；err !=nil表示其它异常，例如传入的查询参数数据类型不匹配。

3、根据老师课堂上讲的
=====
You should only handle errors once. Handling an error means inspecting the error value, and making a single decision.
error只处理一次。

个人理解，日志最后在顶层处理，不要在最底层的DAO层处理。所以相关的error要往上抛。    

但是不应该让普通用户直接看到底层的错误信息，所以除了记录日志之外，还要提供一种面向最终用户的错误信息提示。

4、问题的转化
=====
实际上是转为    

问题：errNoRows作为error，肯定是要往上抛，抛完之后，调用者接收的error是否还是sql.ErrNoRows？

5、问题的测试与解答
=====
##  5.1 准备测试用的数据库
###   5.1.1拟用Postgres数据库，基于Docker部署 
####   (1)拉取postgres Docker镜像
$ docker pull postgres
####   (2) 通过Docker运行Postgres数据库，postgres的密码是Greek#007
$ docker run -d -P --name postgres --mount type=bind,source=/home/longqiping/postgres,destination=/var/lib/postgresql/data -e POSTGRES_PASSWORD="Greek#007" -p 5432:5432 postgres
####   (3)测试连接 Postgres数据库服务器的地址是10.3.134.110
$ psql -h 10.3.134.110 -U postgres --password -p 5432    


###   5.1.2 greek数据库
####   （1）通过psql运行data目录中的install.sq脚本，创建数据库greek和数据库用户greek，greek用户的密码是Greek#007，为用户greek授予数据库greek的完全权限。

install.sql脚本    

drop database greek;    

create database greek;    

drop user greek;    

create user greek with password 'Greek#007';     

grant all privileges on database greek to greek;     



$ psql -h 10.3.134.110 -U postgres --password -p 5432 -f install.sql




####    （2）运行安装脚本，创建执行相关操作所需要的表
$ psql -h 10.3.134.110 -U greek --password -p 5432 -f setup.sql -d greek


setup.sql脚本    

drop table posts;    

create table posts(    

        id      serial primary key,    

        content text,    

        author varchar(255)    

);    

一个简单的表，id是自动生成的作者编号，content表示贴子的内容，author表示作者。

###     5.1.3、准备数据表中的数据
在dataManagement目录中，构建了一个简单的REST的Web服务，用CRUD函数包裹一个Web服务接口，并通过JSON格式来传输数据。这个项目用于管理数据库的表数据。


##     5.2测试问题一
sql.errNoRows肯定是往上抛的，日志只记一次，关键是往上抛之后，errNoRows的信息能否保留而不失真。    

在目录testdatabase目录中，项目里查一个不存在的id=100，    

sql: no rows in result set   

true    

输出结果表明，此时的error是sql.errNoRows。    



有一条id号为4的记录，用data.Post(4)   

结果为    

{4 我的第一次POST longqiping} <nil>    

用data.Post(10)    

{0   } sql: no rows in result set    

用data.Post{"Hello"}    

将出现其它异常。

**结论：将sql.errNoRows往上抛可以保留error的相关信息，如果在底层是sql.errNoRows，往上抛后也同样是sql.errNoRows。**

##    5.3 问题二，既然往上抛后依然能保留sql.errNoRows属性，还有没有必要对sql.errNoRows进行Wrap。
**答案是，不进行Wrap也可以。**     
由于Db.QueryRow不存在panic的问题，所以不需要从panic中recover，因此只需要向包的调用者返回error。   

但是如果要附带其它信息，可以用结构体    

type AppError struct {    

	Message string    

	Err error 

}    

来附带其它信息。



