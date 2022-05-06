package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type AdderJsonReq struct {
	Numbers *[]int `json:"numbers,omitempty"`
}

func main() {
	req := AdderJsonReq{
		Numbers: &[]int{1, 2, 3},
	}
	jsonReq, err := json.Marshal(req) //`{"numbers":[10,20,3]}`
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonReq))

	evalCodeFilePath := "./dynamic-code/adder.go"
	cmd := exec.Command("go", "run", evalCodeFilePath, string(jsonReq))
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
		fmt.Println(stderr.String())
		return
	}
	fmt.Println(out.String())
}
