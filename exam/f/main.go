package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"time"
)

var maxFive int
var minOne int

var (
	lastElementOne   int
	lastElementTwo   int
	lastElementThree int
	lastElementFour  int
	lastElementFive  int
)
var (
	countElementOne   int
	countElementTwo   int
	countElementThree int
	countElementFour  int
	countElementFive  int
)
var (
	lenOne   = [2]int{0, 4}
	lenTwo   = [2]int{5, 8}
	lenThree = [2]int{9, 11}
	lenFour  = [2]int{12, 13}
	lenFive  = [2]int{14, 14}
)

var elements []int

//Изменяем индекс всех вариантов отелей
func setIndexs(countElements int) {
	setCountElements(countElements)
	lenFive[1] = countElements - 1
	lenFive[0] = lenFive[1] - (countElementFive - 1)

	lenFour[1] = lenFive[0] - 1
	lenFour[0] = lenFour[1] - (countElementFour - 1)

	lenThree[1] = lenFour[0] - 1
	lenThree[0] = lenThree[1] - (countElementThree - 1)

	lenTwo[1] = lenThree[0] - 1
	lenTwo[0] = lenTwo[1] - (countElementTwo - 1)

	lenOne[1] = lenOne[0] + (countElementOne - 1)

	fmt.Println(countElementOne, countElementTwo, countElementThree, countElementFour, countElementFive)
}

//Находим максимуму 5-звездочных и минимум 1-звездочных
func setCountElements(countElements int) {
	lastElements := countElements % 5
	countElementThree = countElements / 5
	countElementOne = countElementThree + 2
	countElementTwo = countElementThree + 1
	countElementFour = countElementThree - 1
	countElementFive = countElementThree - 2
	if lastElements > 0 {
		countElementOne += 1
		lastElements--
	}
	if lastElements > 0 {
		countElementTwo += 1
		lastElements--
	}
	if lastElements > 0 {
		countElementThree += 1
		lastElements--
	}
	if lastElements > 0 {
		countElementFour += 1
	}
}

//Ищем все отели с максимальной оценкой
func findFiveElements(elements []int) error {
	//fmt.Println("Five", lenOne, lenTwo, lenThree, lenFour, lenFive)
	if lenFive[0] > lenFive[1] || lenFive[1] >= len(elements) {
		return errors.New("Достигнут предел возможного минимального значения")
	}
	for i := lenFive[0]; i <= lenFive[1]; i++ {
		if elements[i] != elements[i-1] {
			break
		}
		if lenFive[0]+1 > lenFive[1] {
			return errors.New("Достигнут предел возможного минимального значения")
		}
		lenFive[0] += 1
		lenFour[1] += 1
		countElementFive--
		countElementFour++
	}
	err := findFourElements(elements)
	if err != nil {
		return err
	}
	return nil
}

//Ищем все отели с оценкой 4
func findFourElements(elements []int) error {
	//fmt.Println("Four", lenOne, lenTwo, lenThree, lenFour, lenFive)
	for i := lenFour[0]; i <= lenFour[1]; i++ {
		if elements[i] != elements[i-1] {
			break
		}
		lenFour[0] += 1
		countElementFour--
		lenThree[1] += 1
		countElementThree++
		if lenFour[0] == lenFive[0] {
			lenFour[1] = lenFive[0]
			countElementFour++
			if lenFive[0]+1 > lenFive[1] {
				return errors.New("Больше увеличить набор отелей 5 звезд не получится")
			}
			lenFive[0] += 1
			countElementFive--
			break
		}
	}
	if countElementFour <= countElementFive {
		add := (countElementFive-countElementFour)/2 + 1
		lenFour[1] = lenFour[1] + add
		countElementFour += add
		lenFive[0] = lenFive[0] + add
		countElementFive -= add
		err := findFiveElements(elements)
		if err != nil {
			return err
		}
	}
	err := findThreeElements(elements)
	if err != nil {
		return err
	}
	return nil
}

//Ищем все отели с оценкой 3
func findThreeElements(elements []int) error {
	//fmt.Println("Three", lenOne, lenTwo, lenThree, lenFour, lenFive)
	for i := lenThree[0]; i <= lenThree[1]; i++ {
		if elements[i] != elements[i-1] {
			break
		}
		lenThree[0] += 1
		countElementThree--
		lenTwo[1] += 1
		countElementTwo++
		if lenThree[0] == lenFour[0] {
			lenThree[1] = lenFour[0]
			countElementThree++
			lenFour[0] += 1
			countElementFour--
			break
		}
	}
	if countElementThree <= countElementFour {
		add := (countElementFour-countElementThree)/2 + 1
		lenThree[1] += add
		countElementThree += add
		lenFour[0] += add
		countElementFour -= add
		err := findFourElements(elements)
		if err != nil {
			return err
		}
	}
	err := findTwoElements(elements)
	if err != nil {
		return err
	}
	return nil
}

//Ищем все отели с оценкой 2
func findTwoElements(elements []int) error {
	fmt.Println("Two", lenOne, lenTwo, lenThree, lenFour, lenFive)
	for i := lenTwo[0]; i <= lenTwo[1]; i++ {
		if elements[i] != elements[i-1] {
			break
		}
		lenTwo[0] += 1
		countElementTwo--
		lenOne[1] += 1
		countElementOne++
		if lenTwo[0] == lenThree[0] {
			lenTwo[1] = lenThree[0]
			countElementTwo++
			lenThree[0] += 1
			countElementThree--
			break
		}
	}
	if countElementTwo <= countElementThree {
		add := (countElementThree-countElementTwo)/2 + 1
		lenTwo[1] = lenTwo[1] + add
		countElementTwo += add
		lenThree[0] = lenThree[0] + add
		countElementThree -= add
		err := findThreeElements(elements)
		if err != nil {
			return err
		}
	}
	err := findOneElements(elements)
	if err != nil {
		return err
	}
	return nil
}

//Проверяем количество с 1 звездой и с 2 звездами
func findOneElements(elements []int) error {
	fmt.Println("One", lenOne, lenTwo, lenThree, lenFour, lenFive)
	if countElementOne <= countElementTwo {
		add := (countElementTwo-countElementOne)/2 + 1
		lenOne[1] = lenOne[1] + add
		countElementOne += add
		lenTwo[0] = lenTwo[0] + add
		countElementTwo -= add
		err := findTwoElements(elements)
		if err != nil {
			return err
		}
	}
	return nil
}

func stars(countElements int, elements []int) []int {
	var elementsSort []int
	elementsSort = append(elementsSort, elements...)
	sort.Ints(elementsSort)
	setIndexs(countElements)
	fmt.Println(elementsSort)
	err := findFiveElements(elementsSort)
	if err != nil {
		return wrongResult(countElements)
	}
	result := createResult(elements, elementsSort, countElements)
	fmt.Println(createResult(elementsSort, elementsSort, countElements))
	return result
}

func createResult(elements, elementsSort []int, countElements int) []int {
	result := make([]int, countElements)
	for i, e := range elements {
		switch {
		case e < elementsSort[lenTwo[0]]:
			result[i] = 1
		case e >= elementsSort[lenTwo[0]] && e < elementsSort[lenThree[0]]:
			result[i] = 2
		case e >= elementsSort[lenThree[0]] && e < elementsSort[lenFour[0]]:
			result[i] = 3
		case e >= elementsSort[lenFour[0]] && e < elementsSort[lenFive[0]]:
			result[i] = 4
		case e >= elementsSort[lenFive[0]]:
			result[i] = 5
		}
	}
	return result
}

func wrongResult(countElements int) []int {
	result := make([]int, countElements)
	for i := 0; i < countElements; i++ {
		result[i] = -1
	}
	return result
}

func scanNumbers() [][]int {
	input := bufio.NewReader(os.Stdin)
	var t int

	fmt.Fscan(input, &t)

	var result [][]int

	for i := 0; i < t; i++ {
		var countElements int //элемент
		var element int
		elements = []int{}

		fmt.Fscan(input, &countElements)
		for y := 0; y < countElements; y++ {
			fmt.Fscan(input, &element)
			elements = append(elements, element)
		}
		result = append(result, stars(countElements, elements))
	}
	return result
}

func printResult(result [][]int) {
	for _, res := range result {
		for _, r := range res {
			fmt.Print(r, " ")
		}
		fmt.Print("\n")
	}
}

func main() {
	for i := 0; i < 50; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(2 * time.Second)

}
