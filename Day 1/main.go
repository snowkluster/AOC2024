package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
  "math"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func day1s1(data []byte) (sum int){
  var arr1 []int
  var arr2 []int
  scanner := bufio.NewScanner(bytes.NewReader(data))
  for scanner.Scan() {
    line := strings.Split(scanner.Text(), "   ")
    int1, err := strconv.ParseInt(line[0] ,10, 32)
    check(err)
    int2, err := strconv.ParseInt(line[1] ,10, 32)
    check(err)
    arr1 = append(arr1, int(int1))
    arr2 = append(arr2, int(int2))
  }
  sort.Ints(arr1)
  sort.Ints(arr2)
  for i:= range len(arr1) {
      distance := math.Abs(float64(arr1[i] - arr2[i]))
      sum += int(distance)
  }
  return
}

func day1s2(data []byte) (find int){
  var arr1 []int
  var arr2 []int
  scanner := bufio.NewScanner(bytes.NewReader(data))
  for scanner.Scan() {
    line := strings.Split(scanner.Text(), "   ")
    int1, err := strconv.ParseInt(line[0] ,10, 32)
    check(err)
    int2, err := strconv.ParseInt(line[1] ,10, 32)
    check(err)
    arr1 = append(arr1, int(int1))
    arr2 = append(arr2, int(int2))
  }
  sort.Ints(arr1)
  sort.Ints(arr2)
  freq := make(map[int]int)
  for _, num := range arr2 {
    freq[num]++
  }
  for _, num := range arr1 {
		find += freq[num] * num
	}
  return
}

func main() {
	data, err := os.ReadFile("input.txt")
  const prec = uint(100)
  check(err)
  sum := day1s1(data)
  find := day1s2(data)
  fmt.Println(sum)
  fmt.Println(find)
}
