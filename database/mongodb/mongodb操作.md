# mongodb操作
```sql
1、MongoDB 创建数据库
键入以下命令在 MongoDB 中创建数据库：
use database_name
例如1
创建一个数据库beginnersbook，所以命令应该是：
use beginnersbook

2、创建 MongoDB 中的集合
方法一：动态创建 MongoDB 中的集合
例如：我们在数据库beginnersbookdb中没有集合beginnersbook。此命令将动态创建名为beginnersbook的集合，并使用指定的键和值对在其中插入文档。

> use beginnersbookdb switched to db beginnersbookdb
> db.beginnersbook.insert({
  name: "Chaitanya",
  age: 30,
  website: "beginnersbook.com"
})

您将在命令提示符中看到此响应。
WriteResult({ "nInserted" : 1 })

> db.beginnersbook.find()
{ "_id" : ObjectId("59bcb8c2415346bdc68a0a66"), "name" : "Chaitanya",
 "age" : 30, "website" : "beginnersbook.com" }

要检查是否已成功创建集合，请使用以下命令。
show collections
> show collections
beginnersbook

方法二：在插入文档之前使用选项创建集合
我们还可以在实际插入数据之前创建集合。此方法为您提供了在创建集合时可以设置的选项。
语法：
db.createCollection(name, options)
> db.createCollection("students")
{ "ok" : 1 }
让我们看看我们在创建集合时可以提供的选项：
capped：类型：布尔值。
此参数仅为true和false。这指定了集合可以拥有的最大条目的上限。一旦集合达到该限制，它就会开始覆盖旧条目。
这里需要注意的是，当您将capped选项设置为true时，您还必须指定size参数。
size：类型：数字。
指定集合的​​最大大小（上限集合），以字节为单位。
max：类型：数字。
指定集合可容纳的最大文档数。
autoIndexId：类型：布尔值
此参数的默认值为false。如果将其设置为true，则会自动为每个文档创建索引字段_id。我们将在 MongoDB 索引教程中了解索引
例子：
db.createCollection("teachers", { capped : true, size : 9232768} )
{ "ok" : 1 }

```
#### 插入文档
```sql
将文档插入集合的语法：
db.collection_name.insert()


MongoDB 使用insert()插入文档示例
MongoDB 示例：在集合中插入多个文档
这里我们将文档插入名为beginnersbook的集合中。下面示例中的字段course是一个包含多个键值对的数组。
db.beginnersbook.insert(  
   {  
     name: "Chaitanya",  
     age: 30,
     email: "[email protected]",
     course: [ { name: "MongoDB", duration: 7 }, { name: "Java", duration: 30 } ]
   }  
)

您应该看到一条成功的写消息，如下所示：
WriteResult({ "nInserted" : 1 })

如果集合不存在，则insert()方法创建集合，但如果集合存在，则将文档插入其中
MongoDB 插入文档


您还可以通过键入以下命令验证文档是否已成功插入：
db.collection_name.find()

在上面的示例中，我们将文档插入名为beginnersbook的集合中，因此命令应为：
> db.beginnersbook.find()
{ "_id" : ObjectId("59bce797668dcce02aaa6fec"), "name" : "Chaitanya", "age" : 30, 
"email" : "[email protected]", "course" : [ { "name" : "MongoDB", 
"duration" : 7 }, { "name" : "Java", "duration" : 30 } ] }

MongoDB 示例：在集合中插入多个文档
要在集合中插入多个文档，我们定义一个文档数组，稍后我们在数组变量上使用insert()方法，如下例所示。这里我们在名为students的集合中插入三个文档。此命令将在students集合中插入数据，如果集合不存在，则它将创建集合并插入这些文档。
var beginners =
 [
    {
    "StudentId" : 1001,
    "StudentName" : "Steve",
        "age": 30
    },
    {
    "StudentId" : 1002,
    "StudentName" : "Negan",
        "age": 42
    },
    {
    "StudentId" : 3333,
    "StudentName" : "Rick",
        "age": 35
    },
];
db.students.insert(beginners);

你会看到这个输出：
BulkWriteResult({
        "writeErrors" : [ ],
        "writeConcernErrors" : [ ],
        "nInserted" : 3,
        "nUpserted" : 0,
        "nMatched" : 0,
        "nModified" : 0,
        "nRemoved" : 0,
        "upserted" : [ ]
})

如您所见，它在nInserted前面显示数字 3。这意味着此命令已插入 3 个文档。
MongoDB 插入文档

验证文档是否在集合中。运行此命令：
db.students.find()

你知道吗？ 您可以以 JSON 格式打印输出数据，以便您可以轻松阅读。要以 JSON 格式打印数据，请运行命令db.collection_name.find().forEach(printjson)

所以在我们的例子中命令是这样的：
db.students.find().forEach(printjson)

在下面的屏幕截图中，您可以看到差异。首先，我们使用普通的find()方法打印文档，然后使用 JSON 格式打印相同集合的文档。 JSON 格式的文档简洁易读。

```