package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	elastic "gopkg.in/olivere/elastic.v5"
)

var elasticClient *elastic.Client

func Init() {
	httpClient := &http.Client{Transport: DefaultHeaderTransport{}.next}

	var err error
	elasticClient, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetDecoder(&Decoder{}),
		elastic.SetHttpClient(httpClient))

	fmt.Println(elasticClient)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Init()
	Term()
}

type Decoder struct{}

func (u *Decoder) Decode(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}
func (u *Decoder) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.ConfigFastest.Unmarshal(data, v)
}
func (u *Decoder) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.ConfigFastest.Marshal(v)
}

type DefaultHeaderTransport struct {
	N    int64
	next http.RoundTripper
}

func (tr *DefaultHeaderTransport) RoundTripper(r *http.Request) (*http.Response, error) {
	r.Header.Add("service", "search-microservice")
	if tr.next != nil {
		return tr.next.RoundTrip(r)
	}
	return http.DefaultTransport.RoundTrip(r)
}

func Get() {
	res, err := elasticClient.Get().Index("autocomplete_v1").Id("1").Do(context.TODO())

	if err != nil {
		fmt.Println(err)
		return
	}

	if res.Source == nil {
		fmt.Println(err)
		return
	}

	var data Product
	err = jsoniter.ConfigFastest.Unmarshal(*res.Source, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data.ProductName)
}

func Term() {
	searchService := elasticClient.Search().Index("autocomplete_v1").Timeout("1s")

	q := "samsung"
	query := elastic.NewTermQuery("name", q)

	searchService.Query(query)
	searchService.Size(3)
	searchService.From(0)

	ascending := false
	searchService.Sort("score", ascending)

	searchResult, err := searchService.Do(context.TODO())

	if err != nil {
		fmt.Println(err)
		return
	}

	var items []Product
	if searchResult != nil {
		if searchResult.Hits == nil || searchResult.Hits.Hits == nil || len(searchResult.Hits.Hits) == 0 {
			fmt.Println(len(searchResult.Hits.Hits))
			return
		}

		for _, hit := range searchResult.Hits.Hits {
			if hit.Source == nil {
				continue
			}
			var tmp Product
			if err := jsoniter.ConfigFastest.Unmarshal(*hit.Source, &tmp); err == nil {
				items = append(items, tmp)
			}
		}
	}

	fmt.Println(items)
}

type Product struct {
	CategoryRecommendation struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"category_recommendation"`
	ID          int64  `json:"id"`
	ProductName string `json:"name"`
	Score       int64  `json:"score"`
}
