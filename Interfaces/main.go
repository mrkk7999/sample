package main

import "fmt"

type MyInterface interface {
	Add() int
	Substract() int
}
type Values struct {
	a int
	b int
}

func (v *Values) Add() int {
	return v.a + v.b
}

func (v Values) Substract() int {
	return v.b - v.a
}

func Operations(temp MyInterface) {
	fmt.Println(temp.Add())
	fmt.Println(temp.Substract())
}
func main() {
	v := &Values{
		a: 7,
		b: 9,
	}
	Operations(v)

	//v.Add()
	//v.Substract()
}
