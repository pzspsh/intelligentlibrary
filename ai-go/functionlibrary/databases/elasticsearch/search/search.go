package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

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

func Search(es *elastic.Client, index, id, _type string) error {
	search, err := es.Get().Index(index).Id(id).Do(context.Background())
	if err != nil {
		fmt.Printf("search err:%v", err)
		return err
	}
	if search.Found {
		fmt.Printf("search id:%v\n", search.Id)
		fmt.Printf("search index:%v\n", search.Index)
		fmt.Printf("search source:%v", string(search.Source))
	}
	return nil
}

func SearchQuery(es *elastic.Client, index, _type string) error {
	// 获取index的所有数据
	res, err := es.Search(index).Do(context.Background())
	if err != nil {
		fmt.Printf("search query err:%v\n", err)
		return err
	}
	fmt.Println(res)
	// printdata(res, err)
	// for _,item := range res.Each(reflect.TypeOf())
	// 相等匹配
	// q := elastic.NewQueryStringQuery("age:99")
	// res, err = es.Search(index).Type(_type).Query(q).Do(context.Background())
	// if err != nil {
	// 	fmt.Printf("search query err:%v", err)
	// 	return err
	// }
	// fmt.Println(res)
	return nil
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

// 打印查询到的数据
func Printdata(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	var typ DataInfo
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(DataInfo)
		fmt.Printf("%#v\n", t)
	}
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
	err = Search(ES, "createdemo", "2", "createdemo")
	if err != nil {
		fmt.Printf("search fail:%v", err)
	}
	// err = SearchQuery(ES, "createdemo", "createdemo")
	// if err != nil {
	// 	fmt.Printf("search fail:%v", err)
	// }
}
