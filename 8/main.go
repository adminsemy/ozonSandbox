package main

import (
	"bufio"
	"fmt"
	"os"
)

//Все найденные гексограммы и их границы
var findGexogens = make(map[byte][]int)

//Потерянный элемент карты, который не на границе
type Lost struct {
	region byte
	position int
	borders []int
}

var losts []Lost

//Ищем потерящки. Если есть, то объединяем с основным регионом
func findLosts() {
	currentLen := len(losts)
	for k, l := range losts {
		if inRegion(l.region, l.position) {
			addNeighbors(l.region, l.borders)
			if len(losts) == 1 {
				losts = []Lost{}
				break
			}
			losts = append(losts[:k], losts[k + 1:]...)
			break
		}
	}
	changeLen := len(losts)
	//если что-то нашло, то пробегем еще раз
	if currentLen != changeLen {
		findLosts()
	}
}

//Проверяем, принадлежит ли региону текущий элекмент
//Если да, то убираем его из списка
func inRegion(region byte, currentPosition int) bool{
	
	for _, r := range findGexogens[region] {
		if r == currentPosition {
			return true
		}
		
	}

	return false
}

//Добавляем еще соседей к существующему списку
func addNeighbors(region byte, neighbors []int) {
	var addNeighbors []int
	next:
	for _, n := range neighbors {
		for _, r := range findGexogens[region] {
			if r == n {
				continue next
			}
		}
		addNeighbors = append(addNeighbors, n)
	}
	findGexogens[region] = append(findGexogens[region], addNeighbors...)
}

//Проверка коректности позиции. Не выходит ли за границы
func isCorrectPosition(position int, countSymbols int) bool{
	if position >= 0 && position <= countSymbols - 1 {
		return true
	}
	return false
}

//Поиск всех корректных соседей
//Каждая ячейка обозначается номером n где n = m * l + p
//где m - индекс строки в слайче, l - длина строк, p - позиция в строке
//currentPostion = n
//countSymbols = l
//capMap = m
func findNeighbors(currentPosition int, countSymbols int, capMap int) []int{
	var result []int
	numberString := currentPosition/countSymbols
	positionInString := currentPosition - (numberString * countSymbols)
	
	leftNeighbor := currentPosition - 2
	if isCorrectPosition(positionInString - 2, countSymbols) {
		result = append(result, leftNeighbor)
	}

	rightNeighbor := currentPosition + 2
	if isCorrectPosition(positionInString + 2, countSymbols) {
		result = append(result, rightNeighbor)
	}

	leftUpNeighbor := currentPosition - countSymbols - 1
	if isCorrectPosition(positionInString - 1, countSymbols) && numberString - 1 > 0 {
		result = append(result, leftUpNeighbor)
	}

	rightUpNeighbor := currentPosition - countSymbols + 1
	if isCorrectPosition(positionInString + 1, countSymbols) && numberString - 1 > 0 {
		result = append(result, rightUpNeighbor)
	}

	leftDownNeighbor := currentPosition + countSymbols - 1
	if isCorrectPosition(positionInString - 1, countSymbols) && numberString + 1 <= capMap {
		result = append(result, leftDownNeighbor)
	}

	rightDownNeighbor := currentPosition + countSymbols + 1
	if isCorrectPosition(positionInString + 1, countSymbols) && numberString + 1 <= capMap {
		result = append(result, rightDownNeighbor)
	}

	return result
}

//Проходим по всей карте и определяем, что она корректная
func checkRegions(data []string, lenString int) bool {
	findGexogens = make(map[byte][]int)
	losts = make([]Lost, 0)
	var answer bool
	for k, region := range data {
		n := k%2
		for i := n; i < len(region); i += 2 {
			currentPosition := (k + 1) * lenString + i
			neighbors := findNeighbors(currentPosition, lenString, len(data))
			if _, ok := findGexogens[region[i]]; ok {
				if inRegion(region[i], currentPosition) {
					addNeighbors(region[i], neighbors)
					continue
				}
				losts = append(losts, Lost{region[i], currentPosition, neighbors})
			} else {
				findGexogens[region[i]] = neighbors
			}
		}
		findLosts()
	}
	if len(losts) == 0 {
		answer = true
	}
	return answer
}


//Сканирование данных из консоли
func scanNumbers() []bool{
	input := bufio.NewReader(os.Stdin)
	var t int
	
	fmt.Fscan(input, &t)

	var result []bool

	for i := 0; i < t; i++ {
		var n int //количество строк
		var m int //количество символов
		var myMap []string
		fmt.Fscan(input, &n, &m)
		for y := 0; y < n; y++ {
			var row string
			fmt.Fscan(input, &row)
			myMap = append(myMap, row)
		}
		result = append(result, checkRegions(myMap, m))
	}
	return result
}

//Печатаем результат
func printResult(results []bool) {
	for _, r := range results {
		if r {
			fmt.Println("YES")
			continue
		}
		fmt.Println("NO")
	}
}

func main() {
	result := scanNumbers()
	printResult(result)
}