package main

import (
	"fmt"
)

func main() {
	n := 15
	for i := 0; i < n; i++ {
		fibo := getFibonacci(i)
		fmt.Println(i, ": ", fibo)
	}

}

func getFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return getFibonacci(n-1) + getFibonacci(n-2)
}
