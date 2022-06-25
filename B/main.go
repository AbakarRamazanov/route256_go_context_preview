package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var countRows int
	// fmt.Scan(&countRows)
	// var n int
	// var key int
	// for j := 0; j < countRows; j++ {
	// 	m := make(map[int]int)
	// 	fmt.Scan(&n)
	// 	fmt.Fscan(os.Stdin, &n)
	// 	for i := 0; i < n; i++ {
	// 		fmt.Scan(&key)
	// 		m[key]++
	// 	}
	// 	sum := 0
	// 	for key, value := range m {
	// 		h := value / 3
	// 		sum += (value - h) * key
	// 	}
	// 	fmt.Println(sum)
	// }
	printout(calculate(splitString(readin())))
}

func printout(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func calculate(in <-chan map[int]int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := 0
		for m := range in {
			sum = 0
			for key, value := range m {
				sum += (value - (value / 3)) * key
			}
			out <- sum
		}
	}()
	return out
}

func splitString(in <-chan string) <-chan map[int]int {
	out := make(chan map[int]int)
	go func() {
		defer close(out)
		var m map[int]int
		var array []string
		var i int
		for s := range in {
			m = make(map[int]int)
			array = strings.Split(s, " ")
			for _, s := range array {
				i, _ = strconv.Atoi(s)
				m[i]++
			}
			out <- m
		}
	}()
	return out
}

func readin() <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		var countRows int
		var s string
		fmt.Scan(&countRows)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
		for i := 0; i < countRows; i++ {
			reader.ReadString('\n')
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
			out <- s
		}
	}()
	return out
}
