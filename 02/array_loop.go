package main

import "fmt"

func indexLoop() {
	cars := []string{"Toyota", "BMW", "Mercedes"}
	for i := 0; i < len(cars); i++ {
		fmt.Println(i, cars[i])
	}
}

func whileLoop() {
	cars := [3]string{"Toyota", "BMW", "Mercedes"}
	i := 0
	for i < len(cars) {
		fmt.Println(i, cars[i])
		i++
	}
}

func rangeLoop() {
	cars := []string{"Mini Cooper", "Audi A6", "Bugatti Chiron"}
	for index, car := range cars {
		fmt.Println(index, car)
	}
}

/*
Lệnh defer sẽ đưa câu lệnh vào stack chờ đến khi thoát hàm thì mới thực hiện
Những lệnh chờ sau sẽ chồng lên lệnh chờ trước đấy, thế nên lúc thực thi
sẽ thực thi các lệnh chờ sau này, vì vậy áp dụng để reverse được.
*/
func reverseLoop() {
	cars := []string{"Mini Cooper", "Audi A6", "Bugatti Chiron"}
	for index, car := range cars {
		defer fmt.Println(index, car)
	}
}

func rawReverseLoop() {
	cars := []string{"Mini Cooper", "Audi A6", "Bugatti Chiron"}
	for i := len(cars) - 1; i >= 0; i-- {
		fmt.Println(i, cars[i])
	}
}
