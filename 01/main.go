package main

import (
	// Thư viện chuẩn của Google
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Hello Jun, enter your city: ")
	// city, _ := reader.ReadString('\n')
	// fmt.Print("You live in " + city)

	var (
		err    error
		height float64
		weight float64
	)

	height, err = readDataFromKeyboard("Enter your height: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	weight, err = readDataFromKeyboard("Enter your weight: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	bmi := CalculateBMI(height, weight)
	fmt.Println("Chỉ số bmi của bạn = ", bmi)
}

/*
Đọc data từ bàn phím và convert string sang float
*/
func readDataFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")

	if result, err = strconv.ParseFloat(str, 64); err != nil {
		return 0.0, err
	}

	return result, nil
}

func CalculateBMI(height float64, weight float64) (index float64) {
	return weight / (height * height)
}
