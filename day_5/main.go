package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	file, err := os.Open("day_5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	longStr := ""
	defer file.Close()
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		longStr = sc.Text()
	}

	size := len(longStr)
	out := deletePair(longStr, size)
	fmt.Println(out)


	//Part II
	len, char := deleteOne(longStr)
	fmt.Println(len)
	fmt.Println(char)
}

func deletePair(longStr string, size int) int {
	size2 := 0
	for {
		if size > size2 {
			size = len(longStr)
			for i := range longStr {
				if i+1 < len(longStr) {
					if longStr[i] != longStr[i+1] {
						if string(longStr[i]) == strings.ToLower(string(longStr[i+1])) ||
							string(longStr[i]) == strings.ToUpper(string(longStr[i+1])) {
							if i+2 <= len(longStr) {
								longStr = longStr[:i] + longStr[i+2:]
							}
							break
						}
					}
				}
			}
			size2 = len(longStr)
		}
		if size2 == size {
			break
		}
	}

	return len(longStr)
}

func deleteOne(longStr string) (int, string) {

	usedPairs := map[string]bool{}

	shortestStr := len(longStr)
	shortestChar := ""
	newLen := len(longStr)

	for i := range longStr {
		if usedPairs[strings.ToLower(string(longStr[i]))] == false {
			usedPairs[strings.ToLower(string(longStr[i]))] = true
			usedPairs[strings.ToUpper(string(longStr[i]))] = true

			testStr := longStr
			testStr = strings.Replace(testStr, strings.ToLower(string(longStr[i])), "", -1)
			testStr = strings.Replace(testStr, strings.ToUpper(string(longStr[i])), "", -1)
			newLen = deletePair(testStr, len(testStr))
		}
		if newLen < shortestStr {
			shortestStr = newLen
			shortestChar = string(longStr[i])
			fmt.Println(shortestChar)
		}

	}

	return shortestStr, shortestChar
}
