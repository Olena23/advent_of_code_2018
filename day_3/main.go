package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var sum = 0
	var fabric [4000][4000]int
	var allStrings []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newStr := sc.Text()
		allStrings = append(allStrings, newStr)
	}

	for i := range allStrings {
		one := strings.Split(allStrings[i], " ")
		fromLeft := strings.Split(one[2], ",")[0]
		fromLeftInt, _ := strconv.Atoi(fromLeft)
		fromTop := strings.Split(one[2], ",")[1]
		fromTop = fromTop[0:len(fromTop)-1]
		fromTopInt, _ := strconv.Atoi(fromTop)
		width := strings.Split(one[3], "x")[0]
		widthInt, _ := strconv.Atoi(width)
		height := strings.Split(one[3], "x")[1]
		heightInt, _ := strconv.Atoi(height)

		for m := 0; m < heightInt; m++ {
			for k := 0; k < widthInt; k++ {
				if fabric[fromTopInt + m][fromLeftInt + k] == 1 {
					sum++
					fabric[fromTopInt+ m][fromLeftInt + k] = 2
				} else if fabric[fromTopInt+ m][fromLeftInt + k] == 2 {
					continue
				} else {
					fabric[fromTopInt+ m][fromLeftInt + k] = 1
				}
			}
		}


	}

	fmt.Println(sum)

	part2()
}

func part2 ()  {

	file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var sum = 0
	var fabric [4000][4000]int
	var allStrings []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newStr := sc.Text()
		allStrings = append(allStrings, newStr)
	}

	for i := range allStrings {
		one := strings.Split(allStrings[i], " ")
		fromLeft := strings.Split(one[2], ",")[0]
		fromLeftInt, _ := strconv.Atoi(fromLeft)
		fromTop := strings.Split(one[2], ",")[1]
		fromTop = fromTop[0:len(fromTop)-1]
		fromTopInt, _ := strconv.Atoi(fromTop)
		width := strings.Split(one[3], "x")[0]
		widthInt, _ := strconv.Atoi(width)
		height := strings.Split(one[3], "x")[1]
		heightInt, _ := strconv.Atoi(height)

		sumOfsquares := 0
		for m := 0; m < heightInt; m++ {
			for k := 0; k < widthInt; k++ {
				if fabric[fromTopInt + m][fromLeftInt + k] == 1 {
					sum++
					fabric[fromTopInt+ m][fromLeftInt + k] = 2
				} else if fabric[fromTopInt+ m][fromLeftInt + k] == 2 {
					continue
				} else {
					fabric[fromTopInt+ m][fromLeftInt + k] = 1
					sumOfsquares ++
				}
			}
		}
		}



	for i := range allStrings {
		one := strings.Split(allStrings[i], " ")
		fromLeft := strings.Split(one[2], ",")[0]
		fromLeftInt, _ := strconv.Atoi(fromLeft)
		fromTop := strings.Split(one[2], ",")[1]
		fromTop = fromTop[0:len(fromTop)-1]
		fromTopInt, _ := strconv.Atoi(fromTop)
		width := strings.Split(one[3], "x")[0]
		widthInt, _ := strconv.Atoi(width)
		height := strings.Split(one[3], "x")[1]
		heightInt, _ := strconv.Atoi(height)

		sumOfsquares := 0
		for m := 0; m < heightInt; m++ {
			for k := 0; k < widthInt; k++ {
				if fabric[fromTopInt + m][fromLeftInt + k] == 1 {
					sumOfsquares++
				}
			}
		}

		if sumOfsquares == widthInt * heightInt {
			fmt.Println("Found!!!")
			fmt.Println(one)
		}



	}


	}