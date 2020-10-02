package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	out := deletePair(longStr)
	fmt.Println(out)

	//Part II
	len, char := deleteOne(longStr)
	fmt.Println(len)
	fmt.Println(char)
}

func deletePair(longStr string) int {
	size := len(longStr)
	size2 := 0
	for {
		if size > size2 {
			size = len(longStr)
			for i := range longStr {
				if i+1 < len(longStr) {
					if longStr[i] != longStr[i+1] {
						if strings.ToLower(string(longStr[i])) == strings.ToLower(string(longStr[i+1])) {
							if i+2 <= len(longStr) {
								longStr = longStr[:i] + longStr[i+2:]
								i += 2
							}
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

	usedPairs := map[string]struct{}{}

	shortestStr := len(longStr)
	shortestChar := ""
	newLen := len(longStr)

	for i := range longStr {
		if _, ok := usedPairs[strings.ToLower(string(longStr[i]))]; !ok {
			usedPairs[strings.ToLower(string(longStr[i]))] = struct{}{}
			usedPairs[strings.ToUpper(string(longStr[i]))] = struct{}{}

			testStr := longStr
			testStr = strings.Replace(testStr, strings.ToLower(string(longStr[i])), "", -1)
			testStr = strings.Replace(testStr, strings.ToUpper(string(longStr[i])), "", -1)
			newLen = deletePair(testStr)
		}
		if newLen < shortestStr {
			shortestStr = newLen
			shortestChar = string(longStr[i])
			fmt.Println(shortestChar)
		}

	}

	return shortestStr, shortestChar
}
