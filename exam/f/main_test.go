package main

import (
	"log"
	"os"
	"testing"
)

func TestMain(t *testing.T){
	content, err := os.Open("./tests/25")
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()
	os.Stdin = content
	main()
}