package main

import "fmt"

// In go there is only for loop is present for looping
func main() {

	// While loop using for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	//Classic for loops
	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	//Range
	for i := range 4 {
		fmt.Println(i)
	}

}
