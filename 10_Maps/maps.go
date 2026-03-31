package main

import (
	"fmt"
	"maps"
)

// map(hash,object,dict)
func main() {

	//creating map
	m := make(map[string]string)

	//setting an element
	m["name"] = "Arindam"
	m["area"] = "MP"

	//getting an element
	fmt.Println(m["name"], m["area"])
	fmt.Println(m["phone"]) //If key does not exists in the map then it returns zero value

	m1 := make(map[string]int)
	m1["age"] = 30
	m1["num"] = 9
	// fmt.Println(m1["phone"])

	// fmt.Println(len(m1))

	// delete(m1,"num")
	// fmt.Println(m1)

	// clear(m1)

	m2 := map[string]int{"price": 200, "age": 30}
	fmt.Println(m2)

	// v, ok := m2["price"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("All Ok")
	// } else {
	// 	fmt.Println("Not Ok")
	// }

	m3 := map[string]int{"price": 200, "age": 30}
	fmt.Println(maps.Equal(m2,m3))

}
