package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr1 := arr[0]
	arr = arr[1:]
	arr = append(arr, arr1)
	fmt.Println(arr)
}
