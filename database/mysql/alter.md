# 表结构修改(alter)

### 修改

```sql
ALTER TABLE teacher RENAME AS allteacher  --修改表名
ALTER TABLE allteacher ADD age INT(11)  --添加表的字段
--修改表的字段
ALTER TABLE allteacher MODIFY age VARCHAR(11)  --修改约束！
ALTER TABLE allteacher CHANGE age age1 INT(1)  --字段重命名！
```

### 删除

```sql
ALTER TABLE teacher1 DROP age1 --删除表的字段
DROP TABLE if EXISTS teacher1 --删除表，如果表存在。
```
