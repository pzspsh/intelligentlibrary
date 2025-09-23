package main

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/olivere/elastic/v7"
)

func ESConn(url, user, pass string) (*elastic.Client, error) {
	httpClient := &http.Client{}
	client, err := elastic.NewClient(
		elastic.SetHttpClient(httpClient),
		elastic.SetSniff(false),
		elastic.SetURL(url),
		elastic.SetBasicAuth(user, pass),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SelectAll(es *elastic.Client) {
	vulnes, err := es.Scroll("index").TrackTotalHits(true).Scroll("10m").Do(context.Background())
	if err != nil {
		fmt.Println("scroll error:", err)
	}
	fmt.Println(vulnes.TotalHits())
	var dataAll []map[string]any
	scrollid := vulnes.ScrollId
	for {
		if len(vulnes.Hits.Hits) > 0 {
			vulndata := make(map[string]any)
			for _, item := range vulnes.Each(reflect.TypeOf(vulndata)) {
				fmt.Println(item)
				dataAll = append(dataAll, item.(map[string]any))
			}
		} else {
			break
		}
		if vulnes.TotalHits() > int64(len(vulnes.Hits.Hits)) {
			vulnes, err = es.Scroll("index").ScrollId(scrollid).Do(context.Background())
			if err != nil {
				fmt.Println("scrollId error:", err)
			}
			scrollid = vulnes.ScrollId
		} else {
			break
		}
	}
	fmt.Println(len(dataAll))
}

func main() {
	es, err := ESConn("elastic-url", "username", "password")
	if err != nil {
		fmt.Println("es conn error:", err)
	}
	SelectAll(es)
}
