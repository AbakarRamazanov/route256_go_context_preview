package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &a, &b)
		fmt.Println(a + b)
	}
}
