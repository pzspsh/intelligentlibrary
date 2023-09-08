### ElasticSearch Go
```go
/*
@File   : task.go
@Author : pan
@Time   : 2023-09-08 12:52:53
*/
package aps

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

type ApsTaskArgs struct {
	Id        int64          `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string         `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Name      string         `gorm:"column:name;type:varchar(255);default:not null"`
	SesArgs   datatypes.JSON `gorm:"column:ses_args;serializer:json"`
	FisArgs   datatypes.JSON `gorm:"column:fis_args;serializer:json"`
	VvsArgs   datatypes.JSON `gorm:"column:vvs_args;serializer:json"`
	Target    []string       `gorm:"column:target;type:text[];default:null"`
	Status    int            `gorm:"column:status;default:null"`
	Open      int            `gorm:"column:open;default:null"`
	Level     int            `gorm:"column:level;default:null"`
	TaskType  int            `gorm:"column:task_type;default:null"`
	DwNumber  string         `gorm:"column:dw_number;type:varchar(255);default:not null"`
	DwName    string         `gorm:"column:dw_name;type:varchar(255);default:not null"`
	XmNumber  string         `gorm:"column:xm_number;type:varchar(255);default:not null"`
	XmName    string         `gorm:"column:xm_name;type:varchar(255);default:not null"`
	Owner     string         `gorm:"column:owner;type:varchar(255);default:not null"`
	CreatedAt time.Time      `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
	// CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	// UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;default:CURRENT_TIMESTAMP on update current_timestamp" json:"update_at,omitempty"`
}

type ApsTaskControl struct {
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number string `gorm:"column:number;type:varchar(255);unique_index;default:not null"`
	Status int    `gorm:"column:status;default:null"`
}

type ApsTaskProgress struct {
	Id        int64     `gorm:"column:id,primary_key;AUTO_INCREMENT;comment:自增编号"`
	Number    string    `gorm:"column:number;type:varchar(255);unique_index;not null"`
	Total     int       `gorm:"column:total;type:varchar(255);default:null"`
	Success   int       `gorm:"column:success;type:varchar(255);default:null"`
	Failure   int       `gorm:"column:failure;type:varchar(255);default:null"`
	UsedTime  string    `gorm:"column:usedtime;type:varchar(255);default:null"`
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;autoUpdateTime" json:"update_at,omitempty"`
}

type ApsTaskStatistic struct {
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

func (ApsTaskArgs) TableName() string {
	return "aps_task_args"
}

func (ApsTaskControl) TableName() string {
	return "aps_task_control"
}

func (ApsTaskProgress) TableName() string {
	return "aps_task_progress"
}

func (ApsTaskStatistic) TableName() string {
	return "aps_task_statistic"
}

func SelectTask(es *elastic.Client, db *gorm.DB) *[]ApsTaskArgs {
	// res, err := es.Search("tam-task").TrackTotalHits(true).Do(context.Background())
	// res, err := es.Search("tam-task").SearchAfter().TrackTotalHits(true).Do(context.Background())
	// res, err := es.Search("tam-task").From(1000).Size(10).Do(context.Background())
	// res, err := es.Scroll("tam-task").Do(context.Background())
	aftc := []aft.AftCompany{}
	companys := db.Find(&aftc)
	if companys.Error != nil {
		logger.Error("postgres db Find error:%v", companys.Error)
		return nil
	}
	for _, aft := range aftc {
		termquery := elastic.NewTermQuery("dwNumber", aft.Number)
		taskes, err := es.Scroll("tam-task").Query(termquery).Do(context.Background())
		if err != nil {
			logger.Error("aps task args search error:%v", err)
		}
		scrollid := taskes.ScrollId
		for {
			if len(taskes.Hits.Hits) > 0 {
				taskdata := make(map[string]interface{})
				for _, item := range taskes.Each(reflect.TypeOf(taskdata)) {
					number := item.(map[string]interface{})["number"]
					fmt.Println(number)
				}
			} else {
				break
			}
			taskes, _ = es.Scroll("tam-task").Query(termquery).ScrollId(scrollid).Do(context.Background())
			scrollid = taskes.ScrollId
		}
	}
	res, err := es.Scroll("tam-task").Do(context.Background())
	if err != nil {
		logger.Error("aps task args search error:%v", err)
	}
	// fmt.Println(res.TotalHits())
	// totallen := res.TotalHits()
	scrollid := res.ScrollId
	// task := make([]string, 0, res.TotalHits())
	// fmt.Println("totalhits", totallen)
	// fmt.Println(len(res.Hits.Hits))
	// fmt.Println("scrollid", scrollid)
	task := []string{}
	for {
		if len(res.Hits.Hits) > 0 {
			// for _, hit := range res.Hits.Hits {
			// 	fmt.Println(string(hit.Source))
			// task = append(task, string(hit.Source))
			// }
			taskdata := make(map[string]interface{})
			for _, item := range res.Each(reflect.TypeOf(taskdata)) {
				number := item.(map[string]interface{})["number"]
				// number := item.(map[string]interface{})["number"]
				// dwNumber := item.(map[string]interface{})["dwNumber"]
				// sesCs := item.(map[string]interface{})["sesCs"]
				// createTime := item.(map[string]interface{})["createTime"]
				// level := item.(map[string]interface{})["level"]
				// lastTime := item.(map[string]interface{})["lastTime"]
				// owner := item.(map[string]interface{})["owner"]
				// sesSl := item.(map[string]interface{})["sesSl"]
				// fisCs := item.(map[string]interface{})["fisCs"]
				// fisSl := item.(map[string]interface{})["fisSl"]
				// vvsCs := item.(map[string]interface{})["vvsCs"]
				// sescs := sesCs.(map[string]interface{})
				// fmt.Println(sescs["flag"])
				fmt.Println(number)
			}
		} else {
			break
		}
		res, _ = es.Scroll("tam-task").ScrollId(scrollid).Do(context.Background())
		scrollid = res.ScrollId
	}
	fmt.Println("AAAAAAAAAAAAAAAAAAAAA", task)
	data := &[]ApsTaskArgs{}
	return data
}

func (a *ApsTaskArgs) Insert() {

}

func (a *ApsTaskArgs) Update() {

}

func (a *ApsTaskControl) Insert() {

}

func (a *ApsTaskControl) Update() {

}

func (a *ApsTaskProgress) Insert() {

}

func (a *ApsTaskProgress) Update() {

}

func (a *ApsTaskStatistic) Insert() {

}

func (a *ApsTaskStatistic) Update() {

}

```