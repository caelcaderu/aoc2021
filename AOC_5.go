package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//aoc5part1("aoc_5_test_input")
	//aoc5part1("aoc_5_input")
	aoc5part2("aoc_5_test_input")
	aoc5part2("aoc_5_input")
}

func aoc5part1(fname string) {
	//Horiz or Vert lines only
	//On how many points do at least two lines overlap?
	data := getLines(fname)
	for _, d := range data {
		fmt.Println(d.x1, d.y1, "->", d.x2, d.y2)
	}

	//Build a grid of zeros
	xmax, ymax := getMaxCoords(data)
	grid := make([][]int, ymax+1)
	for i := 0; i < ymax+1; i++ {
		grid[i] = make([]int, xmax+1)
	}
	printGrid(grid)

	//Update the grid using the lines
	grid = fillGrid(grid, data, xmax, ymax)
	printGrid(grid)

	//Count 2s+
	count := count2orMore(grid)
	fmt.Println("The count is:", count)
}

func aoc5part2(fname string) {
	//Horiz or Vert lines only
	//On how many points do at least two lines overlap?
	data := getLines(fname)
	for _, d := range data {
		fmt.Println(d.x1, d.y1, "->", d.x2, d.y2)
	}

	//Build a grid of zeros
	xmax, ymax := getMaxCoords(data)
	grid := make([][]int, ymax+1)
	for i := 0; i < ymax+1; i++ {
		grid[i] = make([]int, xmax+1)
	}
	//printGrid(grid)

	//Update the grid using the lines
	grid = fillGrid(grid, data, xmax, ymax)
	//printGrid(grid)
	grid = fillDiagonals(grid, data)
	printGrid(grid)

	//Count 2s+
	count := count2orMore(grid)
	fmt.Println("The count is:", count)
}

func count2orMore(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val > 1 {
				count++
			}
		}
	}
	return count
}

func fillDiagonals(grid [][]int, lines []gridLine) [][]int {
	fmt.Println("Filling the diagonals, line by line...")
	for _, line := range lines {
		if line.y1 == line.y2 || line.x1 == line.x2 {
			continue
		}
		xstart := Min(line.x1, line.x2)
		xend := Max(line.x1, line.x2)
		ystart := Min(line.y1, line.y2)
		yend := Max(line.y1, line.y2)
		if (xstart == line.x1 && ystart == line.y1) ||
			(xstart == line.x2 && ystart == line.y2) {
			for i := 0; i < xend-xstart+1; i++ {
				grid[ystart+i][xstart+i] += 1
			}
		} else {
			for i := 0; i < xend-xstart+1; i++ {
				grid[yend-i][xstart+i] += 1
			}
		}
	}
	return grid
}

func fillGrid(grid [][]int, lines []gridLine, xmax int, ymax int) [][]int {
	fmt.Println("Filling the grid, line by line...")
	//printGrid(grid)
	for _, line := range lines {
		isHorizontal := line.y1 == line.y2
		isVertical := line.x1 == line.x2
		if isHorizontal {
			y := line.y1
			xstart := Min(line.x1, line.x2)
			xend := Max(line.x1, line.x2)
			for i := xstart; i <= xend; i++ {
				grid[y][i] += 1
			}
		} else if isVertical {
			x := line.x1
			ystart := Min(line.y1, line.y2)
			yend := Max(line.y1, line.y2)
			for i := ystart; i <= yend; i++ {
				grid[i][x] += 1
			}
		}
		//printGrid(grid)
	}
	return grid
}

func printGrid(grid [][]int) {
	fmt.Println("The Grid:")
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func getLines(fname string) []gridLine {
	f, _ := os.Open(fname)
	sc := bufio.NewScanner(f)
	allLines := make([]gridLine, 0)
	for sc.Scan() {
		line := sc.Text()
		coords := strings.Split(line, " -> ")
		coords = append(strings.Split(coords[0], ","), strings.Split(coords[1], ",")...)
		coordInts := Apply(strconv.Atoi, coords)
		allLines = append(allLines, gridLine{coordInts[0], coordInts[1], coordInts[2], coordInts[3]})
	}

	f.Close()
	return allLines
}

//Return the maximum x and y values
func getMaxCoords(lines []gridLine) (int, int) {
	xmax := 0
	ymax := 0
	for _, line := range lines {
		if line.x1 > xmax {
			xmax = line.x1
		}
		if line.x2 > xmax {
			xmax = line.x2
		}
		if line.y1 > ymax {
			ymax = line.y1
		}
		if line.y2 > ymax {
			ymax = line.y2
		}
	}
	return xmax, ymax
}

type gridLine struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func Apply(f func(string) (int, error), vs []string) []int {
	out := make([]int, len(vs))
	for i, v := range vs {
		out[i], _ = f(v)
	}
	return out
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
