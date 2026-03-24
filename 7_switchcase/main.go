package main

import (
	"fmt"
)

func main() {
	//Simple Switch
	// i := 2

	// switch i {
	// case 1:
	// 	fmt.Println("One")

	// case 2:
	// 	fmt.Println("Two")

	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Invalid")
	// }

	//Multiple Condition Switch

	// switch time.Now().Weekday(){
	// case time.Saturday,time.Sunday:
	// 	fmt.Println("It is weekend")
	// default:
	// 	fmt.Println("It is a workday")
	// }

	//Type Switch

	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It is integer")
		case string:
			fmt.Println("It is a string")
		case bool:
			fmt.Println("It is a boolean")
		default:
			fmt.Println("Other",t)
		}
	}

whoamI("goo")
whoamI(66)
whoamI(true)
whoamI(99.9)
}
