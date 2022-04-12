package main

import (
	"fmt"
	"sort"
)

func sortSliceWithFunc() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 52},
		{"Jun", 25},
		{"Bob", 43},
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	fmt.Println("By name: ", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("By age: ", people)

}
