package main

import (
	"bufio"
	"fmt"
	"os"
)

func sum(a int, b int) int {
	return a + b
}

func createOutput(data [][]int) {
	for _, v := range data {
		fmt.Println(sum(v[0], v[1]))
	}
}

func scanNumbers(numberCoups *[][]int) {
	in := bufio.NewReader(os.Stdin)
	var a int //переменная a
	var b int //переменная b
	var t int //количество итераций считывания данных из консоли
	
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(in, &a, &b)
		*numberCoups = append(*numberCoups, []int{a, b})
	}
}

func main() {
	//Данные для подсчета 
	var numberCoups [][]int

	scanNumbers(&numberCoups)	
	createOutput(numberCoups)
}