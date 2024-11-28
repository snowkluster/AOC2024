package main

import (
  "fmt"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	fmt.Println("Welcome to AOC 2024")
	data, err := os.ReadFile("input.txt")
  check(err)
  fmt.Print(string(data))
}
