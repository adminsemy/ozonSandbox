package main

import (
	"bufio"
	"fmt"
	"os"
)


func findMaxLenth(requests []int) int{
	var result int
	switch len(requests) {
		case 1:
			return 1
		case 2:
			return 2
	}
	firstClient := requests[0]
	secondClient := requests[1]
	tempSequence := []int{firstClient, secondClient}
	var consistent []int
	for i := 2; i < len(requests); i++ {
		c := requests[i]
		if c == firstClient && firstClient == secondClient {
			tempSequence = append(tempSequence, c)
			continue
		}
		if c != firstClient && firstClient == secondClient {
			secondClient = c
			tempSequence = append(tempSequence, c)
			continue
		}
		if c == firstClient && c != secondClient {
			if len(consistent) == 0 || consistent[0] != c{
				if c == requests[i-1] {
					consistent = []int{c}
				} else {
					consistent = []int{}
				}
			}
			consistent = append(consistent, c)
			tempSequence = append(tempSequence, c)
			continue
		}
		if c != firstClient && c == secondClient {
			if len(consistent) == 0 || consistent[0] != c{
				if c == requests[i-1] {
					consistent = []int{c}
				} else {
					consistent = []int{}
				}
			}
			consistent = append(consistent, c)
			tempSequence = append(tempSequence, c)
			continue
		}
		if c != firstClient && c != secondClient {
			maxLengh := len(tempSequence)
			if result < maxLengh {
				result = maxLengh
			}
			firstClient = requests[i-1]
			secondClient = c
			if len(consistent) == 0 {
				tempSequence = []int{firstClient}
			} else {
				tempSequence = consistent
			}
			tempSequence = append(tempSequence, secondClient)
			consistent = []int{}
			continue
		}
	}
	maxLengh := len(tempSequence)
	if result < maxLengh {
		result = maxLengh
	}
	return result
}

func scanNumbers() []int{
	input := bufio.NewReader(os.Stdin)
	var result []int
	var t int
	

	fmt.Fscan(input, &t)
	
    for i := 0; i < t; i++ {
		var lenRequests int
		var requests []int
		fmt.Fscan(input, &lenRequests)
		for y := 0; y < lenRequests; y++ {
			var client int
			fmt.Fscan(input, &client)
			requests = append(requests, client)
		}
		result = append(result, findMaxLenth(requests))
	}

	return result
}

func printResult(result []int) {
	for _, res := range result {
		fmt.Println(res)
	}
}

func main() {
	result := scanNumbers()
	printResult(result)

}