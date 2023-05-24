package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/olivere/elastic/v7"
)

type Elastic struct {
	ElasticUrl  string
	ElasticUser string
	ElasticPass string
}

func (e *Elastic) ESConn() (*elastic.Client, error) {
	httpClient := &http.Client{}
	client, err := elastic.NewClient(
		elastic.SetHttpClient(httpClient),
		elastic.SetSniff(false),
		elastic.SetURL(e.ElasticUrl),
		elastic.SetBasicAuth(e.ElasticUser, e.ElasticPass),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type DataInfo struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Data Data   `json:"data,omitempty"`
}

type Data struct {
	Ip   string `json:"ip,omitempty"`
	Port string `json:"port,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type Indexdata struct {
	Index    string
	Id       string
	Mapdata  string
	DataInfo DataInfo
}

func (i *Indexdata) CreateIndex(es *elastic.Client) error {
	i.Index = "democreat"
	// map 格式一定要对，不然创建失败
	i.Mapdata = `{
	"mappings":{
		"properties":{
			"name": 	{"type":	"keyword"},
			"age": 		{"type":	"long"},
			"ip": 		{"type":	"text"},
			"port": 	{"type":	"long"},
			"desc": 	{"type":	"text"}
			}
		}
	}`
	ctx := context.Background()
	// 判断索引是否存在
	exists, err := es.IndexExists(i.Index).Do(ctx)
	if err != nil {
		return err
	}
	// DataInfo需转字符串
	if !exists {
		_, err := es.CreateIndex(i.Index).Body(i.Mapdata).Do(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// 结构体方式创建数据
func (i *Indexdata) CreateStruct(es *elastic.Client) error {
	data := Data{
		Ip:   "192.168.0.1",
		Port: "8080",
		Desc: "hello world!",
	}
	datainfo := DataInfo{
		Name: "pan",
		Age:  25,
		Data: data,
	}
	i.DataInfo = datainfo
	// 索引不能使用这种了些 intlligent-Demo es创建索引要求都是小写
	// Type()没有会报错，但是es 8.0.0 以后完全不支持，8.0版本之后去掉Type()试一下
	// 我的版本是7.2.0
	put, err := es.Index().Index(i.Index).Id(i.Id).BodyJson(i.DataInfo).Do(context.Background())
	if err != nil {
		fmt.Printf("create err:%v", err)
		return err
	}
	fmt.Printf("index id:%v, index:%v, type:%v", put.Id, put.Index, put.Type)
	return nil
}

// 字符串方式创建数据
func (i *Indexdata) CreateString(es *elastic.Client) error {
	mapdata := `{"name":"pan","age":26,"data":{"ip":"192.168.0.1","port":"8080","desc":"hello world"}}`
	i.Mapdata = mapdata
	i.Index = "createdemo"
	i.Id = "3"
	put, err := es.Index().Index(i.Index).Id(i.Id).BodyJson(i.Mapdata).Do(context.Background())
	if err != nil {
		fmt.Printf("createString err:%v", err)
		return err
	}
	fmt.Printf("createString successful:%v", put)
	return nil
}

func main() {
	var es = &Elastic{}
	es.ElasticUrl = "http://ip:port"
	es.ElasticUser = "username"
	es.ElasticPass = "password"
	ES, err := es.ESConn()
	if err != nil {
		fmt.Println("连接es 失败:", err)
	}
	fmt.Println("连接es 成功：", ES)

	var mapdata = Indexdata{}
	// // 创建索引
	// err = mapdata.CreateIndex(ES)
	// if err != nil {
	// 	fmt.Printf("create index err:%v", err)
	// }
	// // 结构体格式创建索引
	// err = mapdata.CreateStruct(ES)
	// if err != nil {
	// 	fmt.Println("创建失败")
	// }
	// 字符串格式创建索引
	err = mapdata.CreateString(ES)
	if err != nil {
		fmt.Printf("createstring err:%v", err)
	}
}
