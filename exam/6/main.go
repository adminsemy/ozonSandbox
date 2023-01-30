package main

import (
	"bufio"
	"fmt"
	"os"
)

func findCouples(ring map[int][2]int, rows, firstElement int) [][2]int{
	result := make([][2]int, rows / 2)
	var prevElement int
	var nextElement int

	//Определяем половину списка для формирования ответа
	half := rows / 2
	for i := 0; i < rows; i++ {
		elem := ring[firstElement]
		nextElement = elem[0]
		if prevElement == nextElement {
			nextElement = elem[1]
		}		
		if half > i {
			result[i][0] = firstElement
		} else {
			result[i - half][1] = firstElement
		}
		prevElement = firstElement
		firstElement = nextElement

	}
	return result
}

//Сканирование данных из консоли
func scanNumbers() [][][2]int{
	input := bufio.NewReader(os.Stdin)
	var t int
	
	fmt.Fscan(input, &t)

	var result [][][2]int

	for i := 0; i < t; i++ {
		var n int //количество строк
		var element int //элемент
		var firstNeightbor int //первый сосед
		var secondNeightbor int //второй сосед
		
		ring := make(map[int][2]int)

		firstElement := 0

		fmt.Fscan(input, &n)
		for y := 0; y < n; y++ {
			fmt.Fscan(input, &element, &firstNeightbor, &secondNeightbor)
			if y == 0 {
				firstElement = element
			}
			ring[element] = [2]int{firstNeightbor, secondNeightbor}
		}
		data := findCouples(ring, n, firstElement)
		result = append(result, data)
	}
	return result
}

func printResult(result [][][2]int) {
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