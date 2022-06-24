package main

import "fmt"

func main() {
	var countRows int
	fmt.Scanf("%d\n", &countRows)

	var n int
	var key int
	for j := 0; j < countRows; j++ {
		m := make(map[int]int)
		fmt.Scanf("%d\n", &n)
		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &key)
			m[key]++
		}
		// for key, value := range m{
		// TODO calculation final sum
		// }
		fmt.Println(m)
	}
}
