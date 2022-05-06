package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type AdderJsonReq struct {
	Numbers *[]int `json:"numbers,omitempty"`
}

func main() {
	// get CLI arguments
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(`{"error": "please supply a json string as argument"}`)
		return
	}
	var jsonReq AdderJsonReq
	err := json.Unmarshal([]byte(args[1]), &jsonReq)
	if err != nil {
		fmt.Printf(`{"error": "%v"}`, err)
		return
	}
	if jsonReq.Numbers == nil {
		fmt.Printf(`{"error": "please supply numbers as field in the json"}`)
		return
	}
	numbers := *jsonReq.Numbers

	sum := 0
	for _, v := range numbers {
		sum += v
	}

	fmt.Printf(`{"type": "%T", "value": %v}`, sum, sum)
}
