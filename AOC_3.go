package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	ftest, _ := os.Open("aoc_3_test_input")
// 	f, _ := os.Open("aoc_3_input")
// 	//aoc3part1(ftest)
// 	//aoc3part1(f)
// 	//aoc3part2(ftest)
// 	aoc3part2(f)
// 	ftest.Close()
// 	f.Close()
// }

func aoc3part1(f *os.File) {
	linelen := 12 //5 for the test file
	gamma := 0
	epsilon := 0
	linecount := 0
	counts := make([]int, linelen) //initially all zeros
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		linecount++
		tex := sc.Text()
		//fmt.Println(tex)
		for i := 0; i < linelen; i++ {
			counts[i] += int(tex[i] - '0')
		}
	}
	for i := 0; i < linelen; i++ {
		if counts[i] > linecount/2 {
			//More 1s than 0s
			gamma += 1
		} else {
			epsilon += 1
		}
		if i != linelen-1 {
			gamma *= 2
			epsilon *= 2
		}
	}
	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Println(gamma * epsilon)
}

func aoc3part2(f *os.File) { //BAD ANSWER: 3143647
	linelen := 12 //5 for the test file, 12 for real
	linecount := 0
	//	counts := make([]int, linelen) //initially all zeros
	sc := bufio.NewScanner(f)
	alllines := make([]string, 0)
	for sc.Scan() {
		linecount++
		tex := sc.Text()
		alllines = append(alllines, tex)
		//fmt.Println(tex)
		// for i := 0; i < linelen; i++ {
		// 	counts[i] += int(tex[i] - '0')
		// }
	}

	// oxy := make([]int, linelen)
	// co2 := make([]int, linelen)
	// for i := 0; i < linelen; i++ {
	// 	if counts[i] >= linecount/2 {
	// 		//More 1s than 0s
	// 		oxy[i] = '1'
	// 		co2[i] = '0'
	// 	} else {
	// 		co2[i] = '1'
	// 		oxy[i] = '0'
	// 	}
	// }
	// fmt.Println("Oxy mask:", oxy)
	// fmt.Println("Co2 mask:", co2)

	// //Find out who matches the mask
	// oxymatches := make([]int, linecount)
	// comatches := make([]int, linecount)
	// for j := 0; j < linecount; j++ {
	// 	for i := linelen - 1; i >= 0; i-- {
	// 		if int(alllines[j][i]) != oxy[i] {
	// 			oxymatches[j] = i
	// 		}
	// 		if int(alllines[j][i]) != co2[i] {
	// 			comatches[j] = i
	// 		}
	// 	}
	// }

	// We now have all the mask match distances
	// Who matches farthest?
	// fmt.Println("Oxy matching...")
	// oxymax := oxymatches[0]
	// oxymaxi := 0
	// for i, v := range oxymatches {
	// 	if v >= oxymax {
	// 		oxymaxi = i
	// 		oxymax = v
	// 		fmt.Println(i, v, alllines[i])
	// 	}
	// }

	// fmt.Println("Co2 matching...")
	// comax := comatches[0]
	// comaxi := 0
	// for i, v := range comatches {
	// 	if v >= comax {
	// 		comaxi = i
	// 		comax = v
	// 		fmt.Println(i, v, alllines[i])
	// 	}
	// }

	//BIT BY BIT
	oxyremaining := alllines
	oxymask := '0'
	oxybin := ""
	for i := 0; i < linelen; i++ {
		oxycount := 0
		for _, x := range oxyremaining {
			if x[i] == '1' {
				oxycount++
			}
		}
		if float32(oxycount) >= float32(len(oxyremaining))/2.0 {
			oxymask = '1'
		} else {
			oxymask = '0'
		}
		future := make([]string, 0)
		for _, x := range oxyremaining {
			if int(x[i]) == int(oxymask) {
				future = append(future, x)
			}
		}
		fmt.Println(len(future))
		if len(future) == 1 {
			fmt.Println("Best oxy:", future[0])
			oxybin = future[0]
			break
		}
		oxyremaining = future
	}

	coremaining := alllines
	comask := '0'
	cobin := ""
	for i := 0; i < linelen; i++ {
		cocount := 0
		for _, x := range coremaining {
			if x[i] == '1' {
				cocount++
			}
		}
		fmt.Println("Diagnostic:", i, cocount, len(coremaining)/2.0)
		if float32(cocount) < float32(len(coremaining))/2.0 {
			comask = '1'
		} else {
			comask = '0'
		}
		future := make([]string, 0)
		for _, x := range coremaining {
			if int(x[i]) == int(comask) {
				future = append(future, x)
				fmt.Println(x)
			}
		}
		fmt.Println(len(future))
		if len(future) == 1 {
			fmt.Println("Best co:", future[0])
			cobin = future[0]
			break
		}
		if len(future) == 0 {
			fmt.Println("Best co (tie!):", coremaining[len(coremaining)-1])
			cobin = coremaining[len(coremaining)-1]
			break
		}
		coremaining = future
	}

	//Convert to decimal
	oxydec := 0
	codec := 0
	for i := 0; i < linelen; i++ {
		oxydec += int(oxybin[i] - '0')
		codec += int(cobin[i] - '0')
		if i != linelen-1 {
			oxydec *= 2
			codec *= 2
		}
	}
	fmt.Println(oxydec)
	fmt.Println(codec)
	fmt.Println(oxydec * codec)
}
