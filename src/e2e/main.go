//Run with: go run src/e2e/main.go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {
	//do_Ping()
	do_Login("lucky.ong@gmail.com", "Tokopedia789&*(")
}

const URL_LENDINGAPP = "http://127.0.0.1:9003" //localhost
const _aaa = 1

func do_Ping() {
	resp, err := http.Get(URL_LENDINGAPP + "/ping")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}

func do_Login(email_phone, password string) {
	url := URL_LENDINGAPP + "/v1/login"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("email_phone", email_phone)
	_ = writer.WriteField("password", password)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Printf("\n%+v\n", res)
	fmt.Println(string(body))
}
