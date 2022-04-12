package main

import "fmt"

func testJapaneseDictionary() {
	dict := map[string]string{}
	dict["Sayonara"] = "Tạm biệt"
	dict["Onegai"] = "Làm ơn"
	dict["Konichiwa"] = "Chào buổi chiều"
	dict["Ohayo"] = "Chào buổi sáng"
	for key, value := range dict {
		fmt.Println(key, " : ", value)
	}
}

func testMapInterface() {
	config := map[string]interface{}{
		"host": "localhost",
		"port": 8080,
		"user": "junkyu",
		"pass": "123@456",
		"TLS":  true,
	}

	for key, value := range config {
		fmt.Println(key, " : ", value)
	}
}
