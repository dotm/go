package main

import (
	"os"
	"time"

	"learn-go-with-tests/go-fundamentals/math/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
