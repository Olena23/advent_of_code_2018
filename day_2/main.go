package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"regexp"
	"index/suffixarray"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	sum2 := 0
	sum3 := 0
	file, err := os.Open("day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newStr := sc.Text()
		usedChars := make(map[string]bool)
		usedPairs := make(map[int]bool)
		for _, val := range newStr {
			if !usedChars[string(val)] {
				r := regexp.MustCompile(string(val))
				index := suffixarray.New([]byte(newStr))
				results := index.FindAllIndex(r, -1)
				usedChars[string(val)] = true
				if !usedPairs[2] {
					if len(results) == 2 {
						sum2++
						usedPairs[2] = true
					}
				}
				if !usedPairs[3] {
					if len(results) == 3 {
						sum3++
						usedPairs[3] = true
					}
				}
			}

		}
	}
	fmt.Println(sum2 * sum3)

	part2()

}

func part2()  {
	file, err := os.Open("day_2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var allStrings []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		newStr := sc.Text()
		allStrings = append(allStrings, newStr)
	}

	for ind, val := range allStrings {
		newStr := val
		for _, secondStr := range allStrings[ind+1:] {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(newStr, secondStr, false)
			sum := 0
			for i := range diffs {
				if diffs[i].Type == 0 {
					sum += len(diffs[i].Text)
				}
			}
			if sum == len(newStr)-1 {
				fmt.Println("Found!!!")
				fmt.Println(diffs)
				fmt.Println(dmp.DiffPrettyText(diffs))
			}
		}
	}
}