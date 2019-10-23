package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

var elasticSearchURL = "http://localhost:9200"

func main() {
	populateElasticSearch(elasticSearchURL)
	printIndices(elasticSearchURL)

	fmt.Println("Successfully added 1000 documents to index: bank")
}

func populateElasticSearch(elasticSearchURL string) {
	url := elasticSearchURL + "/_bulk?pretty&refresh"
	file := getData()
	resp, err := http.Post(url, "application/json", file)
	if err != nil {
		fmt.Println("Error POST:", err)
		panic(err)
	}
	defer resp.Body.Close()
	defer file.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Populating result:\n", string(body))

	return
}

func printIndices(elasticSearchURL string) {
	resp, err := http.Get(elasticSearchURL + "/_cat/indices?v")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func getData() (file *os.File) {
	path := getPathToData()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	return
}

func getPathToData() (pathToJSON string) {
	//default to accounts.json
	pathToPlayground, _ := getPathToPlayground()
	pathToJSON = pathToPlayground + "/accounts-data.txt"
	return
}

func getPathToPlayground() (path string, err error) {
	_, path, _, _ = runtime.Caller(1)
	path = filepath.Dir(path)
	return
}
