package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sum := 0
	allFrequencies := make(map[int]struct{}, 0)

	scan(sum, allFrequencies)
}

func scan(sum int, allFrequencies map[int]struct{}) int {
	file, err := os.Open("day_1/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newInt, _ := strconv.Atoi(sc.Text())
		sum += newInt
		if _, ok := allFrequencies[sum]; ok {
			fmt.Println("Found!")
			fmt.Println(sum)
			return sum
		}

		allFrequencies[sum] = struct{}{}
	}
	scan(sum, allFrequencies)
	return 0
}
