package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func diffcheck(intarr []int) (check bool) {
	for i := 0; i < len(intarr) - 1; i++ {
		diff := math.Abs(float64(intarr[i] - intarr[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func decend(intarr []int) bool {
	for i := 0; i < len(intarr)-1; i++ {
		if intarr[i] <= intarr[i+1] {
			return false
		}
	}
	return true
}

func ascend(intarr []int) bool {
	for i := 0; i < len(intarr)-1; i++ {
		if intarr[i] >= intarr[i+1] {
			return false
		}
	}
	return true
}

func day2s1(data []byte) (count int) {
	var intarr []int
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		intarr = nil
		for _, str := range arr {
			val, _ := strconv.Atoi(str)
			intarr = append(intarr, val)
		}
		if (ascend(intarr) || decend(intarr)) && diffcheck(intarr) {
			count++
		}
	}
	return
}

func main() {
	data, _ := os.ReadFile("input.txt")
	count := day2s1(data)
	fmt.Println(count)
}