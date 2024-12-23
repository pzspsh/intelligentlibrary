# orm 连接映射操作数据库-postgres

# gorm 自定义时间、字符串数组类型

> GORM 是 GO 语言中一款强大友好的 ORM 框架，但在使用过程中内置的数据类型不能满足以下两个需求，如下：
>
> 1. `time.Time`类型返回的是 `2023-10-03T09:12:08.53528+08:00`这种字符串格式，需要额外处理，我们更希望默认的是是`2023-10-03 09:12:08`这种可读性更高的格式
> 2. 有些数据字段需要存储数组形式，如下`Article `中`Tags`字段希望保存不确定个字符串。直接保存会提示`[error] unsupported data type: &[]`
>
> 官方提供了` Scanner` 和 `Valuer`两个接口，来满足自定义数据的存储、提取，本文记录上述两种结构解决方法。

```go
type Article struct {
	Tags  []string
}
```

### 自定义时间类型

> 自定义时间类型需满足以下两个需求:
>
> - 返回`2006-01-02 15:04:05`格式
> - `CreatedAt`和`UpdatedAt`使用时能按`GORM`规范自动填充当前时间
>
> 默认的`time.Time`是能被`gorm`自动存储，取出到结构体的，返回`2023-10-03T09:12:08.53528+08:00`格式原因是在于`json`序列化时，这里解决方案是自定义一个数据结构，添加`JSON Marshal接口`，但是自定义的数据类型`gorm`不能识别，所以要额外添加`gorm`的` Scanner` 和 `Valuer`两个接口

```go
type CustomTime time.Time

// GORM Scanner 接口, 从数据库读取到类型
func (t *CustomTime) Scan(value any) error {

	if v, ok := value.(time.Time); !ok {
		return errors.Errorf("failed to unmarshal CustomTime value: %v", value)
	} else {
		*t = CustomTime(v)
		return nil
	}
}

// GORM Valuer 接口, 保存到数据库
func (t CustomTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t), nil
}

// JSON Marshal接口，CustomTime结构体转换为json字符串
func (t *CustomTime) MarshalJSON() ([]byte, error) {
	t2 := time.Time(*t)
	return []byte(fmt.Sprintf(`"%v"`, t2.Format("2006-01-02 15:04:05"))), nil
}
```

### 自定义字符串数组

> 代码比较简单，直接定义一个类型实现` Scanner` 和 `Valuer`两个接口，使用中将列定义为`Strings`类型即可

```go
type Strings []string

func (s *Strings) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), s)
}
func (s Strings) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}
```

### 测试与完整代码

```go
package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Strings []string

func (s *Strings) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), s)
}
func (s Strings) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return string(b), err
}

type CustomTime time.Time

// GORM Scanner 接口, 从数据库读取到类型
func (t *CustomTime) Scan(value any) error {

	if v, ok := value.(time.Time); !ok {
		return errors.Errorf("failed to unmarshal CustomTime value: %v", value)
	} else {
		*t = CustomTime(v)
		return nil
	}
}

// GORM Valuer 接口, 保存到数据库
func (t CustomTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return time.Time(t), nil
}

// JSON Marshal接口，CustomTime结构体转换为json字符串
func (t *CustomTime) MarshalJSON() ([]byte, error) {
	t2 := time.Time(*t)
	return []byte(fmt.Sprintf(`"%v"`, t2.Format("2006-01-02 15:04:05"))), nil
}

// fmt.Printf, 【可选方法】
func (t CustomTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

type Article struct {
	ID        uint `gorm:"primaryKey"`
	Tags      Strings
	CreatedAt CustomTime
	UpdatedAt CustomTime
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Article{})
	db.Create(&Article{Tags: []string{"go", "java"}})

	var article Article

	db.Last(&article)
	fmt.Printf("article is: %v\n", article)
	b, _ := json.Marshal(&article)
	fmt.Printf("article json is: %s\n", string(b))

	time.Sleep(time.Second * 30)
	article.Tags = append(article.Tags, "python")
	db.Save(&article)

	db.Last(&article)
	fmt.Printf("updated article is: %v\n", article)
	b, _ = json.Marshal(&article)
	fmt.Printf("updated article json is: %s\n", string(b))
}
```

#### 测试结果

> 字符串数组

![在这里插入图片描述](../../../../../images/PG字符串数组测试结果.png)

> 自定义时间,可以看到满足`2006-01-02 15:04:05`格式输出，以及时间自动添加和更新
> ![在这里插入图片描述](../../../../../images/PG自定义类型测试结果.png)

```go
package abc

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/lib/pq"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type DemoTest struct {
	Id        int64      `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号" json:"id"`
    Number    string     `gorm:"column:number;type:varchar(255);uniqueIndex;not null;comment:C段编号" json:"number"`
	Port      MapStrings `gorm:"column:port" json:"port,omitempty"`
	Target    string     `gorm:"column:target;type:varchar(255);default:null" json:"target,omitempty"`
}

type Port struct {
	Port    int            `json:"port,omitempty"`
	Url     string         `json:"url,omitempty"`
	Title   string         `json:"title,omitempty"`
    Status  int            `json:"status,omitempty"`
	Fingers pq.StringArray `json:"fingers,omitempty"`
}

func (DemoTest) TableName() string {
	return "demo_test"
}

func (dt *DemoTest) Insert(db *gorm.DB) error {
	err := db.Create(zc)
	if err.Error != nil {
		logger.Error("Number:%v, DemoTest Insert error:%v", dt.Number, err.Error)
		return err.Error
	} else {
		logger.Success("Number:%v, DemoTest Insert Successful", dt.Number)
		return nil
	}
}

func (dt *DemoTest) Update(obj, objId string, db *gorm.DB) error {
	err := db.Where(fmt.Sprintf("%v = ?", obj), objId).Updates(dt)
	if err.Error != nil {
		logger.Error("DemoTest number:%v Update error:%v\n", dt.Number, err.Error)
		return err.Error
	} else {
		logger.Success("DemoTest number:%v Update successfully\n", dt.Number)
		return nil
	}
}

func (dt *DemoTest) Select(sobj, objId string, db *gorm.DB) (bool, int64) {
	err := db.Where(fmt.Sprintf("%v=?", sobj), objId).First(dt)
	if err.Error != nil {
		logger.Error("DemoTest Number:%v select Information error:%v", objId, err.Error)
		return false, -1
	}
	return true, dt.Id
}


func DemoMigrate(es *elastic.Client, db *gorm.DB) {
	demo, err := es.Scroll("index").TrackTotalHits(true).Scroll("10m").Do(context.Background())
	if err != nil {
		fmt.Println("scroll error:", err)
	}
	scrollid := demo.ScrollId
	for {
		if len(demo.Hits.Hits) > 0 {
			demodata := make(map[string]interface{})
			for _, item := range demo.Each(reflect.TypeOf(demodata)) {
				demotest := DemoTest{}
				mapport := []Port{}
				number, ok := item.(map[string]interface{})["number"]
				if ok {
					demotest.Number = number.(string)
					fmt.Println("aaaaaaaaaa", number)
				} else {
					continue
				}
				ip := item.(map[string]interface{})["ip"].(string)
				target := item.(map[string]interface{})["target"].(string)
				ports := item.(map[string]interface{})["port"]
				portlist := ports.([]interface{})
				for _, portinfo := range portlist {
					port := Port{}
					portmap := portinfo.(map[string]interface{})
					sport := portmap["port"]
					switch sport.(type) {
					case string:
						port.Port, _ = strconv.Atoi(fmt.Sprintf("%v", sport))
					default:
						port.Port = int(sport.(float64))
					}
					port.Url = portmap["url"].(string)
					port.Title = portmap["title"].(string)
					status := portmap["status"]
					switch status.(type) {
					case string:
						port.Status, _ = strconv.Atoi(fmt.Sprintf("%v", status))
					default:
						port.Status = int(status.(float64))
					}
					port.Server = portmap["server"].(string)
					fingers, ok := portmap["fingers"]
					if ok {
						switch fingers := fingers.(type) {
						case []interface{}:
							port.Fingers = TostringArray(fingers)
						case []string:
							port.Fingers = fingers
						default:
							port.Fingers = []string{}
						}
					}
					mapport = append(mapport, port)
				}
				demotest.Ip = ip
				demotest.Target = target
				demotest.Port = mapport
				demotest.Insert(db)
				fmt.Println(demotest)
			}
		} else {
			break
		}
		if demo.TotalHits() > int64(len(demo.Hits.Hits)) {
			demo, err = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
			if err != nil {
				fmt.Println("scrollId error:", err)
			}
			scrollid = demo.ScrollId
		} else {
			break
		}
	}
}

type MapStrings []Port

func (ms *MapStrings) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), ms)
}

func (ms MapStrings) Value() (driver.Value, error) {
	b, err := json.Marshal(ms)
	return string(b), err
}

func TostringArray(ss []interface{}) []string {
	var slicestring []string
	for _, value := range ss {
		slicestring = append(slicestring, value.(string))
	}
	return slicestring
}

```

### pg 数据库创建 json 字段示例

```go
import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TestDemo struct {
	Id        int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号" json:"id"`
	Number    string         `gorm:"column:number;type:varchar(255);uniqueIndex;not null;comment:自定义编号" json:"number"`
	Name      string         `gorm:"column:name;type:varchar(255);default:not null;comment:任务名称" json:"name"`
	Args1     datatypes.JSON `gorm:"column:args1;serializer:json;comment:资产扫描参数" json:"args1"`
	Args2     datatypes.JSON `gorm:"column:args2;serializer:json;comment:指纹识别参数" json:"args2"`
	Args3     datatypes.JSON `gorm:"column:args3;serializer:json;comment:漏洞扫描参数" json:"args3"`
	Target    pq.StringArray `gorm:"column:target;type:text[];default:null;comment:任务目标数据" json:"target"`
	Status    int            `gorm:"column:status;default:null;comment:任务状态" json:"status,omitempty"`
	Dumber    string         `gorm:"column:dumber;type:varchar(255);default:not null;comment:单位编号" json:"dw_number,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime;comment:更新时间" json:"update_at,omitempty"`
}

func (TestDemo) TableName() string {
	return "test_demo"
}
```

```go
import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/lib/pq"
)

type SubDomain struct {
	Id         int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号" json:"id"`
	Domains    pq.StringArray `gorm:"column:domains;type:text[];default:null;comment:域名列表" json:"domains,omitempty"`
	Ips        pq.StringArray `gorm:"column:ips;type:text[];default:null;comment:ip列表" json:"ips,omitempty"`
	DomainsUrl DdomainData    `gorm:"column:domainsUrl;embedded" json:"domainsUrl,omitempty"`
	IpSegments IpsegmentData  `gorm:"column:ipSegments;embedded" json:"ipSegments,omitempty"`
	Protocol   Jsondata       `gorm:"column:protocol;embedded;comment:协议统计" json:"protocol,omitempty"`
	Port       Jsondata       `gorm:"column:port;embedded;comment:端口统计" json:"port,omitempty"`
	Location   Jsondata       `gorm:"column:location;embedded;comment:" json:"location,omitempty"`
	Status     int            `gorm:"column:status;default:null" json:"status,omitempty"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime;comment:更新时间" json:"updated_at,omitempty"`
}

func (SubDomain) TableName() string {
	return "test_subdomain"
}

type DdomainData []DomainList

func (ms *DdomainData) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), ms)
}

func (ms DdomainData) Value() (driver.Value, error) {
	b, err := json.Marshal(ms)
	return string(b), err
}

type IpsegmentData []IpSegmentList

func (ms *IpsegmentData) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), ms)
}

func (ms IpsegmentData) Value() (driver.Value, error) {
	b, err := json.Marshal(ms)
	return string(b), err
}

type DomainList struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号" json:"id"`
	Number    string    `gorm:"column:number;type:varchar(255);uniqueIndex;default:not null;comment:域名编号" json:"number"`
	Level     int       `gorm:"column:level;default:null" json:"level,omitempty"`
	Cname     string    `gorm:"column:cname;type:text;default:null" json:"cname,omitempty"`
	Ip        string    `gorm:"column:ip;type:varchar(255);default:null" json:"ip,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime;comment:更新时间" json:"updated_at,omitempty"`
}

type IpSegmentList struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号" json:"id"`
	Number    string    `gorm:"column:number;type:varchar(255);uniqueIndex;default:not null;comment:IP段编号" json:"number"`
	IpSegment string    `gorm:"column:ip_segment;type;varchar(255);default:null" json:"ipSegment,omitempty"`
	Count     int       `gorm:"column:count;default:null" json:"count,omitempty"`
	Data      DataList  `gorm:"column:data;embedded" json:"data,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;comment:创建时间" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime;comment:更新时间" json:"updated_at,omitempty"`
	Email     string    `gorm:"type:varchar(100);unique_index"`
	Role      string    `gorm:"size:255"`        //设置字段的大小为255个字节
	MeNumber  *string   `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
	Num       int       `gorm:"AUTO_INCREMENT"`  // 设置 Num字段自增
	IgnoreMe  int       `gorm:"-"`               //忽略这个字段
	Id        uint      `json:"id" gorm:"column:id;type:int(10) unsigned not null AUTO_INCREMENT;primaryKey;"`
    UserName  string    `json:"user_name" gorm:"column:user_name;type:varchar(16) not null;default:'';index:idx_user_name"`
    Password  string    `json:"-" gorm:"column:password;type:varchar(128) not null;default:''"`
    Status    uint      `json:"status" gorm:"column:status;type:tinyint(1) unsigned not null;default:0;index:idx_status"`
    CreatTime int64     `json:"created_time" gorm:"column:created_time;type:int(11) not null;default:0;index:idx_created_time"`
    UpdatTime int64     `json:"updated_time" gorm:"column:updated_time;type:int(11) not null;default:0;index:idx_updated_time"`
    DeletTime int64     `json:"-" gorm:"column:deleted_time;type:int(11) not null;default:0"`
	Avatar    []byte    `gorm:"cloumn:avatar;type:blob"`
}

type Data struct {
	Ip     string `json:"ip,omitempty"`
	Domain string `json:"domain,omitempty"`
}

func (IpSegmentList) TableName() string {
	return "test_ipsegment"
}

type DataList []Data

func (ms *DataList) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), ms)
}

func (ms DataList) Value() (driver.Value, error) {
	b, err := json.Marshal(ms)
	return string(b), err
}

type Jsondata map[string]int

func (jd *Jsondata) Scan(value any) error {
	v, _ := value.(string)
	return json.Unmarshal([]byte(v), jd)
}

func (jd Jsondata) Value() (driver.Value, error) {
	b, err := json.Marshal(jd)
	return string(b), err
}

```



### 字段标签

在声明模型时，标签是可选的，GORM支持以下标签：标签不区分大小写，但首选‘ camelCase ’。如果使用多个标签，它们应该用分号（';'）分隔。对解析器具有特殊意义的字符可以使用反斜杠（' \ '）进行转义，从而允许将它们用作参数值。

| 标签名                 | 说明                                                         |
| :--------------------- | :----------------------------------------------------------- |
| column                 | 指定 db 列名                                                 |
| type                   | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：`not null`、`size`, `autoIncrement`… 像 `varbinary(8)` 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：`MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT` |
| serializer             | 指定将数据序列化或反序列化到数据库中的序列化器, 例如: `serializer:json/gob/unixtime` |
| size                   | 定义列数据类型的大小或长度，例如 `size: 256`                 |
| primaryKey             | 将列定义为主键                                               |
| unique                 | 将列定义为唯一键                                             |
| default                | 定义列的默认值                                               |
| precision              | 指定列的精度                                                 |
| scale                  | 指定列大小                                                   |
| not null               | 指定列为 NOT NULL                                            |
| autoIncrement          | 指定列为自动增长                                             |
| autoIncrementIncrement | 自动步长，控制连续记录之间的间隔                             |
| embedded               | 嵌套字段                                                     |
| embeddedPrefix         | 嵌入字段的列名前缀                                           |
| autoCreateTime         | 创建时追踪当前时间，对于 `int` 字段，它会追踪时间戳秒数，您可以使用 `nano`/`milli` 来追踪纳秒、毫秒时间戳，例如：`autoCreateTime:nano` |
| autoUpdateTime         | 创建/更新时追踪当前时间，对于 `int` 字段，它会追踪时间戳秒数，您可以使用 `nano`/`milli` 来追踪纳秒、毫秒时间戳，例如：`autoUpdateTime:milli` |
| index                  | 根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 [索引](https://gorm.io/zh_CN/docs/indexes.html) 获取详情 |
| uniqueIndex            | 与 `index` 相同，但创建的是唯一索引                          |
| check                  | 创建检查约束，例如 `check:age > 13`，查看 [约束](https://gorm.io/zh_CN/docs/constraints.html) 获取详情 |
| <-                     | 设置字段写入的权限， `<-:create` 只创建、`<-:update` 只更新、`<-:false` 无写入权限、`<-` 创建和更新权限 |
| ->                     | 设置字段读的权限，`->:false` 无读权限                        |
| -                      | 忽略该字段，`-` 表示无读写，`-:migration` 表示无迁移权限，`-:all` 表示无读写迁移权限 |
| comment                | 迁移时为字段添加注释                                         |