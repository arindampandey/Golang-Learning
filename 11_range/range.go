package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	for i, num := range nums {
		fmt.Println(num, i)
	}

	m := map[string]int{"Age":18,"Heght":6}

	for k,v := range m{
		fmt.Println(k,v)
	}

	//Unicode code point rune
	// starting type of rune 
	for i,c:= range "StringName"{
		// fmt.Println(i,c)
		fmt.Println(i,string(c))
	}

}
