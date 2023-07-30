package main

import (
	"fmt"
)

// Golang Syntax

// Variables
// primitive data types
var num int = 10
var decimalnum float32 = 10.2
var decimalnum2 float64 = 10.2
var str string = "Name"
var boolvar bool = true

const constvar string = "const"

// reference data types

// Array

func Arraytesting() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	// I is the index, v is the value
	for i, v := range arr1 {
		fmt.Println(i, v)
	}
	fmt.Println("The size of this array is ", len(arr2))
}

// slice - dynamic array ArrayList
func Slicetesting() {
	arr := []int{1, 2, 3, 4, 5}
	// slice
	slice := arr[0:3]
	fmt.Println(slice)
	// append
	slice = append(slice, 5)
	fmt.Println(slice, len(slice))

	// delete element from slice
}

// map
func Maptesting() {
	people := make(map[string]string)
	// add
	people["name"] = "Jialong"
	people["age"] = "25"
	// iterate map
	for key, value := range people {
		fmt.Println(key, value)
	}
	// size
	fmt.Println("The size of this map is ", len(people))
	// update
	people["name"] = "JL"
	// delete
	delete(people, "age")
	// how to check if a map has the key
	key, ok := people["age"]
	if ok {
		fmt.Println("The key is ", key)
	} else {
		fmt.Println("The key is not found")
	}
}

// struct - class / Object without method

//type Rectangle struct {
//	length int
//	width  int
//}

//// Area receiver function - method
//func (r Rectangle) Area() int {
//	return r.length * r.width
//}
//func (r Rectangle) Perimeter() int {
//	return 2 * (r.length + r.width)
//}

//func Structtesting() {
//	// instance
//	rec := Rectangle{length: 10, width: 20}
//	fmt.Println(rec.Area())
//	fmt.Println(rec.Perimeter())
//}

// function
// Pass by Value, Pass by Reference
type account struct {
	balance float64
}

var testfunc = func() {
	fmt.Println("Hello World")
}

func testFuncwithParams(anotherFunc func()) {
	anotherFunc()
}
func enclouser() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

// A method with a point receiver
func (ac *account) Deposit(amount int) {
	ac.balance += float64(amount)
}
func (ac *account) Withdraw(amount int) {
	ac.balance -= float64(amount)
}

func add(amount int) int {
	return amount + 10
}
func funcTesting() {
	//MyAccount := account{balance: 100}
	//MyAccount2 := &account{balance: 100}
	//fmt.Println(MyAccount)
	//fmt.Println(MyAccount.balance)
	//MyAccount.Deposit(100)
	//fmt.Println(MyAccount.balance)
	//MyAccount.Withdraw(50)
	//fmt.Println(MyAccount.balance)

	// pass by value
	amount := 10
	fmt.Println(add(amount))
	fmt.Println(amount)
}

// interface

type Rectangle struct {
	Width, Height int
}
type Circle struct {
	Radius int
}

type Shape interface {
	Area() int
	Perimter() int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimter() int {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() int {
	return c.Radius * c.Radius
}
func (c Circle) Perimter() int {
	return 2 * 3 * c.Radius
}

func PrintShape() {
	r := Rectangle{Width: 10, Height: 20}
	c := Circle{Radius: 10}
	fmt.Println(c.Area())
	fmt.Println(c.Perimter())
	fmt.Println(r.Area())
	fmt.Println(r.Perimter())
}

// OOD interface

// struct composition
type Engine struct {
	HP int
}

type Car struct {
	Make  string
	Model string
	Year  int
	Engine
}

// interface implementation

func oodTesting() {
	// create a instance from a composition struct type
	myCar := Car{Make: "Toyota", Model: "Camry", Year: 2020, Engine: Engine{HP: 200}}
	fmt.Println(myCar)
}

// Go routines

//TODO

// Package Management - Go Modules

// Error
func doSomething() (result string, err error) {
	// Try to do something that could fail
	//	result, err = someOperation()
	// If an error occurred, return the error
	if err != nil {
		return "", err
	}

	// If everything went fine, return the result and a nil error
	return result, nil
}

// driver
func main() {
	//fmt.Println("Hello World")
	//Maptesting()
	//Arraytesting()
	//Slicetesting()
	//Structtesting()
	//funcTesting()
	//PrintShape()
	//oodTesting()
	testFuncwithParams(testfunc)

	testenclouser := enclouser()
	fmt.Println(testenclouser())
}
