package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func day3s1(data []byte) (ret int) {
	pat := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pat)
	matches := re.FindAllStringSubmatch(string(data), -1)
	fmt.Println(matches)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		ret += num1 * num2
	}
	return
}

func main() {
	data, _ := os.ReadFile("input.txt")
	sum := day3s1(data)
	fmt.Println(sum)
}