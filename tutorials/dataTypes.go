package main

import "fmt"

func main() {
	var data int
	data = 10
	data1 := 100
	data = 1000
	Data := 1000

	fmt.Println(data)
	fmt.Println(data1)
	fmt.Println(Data)

	//data type
	// 1. Primitive
	// int,float64,string,bool,complex
	// 2. Non-Primitive
	// struct map chan interface array slice rune reflect pointer
	name := `rahul`
	fmt.Println(name)
	// const pi=3.14
	const pi float32=3.14
	const l int =100
	//l=10
	fmt.Print(l)
	fmt.Println(pi)
}