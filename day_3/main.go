package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum = 0
	var fabric [4000][4000]int
	allStrings := make([]string, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newStr := sc.Text()
		allStrings = append(allStrings, newStr)
	}

	for _, str := range allStrings {
		words := strings.Split(str, " ")

		padding := strings.Split(words[2], ",")
		fromLeft, _ := strconv.Atoi(padding[0])
		fromTop, _ := strconv.Atoi(strings.TrimSuffix(padding[1], ":"))

		dimensions := strings.Split(words[3], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if fabric[fromTop+i][fromLeft+j] == 1 {
					sum++
					fabric[fromTop+i][fromLeft+j] = 2
				} else if fabric[fromTop+i][fromLeft+j] == 2 {
					continue
				} else {
					fabric[fromTop+i][fromLeft+j] = 1
				}
			}
		}

	}

	fmt.Println(sum)

	part2()
}

func part2() {

	file, err := os.Open("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var sum = 0
	var fabric [4000][4000]int
	allStrings := make([]string, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newStr := sc.Text()
		allStrings = append(allStrings, newStr)
	}

	for _, str := range allStrings {
		words := strings.Split(str, " ")

		padding := strings.Split(words[2], ",")
		fromLeft, _ := strconv.Atoi(padding[0])
		fromTop, _ := strconv.Atoi(strings.TrimSuffix(padding[1], ":"))

		dimensions := strings.Split(words[3], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])

		sumOfsquares := 0
		for m := 0; m < height; m++ {
			for k := 0; k < width; k++ {
				if fabric[fromTop+m][fromLeft+k] == 1 {
					sum++
					fabric[fromTop+m][fromLeft+k] = 2
				} else if fabric[fromTop+m][fromLeft+k] == 2 {
					continue
				} else {
					fabric[fromTop+m][fromLeft+k] = 1
					sumOfsquares++
				}
			}
		}
	}

	for _, str := range allStrings {
		words := strings.Split(str, " ")

		padding := strings.Split(words[2], ",")
		fromLeft, _ := strconv.Atoi(padding[0])
		fromTop, _ := strconv.Atoi(strings.TrimSuffix(padding[1], ":"))

		dimensions := strings.Split(words[3], "x")
		width, _ := strconv.Atoi(dimensions[0])
		height, _ := strconv.Atoi(dimensions[1])

		sumOfsquares := 0
		for m := 0; m < height; m++ {
			for k := 0; k < width; k++ {
				if fabric[fromTop+m][fromLeft+k] == 1 {
					sumOfsquares++
				}
			}
		}

		if sumOfsquares == width*height {
			fmt.Println("Found!!!")
			fmt.Println(words)
		}

	}

}
