package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseCalories(filename string) [][]int {
	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)

	var prevLineReturn bool = true

	var calsArrArr [][]int
	for scanner.Scan() {
		currLine := scanner.Text()

		// append new array at breakpoint
		if prevLineReturn && currLine != "" {
			calsArrArr = append(calsArrArr, []int{})
		}

		// set prevLineReturn for next iter
		if currLine != "" {
			prevLineReturn = false
		} else {
			prevLineReturn = true
			continue
		}

		// add calorie to current arr
		num, _ := strconv.Atoi(currLine)
		calsArrArr[len(calsArrArr)-1] = append(calsArrArr[len(calsArrArr)-1], num)
	}

	return calsArrArr
}

func processCalories(calsArrArr [][]int) []int {
	var calsArr []int
	for _, v := range calsArrArr {
		calsArr = append(calsArr, sumIntArr(v))
	}
	return calsArr
}

func sumIntArr(arr []int) int {
	var num int = 0
	for _, v := range arr {
		num += v
	}
	return num
}

func intArrMax(arr []int) int {
	var outNum int
	for i, v := range arr {
		if i == 0 || v > outNum {
			outNum = v
		}
	}
	return outNum
}

func main() {

	// slice of arguments without the executable itself
	args := os.Args[1:]

	fmt.Println(intArrMax(processCalories(parseCalories(args[0]))))

}
