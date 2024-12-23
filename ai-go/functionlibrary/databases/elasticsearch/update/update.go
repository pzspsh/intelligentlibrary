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

// 修改  Upsert()表示不存在则插入,去掉Upsert()的话就指修改
func Update(es *elastic.Client, index string, ID string, data map[string]interface{}) error {
	res, err := es.Update().Index(index).Id(ID).Doc(data).Upsert(data).Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Printf("update data err:%v", err)
		return err
	}
	fmt.Printf("update successful %v", res)
	return nil
}

// 批量修改 Upsert()表示不存在则插入
func UpdateBulk(es *elastic.Client, index string, ids []string, docs []interface{}) error {
	buld := es.Bulk().Index(index)
	for i, id := range ids {
		doc := elastic.NewBulkUpdateRequest().Id(id).Doc(docs[i]).Upsert(docs[i])
		buld.Add(doc)
	}
	res, err := buld.Do(context.Background())
	if err != nil {
		fmt.Printf("updatebuld err:%v", err)
		return err
	}
	fmt.Printf("updatebulk successful:%v", res)
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
	data := map[string]interface{}{"age": "100"}
	err = Update(ES, "createdemo", "2", data)
	if err != nil {
		fmt.Printf("update fail:%v", err)
	}
}
