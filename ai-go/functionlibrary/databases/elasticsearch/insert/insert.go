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
	ID   string
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Data Data   `json:"data,omitempty"`
}

type Data struct {
	Ip   string `json:"ip,omitempty"`
	Port string `json:"port,omitempty"`
	Desc string `json:"desc,omitempty"`
}

func Insert(es *elastic.Client, datainfo []*DataInfo) error {
	req := es.Bulk().Index("intlligentdemo")
	for _, data := range datainfo {
		// 如果ID时int则需要转换成string（如：strconv.FormatUint(30, 10)）
		doc := elastic.NewBulkIndexRequest().Id(data.ID).Doc(data)
		req.Add(doc)
	}
	res, err := req.Do(context.Background())
	if err != nil {
		fmt.Println("添加数据失败", err)
		return err
	}
	if !res.Errors {
		return nil
	}
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

}
