# 索引(index)
### MySQL索引是什么
```
索引是一种特殊的数据库结构，由数据表中的一列或多列组合而成，可以用来快速查询数据表中有某一特定值的记录。本节将详细讲解索引的含义、作用和优缺点。

通过索引，查询数据时不用读完记录的所有信息，而只是查询索引列。否则，数据库系统将读取每条记录的所有信息进行匹配。

可以把索引比作新华字典的音序表。例如，要查“库”字，如果不使用音序，就需要从字典的 400 页中逐页来找。但是，如果提取拼音出来，构成音序表，就只需要从 10 多页的音序表中直接查找。这样就可以大大节省时间。

因此，使用索引可以很大程度上提高数据库的查询速度，还有效的提高了数据库系统的性能。
为什么要使用索引
索引就是根据表中的一列或若干列按照一定顺序建立的列值与记录行之间的对应关系表，实质上是一张描述索引列的列值与原表中记录行之间一 一对应关系的有序表。

索引是 MySQL 中十分重要的数据库对象，是数据库性能调优技术的基础，常用于实现数据的快速检索。

在 MySQL 中，通常有以下两种方式访问数据库表的行数据：
1) 顺序访问
顺序访问是在表中实行全表扫描，从头到尾逐行遍历，直到在无序的行数据中找到符合条件的目标数据。

顺序访问实现比较简单，但是当表中有大量数据的时候，效率非常低下。例如，在几千万条数据中查找少量的数据时，使用顺序访问方式将会遍历所有的数据，花费大量的时间，显然会影响数据库的处理性能。
2) 索引访问
索引访问是通过遍历索引来直接访问表中记录行的方式。

使用这种方式的前提是对表建立一个索引，在列上创建了索引之后，查找数据时可以直接根据该列上的索引找到对应记录行的位置，从而快捷地查找到数据。索引存储了指定列数据值的指针，根据指定的排序顺序对这些指针排序。

例如，在学生基本信息表 tb_students 中，如果基于 student_id 建立了索引，系统就建立了一张索引列到实际记录的映射表。当用户需要查找 student_id 为 12022 的数据的时候，系统先在 student_id 索引上找到该记录，然后通过映射表直接找到数据行，并且返回该行数据。因为扫描索引的速度一般远远大于扫描实际数据行的速度，所以采用索引的方式可以大大提高数据库的工作效率。

简而言之，不使用索引，MySQL 就必须从第一条记录开始读完整个表，直到找出相关的行。表越大，查询数据所花费的时间就越多。如果表中查询的列有一个索引，MySQL 就能快速到达一个位置去搜索数据文件，而不必查看所有数据，这样将会节省很大一部分时间。
索引的优缺点
索引有其明显的优势，也有其不可避免的缺点。
优点
索引的优点如下：
通过创建唯一索引可以保证数据库表中每一行数据的唯一性。
可以给所有的 MySQL 列类型设置索引。
可以大大加快数据的查询速度，这是使用索引最主要的原因。
在实现数据的参考完整性方面可以加速表与表之间的连接。
在使用分组和排序子句进行数据查询时也可以显著减少查询中分组和排序的时间
缺点
增加索引也有许多不利的方面，主要如下：
创建和维护索引组要耗费时间，并且随着数据量的增加所耗费的时间也会增加。
索引需要占磁盘空间，除了数据表占数据空间以外，每一个索引还要占一定的物理空间。如果有大量的索引，索引文件可能比数据文件更快达到最大文件尺寸。
当对表中的数据进行增加、删除和修改的时候，索引也要动态维护，这样就降低了数据的维护速度。

使用索引时，需要综合考虑索引的优点和缺点。

索引可以提高查询速度，但是会影响插入记录的速度。因为，向有索引的表中插入记录时，数据库系统会按照索引进行排序，这样就降低了插入记录的速度，插入大量记录时的速度影响会更加明显。这种情况下，最好的办法是先删除表中的索引，然后插入数据，插入完成后，再创建索引。
```