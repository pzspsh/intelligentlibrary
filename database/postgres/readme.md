# postgres数据库
### 一、Postgres数据库安装
#### 1、[官网下载postgres包](https://www.postgresql.org/docs/release/)
#### 2、解压postgres软件包
#### 3、到解压的postgres文件夹目录下创建data文件夹
#### 4、打开cmd终端切换到解压的posgres文件夹目录下
#### 5、执行启动postgres服务
```
1、终端cmd切换到解压的postgres为bin目录下执行如下：
bin>pg_ctl.exe start -D 你上面创建data文件夹路径 # 例如：D:\\pgsql\data

2、终端cmd切换到解压的postgres为bin目录下执行如下:
bin>psql -U postgres  #以postgres用户登录psql
用户 postgres 的口令：输入你的密码
psql (14.0)
输入 "help" 来获取帮助信息.

postgres=#   #这一步就可以你要执行的操作了
```


# PGSQL 学习之数组字段

PGSQL 支技在字段中存储数组，定义时有三种方法：

1字段名[],2指定 ARRAY 关键字 3 直接输入维度， 如： NAME VARCHAR(20) [] 或 NAME VARCHAR(20) ARRAY,

建表：

```sql
--数组类型
create table lr_array1(
    id integer,
    array_i integer[], --数字类型数组
    array_t    text[]);   --text类型数组
```

插入数据

```sql
--数组类型的插入方式两种
--第一种 
insert into test_array1(id, array_i, array_t) values(1 , '{1,2,3}', '{"abc","def"}');    
--第二种
insert into test_array1(id, array_i, array_t) values(2 ,array[4,5,6,7],array['h','d','s']);
```

　　查询

```sql
select array_i from test_array1 where id = 1;--查询数组名即可    
select array_i[1],array_t[1] from test_array1;--通过[]方式获取数据，下标从1开始
--查询array_i数据组中包含有3的 数据
```

```sql
--查询array_i数据组中包含有3的 数据
select * from lr_array1 WHERE array_i @>'{3}'

insert into lr_array1(id, array_i, array_t) values(1, '{1,2,3}', '{"abc","def"}');   
insert into lr_array1(id, array_i, array_t) values(2, '{1,2,3}', '{"abc","def"}');   
insert into lr_array1 values(5,'{3,2,1,0}','{"abc","cde","bef"}')
insert into lr_array1(id, array_i, array_t) values(3, array[4,5,6,7], array['h','d','s']);
```

