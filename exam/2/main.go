package main

import (
	"bufio"
	"fmt"
	"os"
)

var mounths = map[int]int{
	1: 31,
	2: 28,
	3: 31,
	4: 30,
	5: 31,
	6: 30,
	7: 31,
	8: 31,
	9: 30,
	10: 31,
	11: 30,
	12: 31,
}

func checkLeapYear(year int)  bool{
	devideFour := year % 4
	devideHunder := year % 100
	devideFourHunder := year % 400
	if devideFourHunder == 0 || (devideFour == 0 && devideHunder != 0) {
		return true
	}
	return false
}

func checkDate(day, mounth, year int) bool{
	maxDay := mounths[mounth]
	leapYear := checkLeapYear(year)
	if leapYear && mounth == 2 {
		maxDay += 1
	}

	if day > maxDay {
		return false
	}

	return true
}

func scanNumbers() []bool{
	input := bufio.NewReader(os.Stdin)
	var result []bool
	var t int
	

	fmt.Fscan(input, &t)
	
    for i := 0; i < t; i++ {
		var day int
		var mounth int
		var year int
       	fmt.Fscan(input, &day, &mounth, &year)
		result = append(result, checkDate(day, mounth, year))
	}

	return result
}

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