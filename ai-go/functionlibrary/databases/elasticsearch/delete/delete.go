package main

import (
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

func Delete(es *elastic.Client) error {
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
