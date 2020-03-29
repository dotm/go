package main

import (
	"os"
	"path"
	"path/filepath"
)

func main() {
	// get CLI arguments
	args := os.Args
	firstArg := args[1]  //full path of the file to be renamed
	secondArg := args[2] //new name of the files

	dir := filepath.Dir(firstArg)
	newName := path.Join(dir, secondArg)
	os.Rename(firstArg, newName)
}
