package main
//
//import (
//	"bufio"
//	"log"
//	"os"
//	"sort"
//	"fmt"
//	"strings"
//	"strconv"
//)
//
//func main() {
//	file, err := os.Open("day_4/input.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//
//	var allNotes []string
//	lastGuard := 0
//	lastGuard = lastGuard
//	prevTime := 0
//	prevTime = prevTime
//	//curtTime := 0
//
//	//minutesAsleep := map[int]int{}
//	//for i := 0; i < 60; i++ {
//	//	minutesAsleep[i] = 0
//	//}
//
//
//	type minutesInfo struct {
//		minutesAsleep map[int]int
//	}
//
//	sleepingInfo := map[int]minutesInfo{}
//
//	sleepingInfo[0] = [8]
//
//	//for i := 0; i < 60; i++ {
//	//	sleepingInfo[90+i].minutesAsleep[i] = 0
//	//}
//
//
//	fmt.Println(sleepingInfo)
//
//	sc := bufio.NewScanner(file)
//	for sc.Scan() {
//		newStr := sc.Text()
//		allNotes = append(allNotes, newStr)
//	}
//
//	sort.Sort(sort.StringSlice(allNotes))
//
//	for i := range allNotes {
//		time := strings.Split(allNotes[i], "] ")[0]
//		minute, _ := strconv.Atoi(time[len(time)-2:])
//
//		if strings.Contains(allNotes[i], "Guard") {
//			num := strings.Split(allNotes[i], "#")[1]
//			id, _ := strconv.Atoi(strings.Split(num, " ")[0])
//			lastGuard = id
//			prevTime = minute
//		}
//
//		if strings.Contains(allNotes[i], "falls ") {
//
//		}
//
//		if strings.Contains(allNotes[i], "wakes ") {
//
//		}
//	}
//	//fmt.Println(lastGuard)
//}
//
