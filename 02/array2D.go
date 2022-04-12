package main

import "fmt"

func testArray2D() {
	langs := [][]string{{"C#", "C", "Python"},
		{"Java", "Scala", "Perl"},
		{"C++", "Go", "Rust"},
		{"Crystal", "OCAML"}}
	for _, v := range langs {
		for _, lang := range v {
			fmt.Println(lang, " hi ")
		}
	}
}
