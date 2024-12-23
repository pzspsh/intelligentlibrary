# 删除数据(delete)

### 数据库删除

```sql
drop database 数据库名称;
```

### 数据库表删除

```sql
drop table table_name;
```

### 数据库表数据删除

```sql
delete from table_name where id = 1;

例：
delete from demo_table_name where title ='mysql_demo';
```

### 场景一：数据表之间不存在外键联系下同时删除多表数据

```sql
delete 数据表一,数据表二 from 数据表一,数据表二 where 条件

例：
delete student,studentquery from student,studentquery
where student.studentid = 1813004 AND studentquery.studentid=1813004;
```

### 场景二：数据表之间存在外键联系下同时删除多表数据

```sql
// 第一步：关闭数据库外键约束，否则无法删除
set foreign_key_checks = 0;

// 第二步：执行删除语句
delete 数据表一,数据表二 from 数据表一,数据表二 where 条件

// 第三步：开启数据库外键约束
set foreign_key_checks = 1

例：
set foreign_key_checks = 0;

delete student,studentquery from student,studentquery
where student.studentid = studentquery.studentid AND student.studentid=1813004;

set foreign_key_checks = 1
```
