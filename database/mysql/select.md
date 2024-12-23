# 查询数据(select)

### MySQL 数据库中查询数据通用的 SELECT 语法：

```sql
SELECT column_name,column_name FROM table_name [WHERE Clause][LIMIT N][ OFFSET M]

查询语句中你可以使用一个或者多个表，表之间使用逗号(,)分割，并使用WHERE语句来设定查询条件。
SELECT 命令可以读取一条或者多条记录。
你可以使用星号（*）来代替其他字段，SELECT语句会返回表的所有字段数据
你可以使用 WHERE 语句来包含任何条件。
你可以使用 LIMIT 属性来设定返回的记录数。
你可以通过OFFSET指定SELECT语句开始查询的数据偏移量。默认情况下偏移量为0。
```

#### 实例：

```sql
select * from demo_table_name;

select * from demo_table_name where title='mysql_demo';

# BINARY 关键字来设定 WHERE 子句的字符串比较是区分大小写的
select * from demo_table_name where binary title='mysql_demo';
```

### 正则查询

```sql
查找name字段中以'st'为开头的所有数据：
SELECT name FROM table_name WHERE name REGEXP '^st';

查找name字段中以'ok'为结尾的所有数据：
mysql> SELECT name FROM table_name WHERE name REGEXP 'ok$';

查找name字段中包含'mar'字符串的所有数据：
mysql> SELECT name FROM table_name WHERE name REGEXP 'mar';

查找name字段中以元音字符开头或以'ok'字符串结尾的所有数据：
mysql> SELECT name FROM table_name WHERE name REGEXP '^[aeiou]|ok$';
```
