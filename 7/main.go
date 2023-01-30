package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var data = make(map[int][]int)

//Проходим по карте пользователей иищем подходящих друзей
func findRecomend(countUsers int, result [][]int) {
	for i := 1; i <= countUsers; i++ {
		var userFriends = make(map[int]struct{})
		userFriends[i] = struct{}{}
		for _, f := range data[i] {
			userFriends[f] = struct{}{}
		}
		recomendFriends := make(map[int]int)	
		for _, f := range data[i] {
			findFriends(f, userFriends, recomendFriends)
		}
		result[i-1] = findMaxGeneralFriends(recomendFriends)
	}
}

//Проходимся по списку друга и возвращаем список рекомендуемых друзей
func findFriends(friend int, userFriends map[int]struct{}, recomendFriends map[int]int) {
	for _, f := range data[friend] {
		if _, ok := userFriends[f]; ok {
			continue
		}
		recomendFriends[f] += 1
	}
}

//Находим всех тех, у кого наибольшое количество общих друзей
func findMaxGeneralFriends(general map[int]int) []int {
	if len(general) == 0 {
		return []int{}
	}
	var max int
	var list []int
	for i, v := range general {
		if v == max {
			list = append(list, i)
		}
		if v > max {
			list = []int{i}
			max = v
		}
	}
	return list
}

//Сканирование данных из консоли
func scanNumbers() [][]int{
	input := bufio.NewReader(os.Stdin)
	var countUsers int
	var t int
	var user int
	var friend int
	
	fmt.Fscan(input, &countUsers, &t)

	result := make([][]int, countUsers)
	
	for i := 0; i < t; i++ {
		fmt.Fscan(input, &user, &friend)
		data[user] = append(data[user], friend)
		data[friend] = append(data[friend], user)
	}
	findRecomend(countUsers, result)

	return result
}

//Печатаем результат
func printResult(results [][]int) {
	for _, recomend := range results {
		if len(recomend) == 0 {
			fmt.Println(0)
			continue
		}		
		sort.Ints(recomend)
		rec := fmt.Sprint(recomend)
		rec = strings.Trim(rec, "[]")
		fmt.Println(rec)
	}
}

func main() {
	result := scanNumbers()
	printResult(result)	
}