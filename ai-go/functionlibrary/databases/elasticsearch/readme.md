### ElasticSearch Go

```go
/*
@File   : index.go
@Author : pan
@Time   : 2023-09-08 12:52:53
*/
package demo

import (
	"context"
	"fmt"
	"reflect"
	"time"

	logger "dmg/logging"
	"dmg/pkg/aft"

	"github.com/olivere/elastic/v7"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type IndexArgs struct {
	Id        int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string         `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Name      string         `gorm:"column:name;type:varchar(255);default:not null"`
	Args1     datatypes.JSON `gorm:"column:args1;serializer:json"`
	Args2     datatypes.JSON `gorm:"column:args2;serializer:json"`
	Args3     datatypes.JSON `gorm:"column:args3;serializer:json"`
	Target    []string       `gorm:"column:index;type:text[];default:null"`
	Status    int            `gorm:"column:status;default:null"`
	Open      int            `gorm:"column:open;default:null"`
	Level     int            `gorm:"column:level;default:null"`
	TaskType  int            `gorm:"column:index_type;default:null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
	// CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_at,omitempty"`
}

type IndexControl struct {
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number string `gorm:"column:number;type:varchar(255);unique_index;default:not null"`
	Status int    `gorm:"column:status;default:null"`
}

type IndexProgress struct {
	Id        int64     `gorm:"column:id,primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Total     int       `gorm:"column:total;type:varchar(255);default:null"`
	Success   int       `gorm:"column:success;type:varchar(255);default:null"`
	Failure   int       `gorm:"column:failure;type:varchar(255);default:null"`
	UsedTime  string    `gorm:"column:usedtime;type:varchar(255);default:null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

type IndexStatistic struct {
	Id          int64     `gorm:"column:id;primary_key,AUTO_INCREMENT;comment:自增编号"`
	Number      string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Asset       int       `gorm:"column:asset;default:null"`
	Ip          int       `gorm:"column:ip;default:null"`
	Domain      int       `gorm:"column:domain;default:null"`
	Cert        int       `gorm:"column:cert;default:null"`
	Port        int       `gorm:"column:port;default:null"`
	Protocol    int       `gorm:"column:protocol;default:null"`
	Fingerprint int       `gorm:"column:fingerprint;default:null"`
	Vuln        int       `gorm:"column:vuln;default:null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

func (IndexArgs) TableName() string {
	return "Index_args"
}

func (IndexControl) TableName() string {
	return "Index_control"
}

func (IndexProgress) TableName() string {
	return "Index_progress"
}

func (IndexStatistic) TableName() string {
	return "Index_statistic"
}

func Selectindex(es *elastic.Client, db *gorm.DB) *[]IndexArgs {
	// res, err := es.Search("index").TrackTotalHits(true).Do(context.Background())
	// res, err := es.Search("index").SearchAfter().TrackTotalHits(true).Do(context.Background())
	// res, err := es.Search("index").From(1000).Size(10).Do(context.Background())
	// res, err := es.Scroll("index").Do(context.Background())
	aftc := []aft.AftCompany{}
	companys := db.Find(&aftc)
	if companys.Error != nil {
		logger.Error("postgres db Find error:%v", companys.Error)
		return nil
	}
	for _, aft := range aftc {
		termquery := elastic.NewTermQuery("Number", aft.Number)
		indexes, err := es.Scroll("index").Query(termquery).Do(context.Background())
		if err != nil {
			logger.Error("demo index args search error:%v", err)
		}
		scrollid := indexes.ScrollId
		for {
			if len(indexes.Hits.Hits) > 0 {
				indexdata := make(map[string]interface{})
				for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
					number := item.(map[string]interface{})["number"]
					fmt.Println(number)
				}
			} else {
				break
			}
			indexes, _ = es.Scroll("index").Query(termquery).ScrollId(scrollid).Do(context.Background())
			scrollid = indexes.ScrollId
		}
	}
	res, err := es.Scroll("index").Do(context.Background())
	if err != nil {
		logger.Error("demo index args search error:%v", err)
	}
	// fmt.Println(res.TotalHits())
	// totallen := res.TotalHits()
	scrollid := res.ScrollId
	// index := make([]string, 0, res.TotalHits())
	// fmt.Println("totalhits", totallen)
	// fmt.Println(len(res.Hits.Hits))
	// fmt.Println("scrollid", scrollid)
	index := []string{}
	for {
		if len(res.Hits.Hits) > 0 {
			// for _, hit := range res.Hits.Hits {
			// 	fmt.Println(string(hit.Source))
			// index = append(index, string(hit.Source))
			// }
			indexdata := make(map[string]interface{})
			for _, item := range res.Each(reflect.TypeOf(indexdata)) {
				number := item.(map[string]interface{})["number"]
				// number := item.(map[string]interface{})["number"]
				// createTime := item.(map[string]interface{})["createTime"]
				// level := item.(map[string]interface{})["level"]
				// lastTime := item.(map[string]interface{})["lastTime"]
				fmt.Println(number)
			}
		} else {
			break
		}
		res, _ = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
		scrollid = res.ScrollId
	}
	fmt.Println("AAAAAAAAAAAAAAAAAAAAA", index)
	data := &[]IndexArgs{}
	return data
}

func (a *IndexArgs) Insert() {

}

func (a *IndexArgs) Update() {

}

func (a *IndexControl) Insert() {

}

func (a *IndexControl) Update() {

}

func (a *IndexProgress) Insert() {

}

func (a *IndexProgress) Update() {

}

func (a *IndexStatistic) Insert() {

}

func (a *IndexStatistic) Update() {

}

```

```go
package demo

import (
	"context"
	"fmt"
	"reflect"
	"time"

	logger "dmg/logging"
	"dmg/pkg/aft"

	"github.com/olivere/elastic/v7"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type IndexArgs struct {
	Id        int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string         `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Name      string         `gorm:"column:name;type:varchar(255);default:not null"`
	Args1     datatypes.JSON `gorm:"column:args1;serializer:json"`
	Args2     datatypes.JSON `gorm:"column:args2;serializer:json"`
	Args3     datatypes.JSON `gorm:"column:args3;serializer:json"`
	Target    []string       `gorm:"column:index;type:text[];default:null"`
	Status    int            `gorm:"column:status;default:null"`
	Open      int            `gorm:"column:open;default:null"`
	Level     int            `gorm:"column:level;default:null"`
	TaskType  int            `gorm:"column:index_type;default:null"`
	Owner     string         `gorm:"column:owner;type:varchar(255);default:not null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
	// CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_at,omitempty"`
}

type IndexControl struct {
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number string `gorm:"column:number;type:varchar(255);unique_index;default:not null"`
	Status int    `gorm:"column:status;default:null"`
}

type IndexProgress struct {
	Id        int64     `gorm:"column:id,primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Total     int       `gorm:"column:total;type:varchar(255);default:null"`
	Success   int       `gorm:"column:success;type:varchar(255);default:null"`
	Failure   int       `gorm:"column:failure;type:varchar(255);default:null"`
	UsedTime  string    `gorm:"column:usedtime;type:varchar(255);default:null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

type IndexStatistic struct {
	Id          int64     `gorm:"column:id;primary_key,AUTO_INCREMENT;comment:自增编号"`
	Number      string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Asset       int       `gorm:"column:asset;default:null"`
	Ip          int       `gorm:"column:ip;default:null"`
	Domain      int       `gorm:"column:domain;default:null"`
	Cert        int       `gorm:"column:cert;default:null"`
	Port        int       `gorm:"column:port;default:null"`
	Protocol    int       `gorm:"column:protocol;default:null"`
	Fingerprint int       `gorm:"column:fingerprint;default:null"`
	Vuln        int       `gorm:"column:vuln;default:null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

func (IndexArgs) TableName() string {
	return "index_args"
}

func (IndexControl) TableName() string {
	return "index_control"
}

func (IndexProgress) TableName() string {
	return "index_progress"
}

func (IndexStatistic) TableName() string {
	return "index_statistic"
}

func SelectIndex(es *elastic.Client, db *gorm.DB) *[]IndexArgs {
	aftc := []aft.AftCompany{}
	companys := db.Find(&aftc)
	if companys.Error != nil {
		logger.Error("postgres db Find error:%v", companys.Error)
		return nil
	}
	Indexcount := []string{}
	for _, aft := range aftc {
		logger.Info("company info:%v", aft.Number)
		termquery := elastic.NewBoolQuery()
		// termquery := elastic.NewTermQuery("Number", aft.Number)
		// termquery.Must(elastic.NewMatchQuery("Number", aft.Number)) // 特殊字符串匹配不了
		termquery.Must(elastic.NewMatchPhraseQuery("Number", aft.Number)) // 完全匹配
		indexes, err := es.Search().Index("index").Query(termquery).Sort("createTime", true).From(0).Size(10000).Do(context.Background())
		if err != nil {
			logger.Error("demo index args search error:%v", err)
			continue
		}
		if len(indexes.Hits.Hits) > 0 {
			indexdata := make(map[string]interface{})
			for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
				number := item.(map[string]interface{})["number"]
				fmt.Println(number)
				logger.Info("index number:%v", number)
				indexcount = append(indexcount, number.(string))
			}
		}
	}
	fmt.Println("BBBB", len(indexcount))
	// termquery := elastic.NewBoolQuery()
	// termquery.Must(elastic.NewMatchPhraseQuery("Number", "1681180178"))
	// indexes, err := es.Search().Index("index").Query(termquery).Do(context.Background())
	// if err != nil {
	// 	logger.Error("demo index args search error:%v", err)
	// }
	// if len(indexes.Hits.Hits) > 0 {
	// 	indexdata := make(map[string]interface{})
	// 	for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
	// 		number := item.(map[string]interface{})["number"]
	// 		fmt.Println(number)
	// 		logger.Info("index number:%v", number)
	// 		indexcount = append(indexcount, number.(string))
	// 	}
	// }
	// fmt.Println("BBBB", len(indexcount))

	// indexcount := []string{}
	// for _, aft := range aftc {
	// 	// logger.Info("company info:%v", aft.Number)
	// 	termquery := elastic.NewMatchQuery("Number", aft.Number)
	// 	indexes, err := es.Scroll("index").Query(termquery).TrackTotalHits(true).Do(context.Background())
	// 	if err != nil {
	// 		logger.Error("demo index args search error:%v", err)
	// 	}
	// 	// scrollid := indexes.ScrollId
	// 	for {
	// 		if len(indexes.Hits.Hits) > 0 {
	// 			indexdata := make(map[string]interface{})
	// 			for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
	// 				number := item.(map[string]interface{})["number"]
	// 				fmt.Println(number)
	// 				logger.Info("index number:%v", number)
	// 				indexcount = append(indexcount, number.(string))
	// 			}
	// 		} else {
	// 			break
	// 		}
	// 		// indexes, _ = es.Scroll("index").ScrollId(scrollid).Query(termquery).Do(context.Background())
	// 		// scrollid = indexes.ScrollId
	// 	}
	// }
	// fmt.Println("BBBB", len(indexcount))
	data := &[]IndexArgs{}
	return data
}

func (a *IndexArgs) Insert() {

}

func (a *IndexArgs) Update() {

}

func (a *IndexControl) Insert() {

}

func (a *IndexControl) Update() {

}

func (a *IndexProgress) Insert() {

}

func (a *IndexProgress) Update() {

}

func (a *IndexStatistic) Insert() {

}

func (a *IndexStatistic) Update() {

}
```

```go
/*
@File   : index.go
@Author : pan
@Time   : 2023-09-08 12:52:53
*/
package demo

import (
	"context"
	"fmt"
	"reflect"
	"time"

	logger "dmg/logging"
	"dmg/pkg/aft"

	"github.com/olivere/elastic/v7"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type IndexArgs struct {
	Id        int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string         `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Name      string         `gorm:"column:name;type:varchar(255);default:not null"`
	Args1     datatypes.JSON `gorm:"column:args1;serializer:json"`
	Args2     datatypes.JSON `gorm:"column:args2;serializer:json"`
	Args3     datatypes.JSON `gorm:"column:args3;serializer:json"`
	Target    []string       `gorm:"column:index;type:text[];default:null"`
	Status    int            `gorm:"column:status;default:null"`
	Open      int            `gorm:"column:open;default:null"`
	Level     int            `gorm:"column:level;default:null"`
	TaskType  int            `gorm:"column:index_type;default:null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
	// CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_at,omitempty"`
}

type IndexControl struct {
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number string `gorm:"column:number;type:varchar(255);unique_index;default:not null"`
	Status int    `gorm:"column:status;default:null"`
}

type IndexProgress struct {
	Id        int64     `gorm:"column:id,primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Total     int       `gorm:"column:total;type:varchar(255);default:null"`
	Success   int       `gorm:"column:success;type:varchar(255);default:null"`
	Failure   int       `gorm:"column:failure;type:varchar(255);default:null"`
	UsedTime  string    `gorm:"column:usedtime;type:varchar(255);default:null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

type IndexStatistic struct {
	Id          int64     `gorm:"column:id;primary_key,AUTO_INCREMENT;comment:自增编号"`
	Number      string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Asset       int       `gorm:"column:asset;default:null"`
	Ip          int       `gorm:"column:ip;default:null"`
	Domain      int       `gorm:"column:domain;default:null"`
	Cert        int       `gorm:"column:cert;default:null"`
	Port        int       `gorm:"column:port;default:null"`
	Protocol    int       `gorm:"column:protocol;default:null"`
	Fingerprint int       `gorm:"column:fingerprint;default:null"`
	Vuln        int       `gorm:"column:vuln;default:null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

func (IndexArgs) TableName() string {
	return "Index_args"
}

func (IndexControl) TableName() string {
	return "Index_control"
}

func (IndexProgress) TableName() string {
	return "Index_progress"
}

func (IndexStatistic) TableName() string {
	return "Index_statistic"
}

func SelectTask(es *elastic.Client, db *gorm.DB) *[]IndexArgs {
	aftc := []aft.AftCompany{}
	companys := db.Find(&aftc)
	if companys.Error != nil {
		logger.Error("postgres db Find error:%v", companys.Error)
		return nil
	}
	indexcount := []string{}
	// for _, aft := range aftc {
	// 	termquery := elastic.NewBoolQuery()
	// 	// termquery := elastic.NewTermQuery("Number", aft.Number)
	// 	// termquery.Must(elastic.NewMatchQuery("Number", aft.Number)) // 特殊字符串匹配不了
	// 	// matchquery := elastic.NewMatchPhraseQuery("Number", aft.Number)
	// 	termquery.Must(elastic.NewMatchPhraseQuery("Number", aft.Number)) // 完全匹配
	// 	// indexes, err := es.Search().Index("index").Query(termquery).Sort("createTime", true).From(1000).Size(10).Do(context.Background())
	// 	indexes, err := es.Search().Index("index").Query(termquery).From(0).Size(10000).Do(context.Background())
	// 	if err != nil {
	// 		logger.Error("demo index args search error:%v", err)
	// 		continue
	// 	}
	// 	fmt.Println(indexes.Hits.Hits)
	// 	if len(indexes.Hits.Hits) > 0 {
	// 		indexdata := make(map[string]interface{})
	// 		for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
	// 			number := item.(map[string]interface{})["number"]
	// 			logger.Info("Number:%v index Number:%v", aft.Number, number)
	// 			indexcount = append(indexcount, number.(string))
	// 		}
	// 	}
	// }
	// fmt.Println("BBBB", len(indexcount))
	for _, aft := range aftc {
		termquery := elastic.NewBoolQuery()
		termquery.Must(elastic.NewMatchPhraseQuery("Number", aft.Number))
		indexes, err := es.Scroll("index").Query(termquery).TrackTotalHits(true).Scroll("2m").Do(context.Background())
		if err != nil {
			logger.Error("Number:%v  demo index args search error:%v", aft.Number, err)
			continue
		}
		scrollid := indexes.ScrollId
		for {
			if len(indexes.Hits.Hits) > 0 {
				indexdata := make(map[string]interface{})
				for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
					number := item.(map[string]interface{})["number"]
					logger.Info("Number:%v index Number:%v", aft.Number, number)
					indexcount = append(indexcount, number.(string))
				}
			} else {
				break
			}
			if indexes.TotalHits() > int64(len(indexes.Hits.Hits)) {
				indexes, _ = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
				scrollid = indexes.ScrollId
			} else {
				break
			}
		}
	}
	fmt.Println("BBBB", len(indexcount))
	es.CloseIndex("index").Do(context.Background())
	es.OpenIndex("index").Do(context.Background())
	data := &[]IndexArgs{}
	return data
}

func (a *IndexArgs) Insert() {

}

func (a *IndexArgs) Update() {

}

func (a *IndexControl) Insert() {

}

func (a *IndexControl) Update() {

}

func (a *IndexProgress) Insert() {

}

func (a *IndexProgress) Update() {

}

func (a *IndexStatistic) Insert() {

}

func (a *IndexStatistic) Update() {

}

```

```go
/*
@File   : demo.go
@Author : pan
@Time   : 2023-09-08 12:48:31
*/
package demo

import (
	"context"
	logger "dmg/logging"
	"fmt"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
)

type ApsTargetArgs struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Target    string    `gorm:"column:index;type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
	// 	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at,omitempty"`
	// 	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_at,omitempty"`
}

type ApsTargetProgress struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Total     int       `gorm:"column:total;default:null"`
	Success   int       `gorm:"column:success;default:null"`
	Failure   int       `gorm:"column:failure;default:null"`
	UsedTime  string    `gorm:"column:usedtime;type:varchar(255);default:null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

type ApsTargetStatistic struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number      string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Asset       int       `gorm:"column:asset;default:null"`
	Ip          int       `gorm:"column:ip;default:null"`
	Domain      int       `gorm:"column:domain;default:null"`
	Cert        int       `gorm:"column:cert;default:null"`
	Port        int       `gorm:"column:port;default:null"`
	Protocol    int       `gorm:"column:protocol;default:null"`
	Fingerprint int       `gorm:"column:fingerprint;default:null"`
	Vuln        int       `gorm:"column:vuln;default:null"`
	CreatedAt   time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

func (ApsTargetArgs) TableName() string {
	return "index_args"
}

func (ApsTargetProgress) TableName() string {
	return "index_progress"
}

func (ApsTargetStatistic) TableName() string {
	return "index_statistic"
}

func SelectTarget(es *elastic.Client, redata []map[string]string) {
	argsData := ApsTargetArgs{}
	for _, data := range redata {
		Number := data["Number"]
		termatchquery := elastic.NewMatchPhraseQuery("Number", Number)
		indexes, err := es.Scroll("index").Query(termatchquery).TrackTotalHits(true).Scroll("2m").Do(context.Background())
		if err != nil {
			logger.Error("selectTarget es Scroll error:%v", err)
		}
		scrollid := indexes.ScrollId
		for {
			indexdata := make(map[string]interface{})
			if len(indexes.Hits.Hits) > 0 {
				for _, item := range indexes.Each(reflect.TypeOf(indexdata)) {
					number := item.(map[string]interface{})["number"]
					index := item.(map[string]interface{})["index"]
					createTime := item.(map[string]interface{})["createTime"]
					lastTime := item.(map[string]interface{})["lastTime"]
					argsData.Number = number.(string)
					argsData.Target = index.(string)
					argsData.RNumber = Number
					createtime, _ := time.ParseInLocation("2006-01-02T15:04:05", createTime.(string), time.Local)
					argsData.CreatedAt = createtime
					lasttime, _ := time.ParseInLocation("2006-01-02T15:04:05", lastTime.(string), time.Local)
					argsData.UpdatedAt = lasttime
				}
			} else {
				break
			}
			if indexes.TotalHits() > int64(len(indexes.Hits.Hits)) {
				indexes, _ = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
				scrollid = indexes.ScrollId
			} else {
				break
			}
		}
	}
	fmt.Println(indexes, data, termatchquery)
	// res, err := es.Scroll("index").Do(context.Background())
	// if err != nil {
	// 	logger.Error("demo index args search error:%v", err)
	// }
	// scrollid := res.ScrollId
	// index := []string{}
	// for {
	// 	if len(res.Hits.Hits) > 0 {
	// 		for _, hit := range res.Hits.Hits {
	// 			index = append(index, string(hit.Source))
	// 		}
	// 	} else {
	// 		break
	// 	}
	// 	res, _ = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
	// 	scrollid = res.ScrollId
	// }
	// fmt.Println("AAAAAAAAAAAAAAAAAAAAA", index)
}

func (a *ApsTargetArgs) Insert() {

}

func (a *ApsTargetArgs) Update() {

}

func (a *ApsTargetProgress) Insert() {

}

func (a *ApsTargetProgress) Update() {

}

func (a *ApsTargetStatistic) Insert() {

}

func (a *ApsTargetStatistic) Update() {

}

```

### ES 查询超过 1 万+数据量示例方法

```go
func SelectDemo(es *elastic.Client, taskdata []map[string]string) {
	es.CloseIndex("index").Do(context.Background())
	es.OpenIndex("index").Do(context.Background())
	logger.Info("SelectDemo taskdata Len:%v", len(taskdata))
	demoData := []DemoData{}
	for _, data := range taskdata {
		number := data["Number"]
		matchquery := elastic.NewBoolQuery()
		matchquery.Must(elastic.NewMatchPhraseQuery("Number", number))
		demotest, err := es.Scroll("index").Query(matchquery).TrackTotalHits(true).Scroll("100m").Do(context.Background())
		if err != nil {
			es.CloseIndex("index").Do(context.Background())
			es.OpenIndex("index").Do(context.Background())
			logger.Error("Number:%v select demotest es Scroll error:%v", number, err)
		}
		scrollid := demotest.ScrollId
		for {
			if len(demotest.Hits.Hits) > 0 {
				demodatas := make(map[string]interface{})
				for _, item := range demotest.Each(reflect.TypeOf(demodatas)) {
					demodata := DemoData{}
					number := item.(map[string]interface{})["number"]
					name := item.(map[string]interface{})["name"]
					title := item.(map[string]interface{})["title"]
					port := item.(map[string]interface{})["port"]
					alive := item.(map[string]interface{})["alive"]
					createTime := item.(map[string]interface{})["createTime"]
					lastTime := item.(map[string]interface{})["lastTime"]
					createtime, _ := time.ParseInLocation("2006-01-02T15:04:05", createTime.(string), time.Local)
					lasttime, _ := time.ParseInLocation("2006-01-02T15:04:05", lastTime.(string), time.Local)
					fingerprint := item.(map[string]interface{})["fingerprint"]
					waf := item.(map[string]interface{})["waf"]
					demodata.Number = number.(string)
					demodata.Name = name.(string)
					demodata.Tile = title.(string)
					switch port.(type) {
					case string:
						demodata.Port, _ = strconv.Atoi(fmt.Sprintf("%v", port))
					default:
						demodata.Port = int(port.(float64))
					}
					if alive != "" {
						demodata.Alive = int(alive.(float64))
					}
					demodata.CreatedAt = createtime
					demodata.UpdatedAt = lasttime
					if fingerprint != "" {
						switch fingerprint := fingerprint.(type) {
						case string:
							demodata.Fingerprint = []string{fingerprint}
						default:
							demodata.Fingerprint = TostringArray(fingerprint.([]interface{}))
						}
					}
					if waf != "" {
						switch waf := waf.(type) {
						case string:
							demodata.Waf = []string{waf}
						default:
							demodata.Waf = utils.TostringArray(waf.([]interface{}))
						}
					}
					fmt.Println("demoData Data Info:", demodata)
					demoData = append(demoData, demodata)
				}
			} else {
				break
			}
			if demotest.TotalHits() > int64(len(demotest.Hits.Hits)) {
				demotest, err = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
				if err != nil {
					logger.Error("demoData es Scroll ScrollId error:%v", err)
				}
				scrollid = demotest.ScrollId
			} else {
				break
			}
		}
		es.ClearScroll(scrollid)
		es.CloseIndex("index").Do(context.Background())
		es.OpenIndex("index").Do(context.Background())
	}
	fmt.Println("DemoTest Len", len(demoData))
}
```
