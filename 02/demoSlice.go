package main

import (
	"fmt"
)

func testSlice() {
	a := []string{"a", "b", "c", "d"}
	a = append(a, "e")
	fmt.Println(len(a))
	fmt.Println(a[:2])
	fmt.Println(a[2:5])
}

func removeItemSliceNotKeepOrder(a []string, i int) []string {
	a[i] = a[len(a)-1]  // Copy last element to index i.
	a[len(a)-1] = ""    // Erase last element (write zero value).
	return a[:len(a)-1] // Truncate slice.
}

func removeItemSliceNotKeepOrder2(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

/*
Xoá phần tử bảo toàn thứ tự
*/
func removeItemSliceKeepOrder(a []string, i int) []string {
	// Remove the element at index i from a.
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a
}

/*
Xoá phần tử bảo toàn thứ tự. Cách 2
*/
func removeItemSliceKeepOrder2(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}
