package main

import "fmt"

type Person struct {
	Id       string
	FullName string
	Email    string
	Addr     []Address
}

type Address struct {
	Location string
	City     string
	Country  string
}

func (p Person) String() string {
	return fmt.Sprintf("%s : %s : %s : %s : %s : %s", p.Id, p.FullName, p.Email,
		p.Addr[0].Location, p.Addr[0].City, p.Addr[0].Country)
}

func main() {
	type personRequest struct {
		FullName string
		Email    string
		Location string
		City     string
		Country  string
	}

	perRequest := personRequest{
		FullName: "Jun Nguyen",
		Email:    "junnguyen1011@gmail.com",
		Location: "98A Nguy Nhu Kon Tum",
		City:     "Ha Noi",
		Country:  "Viet Nam",
	}

	person := Person{
		Id:       "ox-11",
		FullName: perRequest.FullName,
		Email:    perRequest.Email,
		Addr: []Address{
			Address{
				Location: perRequest.Location,
				City:     perRequest.City,
				Country:  perRequest.Country,
			},
		},
	}

	fmt.Println(person)
}
