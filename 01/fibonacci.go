package main

import (
	"fmt"
)

func main() {
	n := 15
	for i := 1; i <= n; i++ {
		fibo := getFibonacci(i)
		fmt.Println(i, ": ", fibo)
	}

}

func getFibonacci(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	a := 0
	b := 1

	for temp := 3; temp <= n; temp++ {
		b = a + b
		a = b - a
	}
	return b
}
