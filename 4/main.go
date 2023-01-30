package main

import (
	"bufio"
	"fmt"
	"os"
)

//Проверяем задачи на последовательно выполнения
func isYes(countDays int, task []int) bool {
	var yes = true
	//Запоминаем номер задачи и день (индекс) ее выполнения
	
	complitedTasks := make(map[int]int)
	for i, t := range task {
		currentDayask, ok := complitedTasks[t]
		if ok {
			//Если разница между индексами больше 1, то задача делалась непоследовательно
			if (i - currentDayask) == 1 {
				complitedTasks[t] = i
			} else {
				return false
			}
		}
		if !ok {
			complitedTasks[t] = i
		}
	}
	
	return yes
}

func scanNumbers() []bool{
	input := bufio.NewReader(os.Stdin)

	var result []bool
	var t int
	var count int

	fmt.Fscan(input, &t)

	for i := 0; i < t; i++ {
		fmt.Fscan(input, &count)
		var n int
		var tasks []int
		for y := 0; y < count; y++ {
			fmt.Fscan(input, &n)
			tasks = append(tasks, n)
		}
		result = append(result, isYes(count, tasks))
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