# sql 语句实例

```sql
实例一：SQL检索所有行和列 解决方案
用特殊符号*对该表执行 SELECT 查询。
select * from student

SQL检索所有行和列 扩展知识
在 SQL 中，符号*有着特殊含义。该符号使得查询语句返回指定表的所有列。由于没有指定 WHERE 子句，因此所有行都会被提取出来。你也可以使用另一种方法，列出表中的每一列。
select number,name,job,sal,mgr,hiredate,comm,deptno from student


实例二：你有一张表，并且只想查看满足指定条件的行。SQL 筛选行 解决方案
使用 WHERE 子句指明保留哪些行。例如，下面的语句将查找部门编号为 10001 的所有员工。
select * from student where deptno = 10001

SQL 筛选行 扩展知识
可以使用 WHERE 子句来筛选出我们感兴趣的行。如果 WHERE 子句的表达式针对某一行的判定结果为真，那么就会返回该行的数据。
大多数数据库都支持常用的运算符，例如 =、<、>、<=、>=、! 和 <>。除此之外，你可能需要指定多个条件来筛选数据，这时就需要使用 AND、OR 和圆括号

实例三：SQL 查找满足多个查询条件的行 解决方案
使用带有 OR 和 AND 条件的 WHERE 子句。例如，如果你想找出部门编号为 10 的所有员工、有业务提成的所有员工以及部门编号是 20 且工资低于 2000 美元的所有员工。
select * from emp where deptno = 10 or comm is not null or sal <= 2000 and deptno=20

SQL 查找满足多个查询条件的行 扩展知识
你可以组合使用 AND、OR 和圆括号来筛选满足多个查询条件的行。在这个实例中，WHERE 子句找出了如下的数据。
DEPTNO 等于 10，或 COMM 不是 Null，或 DEPTNO 等于 20 且工资不高于 2000 美元的员工。
圆括号里的查询条件被一起评估。例如，试想一下如果采用下面的做法，检索结果会发生什么样的变化。

实例四：指定你感兴趣的列。例如，只查看员工的名字、部门编号和工资。SQL 筛选列 解决方案
select ename,deptno,sal from emp

SQL 筛选列 扩展知识
在 SELECT 语句里指定具体的列名，可以确保查询语句不会返回无关的数据。当在整个网络范围内检索数据时，这样做尤为重要，因为它避免了把时间浪费在检索不需要的数据上。

实例五：SQL 创建有意义的列名
你可能想要修改检索结果的列名，使其更具可读性且更易于理解。考虑下面这个查询，它返回的是每个员工的工资和业务提成。
select sal,comm from emp

sal 指的是什么？是 sale 的缩写吗？是人名吗？ comm 又是什么？是 communication 的缩写吗？显然，检索结果应该让人容易理解。
SQL 创建有意义的列名 解决方案
使用 AS 关键字，并以 original_name AS new_name 的形式来修改检索结果的列名。对于一些数据库而言，AS 不是必需的，但所有的数据库都支持这个关键字。
select sal as salary, comm as commission from emp

实例六：SQL 在WHERE子句中引用别名列
你已经为检索结果集创建了有意义的列名，并且想利用 WHERE 子句过滤掉部分行数据。但是，如果你尝试在 WHERE 子句中引用别名列，查询无法顺利执行。
select sal as salary, comm as commission from emp where salary < 5000

SQL 在WHERE子句中引用别名列 解决方案
把查询包装为一个内嵌视图，这样就可以引用别名列了。
select * from (select sal as salary, comm as commission from emp) x where salary < 5000

内嵌视图的别名为 x。并非所有数据库都需要给内嵌视图取别名，但对于某些数据库而言，确实必须如此。不过，所有的数据库都支持这一点。
SQL 在WHERE子句中引用别名列 扩展知识
在这个简单的实例中，你可以不使用内嵌视图。在 WHERE 子句里直接引用 COMM 列和 SAL 列，也可以达到同样的效果。当你想在 WHERE 子句中引用下列内容时，这个解决方案告诉你该如何做。

实例七：SQL 串联多列的值 解决方案
使用数据库中的内置函数来串联多列的值。
DB2、Oracle 和 PostgreSQL

这些数据库把双竖线作为串联运算符。
select ename||' WORKS AS A '||job as msg from emp where deptno=10

MySQL
该数据库使用 CONCAT 函数。
select concat(ename, ' WORKS AS A ',job) as msg from emp where deptno=10

SQL Server
该数据库使用+作为串联运算符。
select ename + ' WORKS AS A ' + job as msg from emp where deptno=10

SQL 串联多列的值 扩展知识
使用 CONCAT 函数可以串联多列的值。在 DB2、Oracle 和 PostgreSQL 中，||是 CONCAT 函数的快捷方式，在 SQL Server 中则为 +

实例八：SQL 在SELECT语句里使用条件逻辑 解决方案
在 SELECT 语句里直接使用 CASE 表达式来执行条件逻辑。
select ename,sal, case when sal <= 2000 then 'UNDERPAID' when sal >= 4000 then 'OVERPAID' else 'OK' end as status from emp

SQL 在SELECT语句里使用条件逻辑 扩展知识
CASE 表达式能对查询结果执行条件逻辑判断。你可以为 CASE 表达式的执行结果取一个别名，使结果集更有可读性。就本例而言，STATUS 就是 CASE 表达式执行结果的别名。ELSE 子句是可选的，若没有它，对于不满足测试条件的行，CASE 表达式会返回 Null。



```
