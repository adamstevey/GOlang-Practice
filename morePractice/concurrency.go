package main

import (
	"fmt"
	"time"
)

func main() {
	go concurrency()
	printLine()
}

func concurrency() {
	for {
		time.Sleep(1000)
		fmt.Println("running")
	}
}

func printLine() {
	for {
		time.Sleep(100)
		fmt.Println("Printing")
	}
}
