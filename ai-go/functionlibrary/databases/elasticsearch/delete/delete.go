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

func Deleteindex(es *elastic.Client, index string) error {
	res, err := es.DeleteIndex(index).Do(context.Background())
	if err != nil {
		fmt.Printf("delete index err:%v", err)
		return err
	}
	fmt.Printf("delete index successful:%v", res)
	return nil
}

func DeleteData(es *elastic.Client, index string, ID string) error {
	res, err := es.Delete().Index(index).Id(ID).Do(context.Background())
	if err != nil {
		fmt.Printf("delete data err:%v", err)
		return err
	}
	fmt.Printf("delete successful:%v", res)
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
	err = Deleteindex(ES, "index")
	if err != nil {
		fmt.Printf("delete fail :%v", err)
	}
}
