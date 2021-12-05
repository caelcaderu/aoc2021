package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	ftest, _ := os.Open("aoc_4_test_input")
// 	f, _ := os.Open("aoc_4_input")
// 	//aoc4part1(ftest)
// 	//aoc4part1(f)
// 	aoc4part2(ftest)
// 	aoc4part2(f)
// 	ftest.Close()
// 	f.Close()
// }

func aoc4part1(f *os.File) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	seq := sc.Text()
	fmt.Println("Sequence:", seq)
	boards := make([][][]int, 0)

	//Get the boards
	for sc.Scan() { //there is a blank line (not end of input)
		boards = append(boards, getBoard(sc))
	}
	for _, b := range boards {
		printBoard(b)
	}

	litup := make([][5][5]bool, len(boards))

	//Now we're ready to solve!
	seqi := Apply(strconv.Atoi, strings.Split(seq, ","))
	for _, input := range seqi {
		fmt.Println("Drawn:", input)
		for b, boa := range boards {
			for r, row := range boa {
				for c, col := range row {
					if input == col {
						litup[b][r][c] = true
						//Check for victory
						victory := true
						for horiz := 0; horiz < 5; horiz++ {
							if !litup[b][r][horiz] {
								victory = false
							}
						}
						if victory {
							fmt.Println("Horizontal victory on board", b)
							printBoard(boards[b])
							printLit(litup[b])
							fmt.Println("Winning score:", getWinningScore(boards[b], litup[b], input))
							return
						}
						victory = true
						for vert := 0; vert < 5; vert++ {
							if !litup[b][vert][c] {
								victory = false
							}
						}
						if victory {
							fmt.Println("Vertical victory on board", b)
							printBoard(boards[b])
							printLit(litup[b])
							fmt.Println("Winning score:", getWinningScore(boards[b], litup[b], input))
							return
						}
					}
				}
			}
		}
	}
}

func aoc4part2(f *os.File) {
	sc := bufio.NewScanner(f)
	sc.Scan()
	seq := sc.Text()
	fmt.Println("Sequence:", seq)
	boards := make([][][]int, 0)

	//Get the boards
	for sc.Scan() { //there is a blank line (not end of input)
		boards = append(boards, getBoard(sc))
	}
	for _, b := range boards {
		printBoard(b)
	}

	litup := make([][5][5]bool, len(boards))

	//Up to this point exactly like part 1

	//Now we want to solve, but we want to DELETE boards that win
	// until only one is left
	// then find its score when it wins
	seqi := Apply(strconv.Atoi, strings.Split(seq, ","))
	for _, input := range seqi {
		fmt.Println("There are", len(boards), "board remaining...")
		fmt.Println("  Drawn:", input)
		b := 0
		for b < len(boards) {
			//boardActive := true
			for r, row := range boards[b] {
				for c, col := range row {
					if input == col {
						litup[b][r][c] = true
						//Check for victory
						victory := true
						for horiz := 0; horiz < 5; horiz++ {
							if !litup[b][r][horiz] {
								victory = false
							}
						}
						if victory {
							if len(boards) > 1 {
								//Need to delete the board from both boards and litup
								boards = append(boards[:b], boards[b+1:]...)
								litup = append(litup[:b], litup[b+1:]...)
								fmt.Println("Deleting board", b)
								b--
								//boardActive = false
								//break
							} else {
								fmt.Println("Horizontal victory on board", b)
								printBoard(boards[b])
								printLit(litup[b])
								fmt.Println("Winning score:", getWinningScore(boards[b], litup[b], input))
								return
							}
						} else {
							victory = true
							for vert := 0; vert < 5; vert++ {
								if !litup[b][vert][c] {
									victory = false
								}
							}
							if victory {
								if len(boards) > 1 {
									boards = append(boards[:b], boards[b+1:]...)
									litup = append(litup[:b], litup[b+1:]...)
									fmt.Println("Deleting board", b)
									b--
									//boardActive = false
									//break
								} else {
									fmt.Println("Vertical victory on board", b)
									printBoard(boards[b])
									printLit(litup[b])
									fmt.Println("Winning score:", getWinningScore(boards[b], litup[b], input))
									return
								}
							}
						}
					}
					// if !boardActive {
					// 	break
					// }
				}
				// if !boardActive {
				// 	break
				// }
			}
			b++
		}
	}
}

func getWinningScore(b [][]int, blit [5][5]bool, winnum int) int {
	//Sum of all unmarked numbers on the board
	sum := 0
	for r := range b {
		for c := range b[0] {
			if !blit[r][c] {
				sum += b[r][c]
			}
		}
	}

	//Multiply the sum by the winning number
	return sum * winnum
}

func printBoard(b [][]int) {
	for _, row := range b {
		for _, val := range row {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printLit(b [5][5]bool) {
	for _, row := range b {
		for _, val := range row {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}

// func Apply(f func(string) (int, error), vs []string) []int {
// 	out := make([]int, len(vs))
// 	for i, v := range vs {
// 		out[i], _ = f(v)
// 	}
// 	return out
// }

func getBoard(sc *bufio.Scanner) [][]int {
	b := make([][]int, 0)
	for i := 0; i < 5; i++ {
		sc.Scan()
		b = append(b, Apply(strconv.Atoi, strings.Fields(sc.Text())))
	}
	//fmt.Println("Made board:")
	//printBoard(b)
	return b
}
