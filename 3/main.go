package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Преобразуем строку входных данных в слайс чисел
func split(str string) []int{
	var result []int
	s := strings.Split(str, " ")
	for _,v := range s {
		vInt, err := strconv.Atoi(v)
		if err == nil {
			result = append(result, vInt)
		}		
	}
	return result
}

//Формируем список нужных пар
func list (nDevelop int, arr []int) [][]int{
	var result [][]int
	couples := make(map[int]bool) //сохраняем выявленных разработчиков
	for i := 1; i <= nDevelop; i++ {
		if _, ok := couples[i]; ok {
			continue
		}
		var index int //индекс второго  разработчика
		//Результат самой минимальной разницы между уровнями разработчиков
		//Для каждой итерации по разработчику начинаем искать разницу заново
		var abs float64 
		couples[i] = true
		firstDevelop := arr[i-1]
		for i, secondDevelop := range arr {
			i = i + 1
			if _, ok := couples[i]; ok {
				continue
			}
			//Разница модуля числа между двумя разработчиками
			diff := math.Abs(float64(firstDevelop) - float64(secondDevelop))
			if diff == 0 {
				index = i
				break
			}
			if abs == 0 {
				abs = diff
				index = i
				continue
			}
			if diff >= abs {
				continue
			} 
			if diff < abs {
				index = i
				abs = diff
			}
		}
		couples[index] = true
		result = append(result, []int{i, index})
	}
	return result
}			


func scanNumbers() [][][]int{
	input := bufio.NewReader(os.Stdin)


	var sums [][][]int
	var t int
	var count int

	fmt.Fscan(input, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(input, &count)
		var p int
		var prices []int
		for y := 0; y < count; y++ {
			fmt.Fscan(input, &p)
			prices = append(prices, p)
		}
		result := list(count, prices)
		sums = append(sums, result)
	}
	return sums
}

//Печатаем результаты в консоль
func printResult(result [][][]int) {
	for _, res := range result {
		for _, r := range res {
			fmt.Println(r[0], r[1])
		}
		fmt.Println("")
	}
}

func main() {
	result := scanNumbers()
	printResult(result)
}