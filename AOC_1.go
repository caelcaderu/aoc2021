package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	//	part1()
// 	part2()
// }

func part1() {
	//dat, err := os.ReadFile("test_text.txt")
	//f, _ := os.Open("test_text.txt")
	f, _ := os.Open("aoc_1_input")
	sc := bufio.NewScanner(f)
	counter := 0
	sc.Scan()
	first, _ := strconv.Atoi(sc.Text())
	for sc.Scan() {
		second, _ := strconv.Atoi(sc.Text())
		if second > first {
			counter++
		}
		first = second
	}
	fmt.Print(counter)

	f.Close()
}

func part2() {
	//f, _ := os.Open("test_text.txt")
	f, _ := os.Open("aoc_1_input")
	sc := bufio.NewScanner(f)
	counter := 0
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ := strconv.Atoi(sc.Text())
	for sc.Scan() {
		d, _ := strconv.Atoi(sc.Text())
		if d > a {
			counter++
		}
		a = b
		b = c
		c = d
	}
	fmt.Print(counter)
	f.Close()

}
