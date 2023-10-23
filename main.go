package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(5)
	n := runtime.GOMAXPROCS(0)

	fmt.Println(runtime.NumCPU(), n)
}
