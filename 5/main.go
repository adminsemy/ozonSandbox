package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Проверяем на корректность левю и правую части отрезка времени
func checkCorrectTimeRange(clockLeft, clockRight string) error {
	if clockLeft > clockRight {
		return errors.New("Левая часть больше, чем правая")
	}
	return nil
}

//Структура для хранения данных от пользоваталя
type Clock struct {
	left string
	right string
}

type Channel struct {
	Position int8
	Result bool
}

func changeCouples(couples []string, number int8, ch chan <- Channel) {
	timeRanges := make(map[int][]Clock)
	for _, v := range couples {
		couples, err := createCouple(v)
		if err != nil {
			ch <- Channel{number, false}
			return
		}
		leftHour, _ := strconv.Atoi(couples.left[0:2])
		rightHour, _ := strconv.Atoi(couples.right[0:2])
		for i:= leftHour; i <= rightHour; i++ {
			if err := changeTimeRanges(timeRanges[i], couples); err != nil {
				ch <- Channel{number, false}
				return
			}
		}
		for i := leftHour; i <= rightHour; i++ {
			timeRanges[i] = append(timeRanges[i], couples)
		}
	}
	ch <- Channel{number, true}
	return
}

//Проверка вхождения в другие отрезки времени
func changeTimeRanges(timeRanges []Clock, couple Clock) error {
	for _, tr := range timeRanges {
		if (couple.left >= tr.left && couple.left <= tr.right) {
			return errors.New("левая часть входит в диапазон")
		}
		if (couple.right >= tr.left && couple.right <= tr.right) {
			return errors.New("правая часть входит в диапазон")
		}
		if (tr.left >= couple.left && tr.left <= couple.right) {
			return errors.New("левая часть входит в диапазон")
		}
	}
	return nil
}


//Проверка соответствия необходимому диапазону отрезка времени
func timeCorrect(time string, min string, max string, clockRange string) error {
	if time < min || time >= max {
		return errors.New(clockRange + " не входят в необходимый диапазон")
	}
	return nil
}

//Проверяем время на корректные данные
func check(clock string) error{
	
	hour := clock[0:2]
		if err := timeCorrect(hour, "00", "24", "часы"); err != nil {
		return err
	}

	minute := clock[3:5]
	if err := timeCorrect(minute, "00", "60", "минуты"); err != nil {
		return err
	}
	second := clock[6:8]
	if err := timeCorrect(second, "00", "60", "секунды"); err != nil {
		return err
	}
	return nil
}

//Проверяем пары отрезка времени
func createCouple(input string) (Clock, error) {
	couple := split(input, "-")
	clockLeft := couple[0]
	errorOne := check(clockLeft)
	if errorOne != nil {
		return Clock{}, errorOne
	}
	clockRight := couple[1]
	errorTwo := check(clockRight)
	if errorTwo != nil {
		return Clock{}, errorTwo
	}
	if err := checkCorrectTimeRange(clockLeft, clockRight); err  != nil {
		return Clock{}, err
	}
	return Clock{clockLeft, clockRight}, nil
}

//Делим строку на части
func split(str string, split string) []string{
	s := strings.Split(str, split)
	return s
}

//Сканирование данных из консоли
func scanNumbers() []bool{
	input := bufio.NewReader(os.Stdin)
	dataResult := make(chan Channel)
	var t int8

	fmt.Fscan(input, &t)

	result := make([]bool, t)

	for i := int8(0); i < t; i++ {
		var count int16
		fmt.Fscan(input, &count)
		couples := make([]string, count)
		for y:=int16(0); y < count; y++ {
			var c string
			fmt.Fscan(input, &c)
			couples[y] = c
		}
		go changeCouples(couples, i, dataResult)
	}
	for i:= int8(0); i < t; i++ {
		ch := <-dataResult
		result[ch.Position] = ch.Result
	}
	return result
}

//Печатаем результаты в консоль
func printResult(result []bool) {
	for _, res := range result {
		switch res {
		case true:
			fmt.Println("YES")
		case false:
			fmt.Println("NO")
		}	
	}
}

func main() {
	result := scanNumbers()
	printResult(result)
}