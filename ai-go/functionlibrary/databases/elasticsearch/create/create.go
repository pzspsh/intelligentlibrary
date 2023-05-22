package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/olivere/elastic"
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
	Index   string
	Mapdata string
}

func (i *Indexdata) CreateIndex(es *elastic.Client) error {
	// 判断索引是否存在
	exists, err := es.IndexExists(i.Index).Do(context.Background())
	if err != nil {
		return err
	}
	// DataInfo需转字符串
	if !exists {
		_, err := es.CreateIndex(i.Index).Body(i.Mapdata).Do(context.Background())
		if err != nil {
			return err
		}
	}
	return nil
}

// 结构体方式创建数据
func Createstruct(es *elastic.Client) error {
	data := Data{
		Ip:   "192.158.0.1",
		Port: "8080",
		Desc: "hello world!",
	}
	datainfo := DataInfo{
		Name: "pan",
		Age:  25,
		Data: data,
	}
	// 索引不能使用这种了些 intlligent-demo
	put, err := es.Index().Index("intlligentdemo").Type("datainfo").Id("1").BodyJson(datainfo).Do(context.Background())
	if err != nil {
		fmt.Printf("create err:%v", err)
		return err
	}
	fmt.Printf("index id:%v, index:%v, type:%v", put.Id, put.Index, put.Type)
	return nil
}

// 字符串方式创建数据
func Createstring(es *elastic.Client) error {
	return nil
}

func main() {
	var es = &Elastic{}
	es.ElasticUrl = "http://10.0.35.74:9200"
	es.ElasticUser = "elastic"
	es.ElasticPass = "techtech"
	ES, err := es.ESConn()
	if err != nil {
		fmt.Println("连接es 失败:", err)
	}
	fmt.Println("连接es 成功：", ES)
	err = Createstruct(ES)
	if err != nil {
		fmt.Println("创建失败")
	}

}
