package main

import (
	"bufio"
	"fmt"
	"os"
)

var letters = map[string]string{
	"00": "a",
	"100": "b",
	"101": "c",
	"11": "d",
}

func decodeNumbers(numbers string) string{
	var result string
	var setNumbers string
	for _, n := range numbers {
		setNumbers += fmt.Sprint(string(n))
		letter, ok := letters[setNumbers]
		if ok {
			result += letter
			setNumbers = ""
		}
	}
	return result
}

func scanNumbers() []string{
	input := bufio.NewReader(os.Stdin)
	var result []string
	var t int
	

	fmt.Fscan(input, &t)
	
    for i := 0; i < t; i++ {
		var setLetters string 
       	fmt.Fscan(input, &setLetters)
		result = append(result, decodeNumbers(setLetters))
	}

	return result
}

func printResult(result []string) {
	for _, res := range result {
		fmt.Println(res)
	}
}

func main() {
	result := scanNumbers()
	printResult(result)

}