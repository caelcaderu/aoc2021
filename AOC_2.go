package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	ftest, _ := os.Open("aoc_2_test_input")
// 	f, _ := os.Open("aoc_2_input")
// 	//aoc2part1(ftest)
// 	//aoc2part1(f)
// 	aoc2part2(ftest)
// 	aoc2part2(f)
// 	ftest.Close()
// 	f.Close()
// }

func aoc2part1(f *os.File) {
	horiz := 0
	depth := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		command := parts[0]
		n, _ := strconv.Atoi(parts[1])
		if command == "forward" {
			horiz += n
		} else if command == "down" {
			depth += n
		} else {
			depth -= n
		}
	}
	fmt.Println(horiz, depth, horiz*depth)
}

func aoc2part2(f *os.File) {
	horiz := 0
	depth := 0
	aim := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		command := parts[0]
		n, _ := strconv.Atoi(parts[1])
		if command == "forward" {
			horiz += n
			depth += (aim * n)
		} else if command == "down" {
			aim += n
		} else {
			aim -= n
		}
	}
	fmt.Println(horiz, depth, horiz*depth)
}
