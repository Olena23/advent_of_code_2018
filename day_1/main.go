package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"fmt"
)

func main() {
	sum := 0
	var allFrequencies []int
	allFrequencies = append(allFrequencies, 0)

	scan(sum, allFrequencies)
}


func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func scan (sum int, allFrequencies []int) int {
	file, err := os.Open("day_1/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newInt, _ := strconv.Atoi(sc.Text())
		sum += newInt
		if contains(allFrequencies, sum) {
			fmt.Println("Found!")
			fmt.Println(sum)
			return sum
		}
		allFrequencies = append(allFrequencies, sum)
	}
	scan(sum, allFrequencies)
	return 0
}
