package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	out := sortingDataset(generateDataset(readin()))
	s := ""
	for o := range out {
		for _, row := range o.matrix {
			s = fmt.Sprint(row)
			s = s[1 : len(s)-1]
			fmt.Printf("%s\n", s)
		}
		fmt.Println()
	}
	fmt.Println()
}

type Dataset struct {
	matrix [][]int
	clicks []int
}

func sortingDataset(in <-chan Dataset) <-chan Dataset {
	out := make(chan Dataset)
	go func() {
		defer close(out)
		for dataset := range in {
			for _, key := range dataset.clicks {
				sort.SliceStable(dataset.matrix, func(i, j int) bool {
					return dataset.matrix[j][key-1] > dataset.matrix[i][key-1]
				})
			}
			out <- dataset
		}
	}()
	return out
}

func generateDataset(in <-chan []string) <-chan Dataset {
	out := make(chan Dataset)
	var ds Dataset
	go func() {
		defer close(out)
		var splited []string
		var splitedInt int
		for s := range in {
			ds = Dataset{}
			for i := 0; i < len(s)-2; i++ {
				ds.matrix = append(ds.matrix, []int{})
				splited = strings.Split(s[i], " ")
				for _, s := range splited {
					splitedInt, _ = strconv.Atoi(s)
					ds.matrix[i] = append(ds.matrix[i], splitedInt)
				}
			}
			splited = strings.Split(s[len(s)-1], " ")
			ds.clicks = make([]int, 0, len(ds.matrix[0]))
			// splitedInt, _ = strconv.Atoi(splited[0])
			// ds.clicks = append(ds.clicks, splitedInt)
			// splited = splited[1:]
			for _, s := range splited {
				splitedInt, _ = strconv.Atoi(s)
				// if ds.clicks[len(ds.clicks)-1] != splitedInt {
				ds.clicks = append(ds.clicks, splitedInt)
				// }
			}
			out <- ds
		}
	}()
	return out
}

func readin() <-chan []string {
	out := make(chan []string)
	go func() {
		defer close(out)
		var countDataset int
		var row int
		var s []string
		fmt.Scan(&countDataset)
		reader := bufio.NewReader(os.Stdin)
		l := ""
		reader.ReadString('\n')
		for i := 0; i < countDataset; i++ {
			reader.ReadString('\n')
			l, _ = reader.ReadString('\n')
			l = strings.TrimSuffix(l, "\r\n")
			row, _ = strconv.Atoi(strings.Split(l, " ")[0])
			s = make([]string, row+2)
			for j := 0; j < row; j++ {
				s[j], _ = reader.ReadString('\n')
				s[j] = strings.TrimSuffix(s[j], "\r\n")
			}
			s[row], _ = reader.ReadString('\n')
			s[row] = strings.TrimSuffix(s[row], "\r\n")
			s[row+1], _ = reader.ReadString('\n')
			s[row+1] = strings.TrimSuffix(s[row+1], "\r\n")
			out <- s
		}
	}()
	return out
}
