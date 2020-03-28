//run this file to generate code
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	//improvements:
	//use linter after writing to file
	//use golang template with a valid template.go file
	//  where the type is a concrete type such as int64
	writeToFile(generateCode())
}

func generateCode() string {
	typeArray := [...]string{"int", "string"}
	code := "//This file is generated from: " + getPathToThisFile() + "\n"
	code += "//Do NOT modify this file directly!\n"
	code += "package main\n"

	for _, typeName := range typeArray {
		code += fmt.Sprintf(`
func concat%v(a, b %v) string {
	return string(a) + string(b)
}
`, strings.Title(typeName), typeName)
	}
	return code
}

func writeToFile(code string) {
	f, err := os.Create(getPathToGeneratedFile())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.WriteString(code)
}

func getPathToGeneratedFile() (path string) {
	dir, _ := getDirPathToThisFile()
	path = getParentDirectory(dir) + "/generated.go"
	return
}

func getParentDirectory(dir string) (parentDir string) {
	parentDir = filepath.Dir(dir)
	return
}

func getDirPathToThisFile() (dir string, err error) {
	dir = filepath.Dir(getPathToThisFile())
	return
}

func getPathToThisFile() (path string) {
	_, path, _, _ = runtime.Caller(1)
	return
}
