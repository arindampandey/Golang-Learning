package main

import (
	"fmt"
	"slices"
)

//slices -> dynamic
// Msst used concept of Go

func main() {
	//Uninitialized slice is nil
	var num []int
	// fmt.Println(num==nil)
	fmt.Println(len(num))

	// var num1 = make([]int, 2, 5) //Here 2 is length and 5 is capacity
	// fmt.Println(num1)

	//Capacity
	// fmt.Println(cap(num1))

	// num1 = append(num1, 1)
	// num1 = append(num1, 2)
	// num1 = append(num1, 3)
	// num1 = append(num1, 4)
	// num1 = append(num1, 5)
	// fmt.Println(num1)

	// fmt.Println(cap(num1))

	//For making empty slice the length should be 0 in make function


	var num1 = make([]int, 0, 5)
	num1 = append (num1,2)
	var num2 = make([]int,len(num1))

	//Copy function

	copy(num2,num1)
	fmt.Println(num1,num2)


	//Slice Operator 
	var slice = []int{1,2,3,4}
	fmt.Println(slice[0:2])
	fmt.Println(slice[:3])
	fmt.Println(slice[0:])


	//Slice
	var n1 = []int{1,2,3}
	var n2 = []int{1,2,3}

	fmt.Println(slices.Equal(n1,n2))


	//2D Slice

	var s1 = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(s1)
}
