package main
import "fmt"

func main(){
	var nums [5]int

	nums[0]=1
	fmt.Println(nums[1])
	fmt.Println(nums) //When no values are given to indexes so they are initialized with zeroed values

	// Array length
	// fmt.Println(len(nums))


	arr := [4]int{1,2,3,4}
	fmt.Println(arr)

	//2D Arrays

	arr2 := [2][2]int{{2,4},{6,8}}
	fmt.Println(arr2)
}