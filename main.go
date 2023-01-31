package main

import "fmt"

func main() {
	var iStart uint64 = 97
	var iEnd uint64 = 100
	var iCenter uint64
	iCenter = iEnd - (iEnd-iStart)/2
	fmt.Println(iCenter)
}
