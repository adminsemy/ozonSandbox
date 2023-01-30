package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//Все найденные гексограммы и их границы
var findGexogens = make(map[int][]int)

//Вся карта 
var myMap []string

//Элемент, откуда начинается движение фишки
var firstElement int

//Данные в карте. '*' - ход был, '.' - хода не было
var (
	star = '*'
	dot = '.'
)

//Добавляем соседей к элементу
func addNeighbors(currentPosition, neighbor int) {
	findGexogens[currentPosition] = append(findGexogens[currentPosition], neighbor)
}

//Позиция на карте элемента
func position(position, lenMap, lenString int ) (int, int, error){
	if position < 0 {
		return 0, 0, errors.New("позиция ниже нуля")
	}
	if position > lenMap * lenString {
		return 0, 0, errors.New("вышли за пределы карты")
	}
	indexMap := position / lenString
	indexString := position % lenString - 1
	if position % lenString == 0 {
		indexMap -= 1
		indexString = lenString - 1
	}
	if position == lenString {
		indexMap = 0
	}
	if indexMap < 0 || indexString < 0 {
		return 0, 0, errors.New("индекс карты меньше 0")
	}
	return indexMap, indexString, nil
}

//Поиск всех корректных соседей
func findNeighbors(currentPosition int, lenMap int, lenString int) {
	i, _, _ := position(currentPosition, lenMap, lenString)
	
	rightNeightbor := currentPosition + 1
	iRight, sRight, err := position(rightNeightbor, lenMap, lenString)
	if err == nil {
		if i == iRight && myMap[iRight][sRight] == byte(star){
			addNeighbors(currentPosition, rightNeightbor)
		}	
	}

	leftNeightbor := currentPosition - 1
	iLeft, sLeft, err := position(leftNeightbor, lenMap, lenString) 
	if err == nil {
		if i == iLeft && myMap[iLeft][sLeft] == byte(star){
			addNeighbors(currentPosition, leftNeightbor)
		}
	}

	upNeightbor := currentPosition - lenString
	iUp, sUp, err := position(upNeightbor, lenMap, lenString)
	if err == nil {
		if i - 1 == iUp && myMap[iUp][sUp] == byte(star){
			addNeighbors(currentPosition, upNeightbor)
		}
	}

	downNeightbor := currentPosition + lenString
	iDown, sDown, err := position(downNeightbor, lenMap, lenString)
	if err == nil {		
		if i + 1 == iDown && myMap[iDown][sDown] == byte(star){
			addNeighbors(currentPosition, downNeightbor)
		}
	}
	pointEntrance(currentPosition)
}

//Устанавливаем точку входя для фишки
func pointEntrance(currentPosition int) {
	countNeighbors := len(findGexogens[currentPosition])
	if countNeighbors == 1 {
		firstElement = currentPosition
	}
}

//Проходим по сформированной карте и возвращаем результат
func move(lenMap, lenString int) string{
	var move string
	var	lastPoint int
	pointEnter := firstElement
	for i := 0; i < len(findGexogens); i++{
		var nextPointer int
		findElement := findGexogens[pointEnter]
		firstElement := findElement[0]
		if len(findElement) == 1 {
			move += direction(pointEnter, firstElement, lenMap, lenString)
			lastPoint, pointEnter = pointEnter, findElement[0]
			continue
		}
		second := findElement[1]
		if lastPoint == firstElement {
			nextPointer = second
		} else {
			nextPointer = firstElement
		}
		move += direction(pointEnter, nextPointer, lenMap, lenString)
		lastPoint, pointEnter = pointEnter, nextPointer
	}
	return move[:len(move) - 1]
}

func direction(pointEnter, nextPointer, lenMap, lenString int) string{
	iEnter, sEnter, _ := position(pointEnter, lenMap,lenString)
	iPointer, sPointer, _ := position(nextPointer, lenMap, lenString)
	if iEnter < iPointer {
		return "D"
	}
	if iEnter > iPointer {
		return "U"
	}
	if sEnter < sPointer {
		return "R"
	}
	if sEnter > sPointer {
		return "L"
	}
	return ""
}

//Проходимся по карте и формирует ответ
func checkMap(lenMap, lenString int) string{
	for y, str := range myMap {
		for i, s := range str {
			if s == dot {
				continue
			}
			currentPosition := (i + 1) + lenString * y
			findNeighbors(currentPosition, lenMap, lenString)
		}
	}
	return move(lenMap, lenString)

}


//Сканирование данных из консоли
func scanNumbers() []string{
	input := bufio.NewReader(os.Stdin)
	var t int
	
	fmt.Fscan(input, &t)

	var result []string

	for i := 0; i < t; i++ {
		var n int //количество строк
		var m int //количество символов
		findGexogens = make(map[int][]int)
		myMap = []string{}
		firstElement = 0
		
		fmt.Fscan(input, &n, &m)
		for y := 0; y < n; y++ {
			var row string
			fmt.Fscan(input, &row)
			myMap = append(myMap, row)
		}
		result = append(result, checkMap(n, m))
	}
	return result
}

func printResult(result []string) {
	for _, r := range result {
		fmt.Println(r)
	}
}

func main() {
	result := scanNumbers()
	printResult(result)
}