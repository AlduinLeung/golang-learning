package main

import "fmt"

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
	// i is the index, v is the value
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

// struct - class / Object

// function

// interface

//loop if else

// Pass by Value, Pass by Reference

// OOD interface

// Go routines

// Package Management - Go Modules

// driver
func main() {
	//fmt.Println("Hello World")
	//Maptesting()
	//Arraytesting()
	Slicetesting()
}
