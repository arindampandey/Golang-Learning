package main
import "fmt"

const age = 30
// age := 25 (Not allowed !) 

func main(){
	const name = "golang" 
	// name = "JSS" (Not allowed !) 

	fmt.Println(age)

	const(
		port = 9999
		host = "Local"
	)

	fmt.Println(port,host)
}
