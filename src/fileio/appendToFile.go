package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "/Users/nakama/Downloads/a.txt"
	appendToFile(filename, "kodok\n\nkodok")
	fmt.Println("file appended successfully")
}

func appendToFile(filename, stringToBeAdded string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		f = createFile(filename)
	}
	_, err = fmt.Fprintln(f, stringToBeAdded)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func createFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}
