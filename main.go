package main

import (
	"fmt"
	"golearn/public"
)

func main() {
	fmt.Println(public.Stringtoarray("hello"))
	fmt.Println(public.IntToArray(12345))
	fmt.Println(public.FloatToArray(123.45)) //Not the best idea to use float64

	//Using Car struct
	Mycar := public.Car{
		Name: "Toyota",
		Year: 2020,
	}
	fmt.Println(Mycar.GetName())
	fmt.Println(Mycar.GetYear())

	//Using House interface
	Myhouse := public.FamilyHouse{
		Address: "123 Main St",
		Cost:    100000,
	}
	fmt.Println(Myhouse.Location())
	fmt.Println(Myhouse.Price())
}

