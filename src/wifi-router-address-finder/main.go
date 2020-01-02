package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func main() {
	pingPossibleWiFiRouterAddresses()
}

func pingPossibleWiFiRouterAddresses() {
	var addressList []string
	readFromFile(getPathToTextList(), func(address string) { addressList = append(addressList, address) })

	wg := sync.WaitGroup{}
	for _, address := range addressList {
		wg.Add(1)
		go pingAddress(address, &wg)
	}
	wg.Wait()
}

func pingAddress(address string, wg *sync.WaitGroup) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Head("http://" + address)
	if err != nil {
		//fmt.Println("failed to ping", address, err)
		wg.Done()
		return
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		fmt.Println(address, "responded with", http.StatusOK)
	}
	wg.Done()
}

func readFromFile(path string, cb func(string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cb(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getPathToTextList() (path string) {
	dir, _ := getPathToThisFile()
	path = dir + "/address-list.txt"
	return
}

func getPathToThisFile() (path string, err error) {
	_, path, _, _ = runtime.Caller(1)
	path = filepath.Dir(path)
	return
}
