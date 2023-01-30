package main

import (
	"bufio"
	"fmt"
	"os"
)

//Считаем общую сумму по данным
func result(sumProduct map[int]int) int{
	var sum int
	for i, v := range sumProduct {
		discount := v/3
		sum += (v - discount) * i
	}
	return sum
}

//Формируем объединенный список товаров по ценам и их количество
func countSum(count int, price []int) map[int]int{
	sumProducts := make(map[int]int)
	for i := 0; i < count; i++ {
		if _, ok := sumProducts[price[i]]; !ok {
			sumProducts[price[i]] = 1
		} else {
			sumProducts[price[i]] += 1
		}
	}
	return sumProducts
}

//Сканируем данные из консоли и тут же обрабатываем их
//(неправильно, одна функция = одно действие)
func scanNumbers(gr map[int]int) []int{
	input := bufio.NewReader(os.Stdin)

	var sums []int
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
		sumProduct := countSum(count, prices)
		sum := result(sumProduct)
		sums = append(sums, sum)
	}
	return sums
}

//Проходимся по массиву результатов и печатаем его
func printResults(sums []int) {
	for _, v := range sums {
		fmt.Println(v)
	}
}

func main() {
	grProducts := make(map[int]int)
	//Сканируем, преобразуем данные и объединяем для печати на экран
	sums := scanNumbers(grProducts)
	//Печатаем результат
	printResults(sums)
}