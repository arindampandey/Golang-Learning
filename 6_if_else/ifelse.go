package main

import "fmt"

func main() {
	// age := 18

	// if age >= 18 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are a kid")
	// }

	// age := 26

	// if age>= 18{
	// 	fmt.Println("Person is an adult")
	// } else if (age>=12){
	// 	fmt.Println("Person is a teenager")
	// } else{
	// 	fmt.Println("Person is a kid")
	// }

	var role = "admin"
	var haspermit = false

	// OR Operator
	if role == "admin" || haspermit {
		fmt.Println("You are eligible !!")
	}

	// AND Operator
	if role == "admin" && haspermit {
		fmt.Println("You are not eligible !!")
	}


	if num:=5 ; num<10{
		fmt.Println("it is a small number",num)
	}else if num > 2{
		fmt.Println("Hiiii")
	}


	//NOTE: Go does not have ternary operator, you have to use if else for such problems.

}
