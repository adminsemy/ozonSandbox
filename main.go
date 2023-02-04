package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr = append(arr[:8], arr[9:]...)
	fmt.Println(arr)
}
