package main

import "fmt"

//iterating over data structures in go
func main() {

	// Range with slices
	// nums := []int {1, 2, 3, 4, 5}

	// sum := 0
	// for i, num := range nums{
	// 	sum =  sum +num
	// 	fmt.Println(num,"index:",i)

	// }

	// for i:=0; i<len (nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// range with maps

	// m := map[string]string{"fname": "John", "lastName": "Doe"}

	// for k, v := range m {
	// 	fmt.Println("key: = ", k ,"|", "value: = ", v)
	// }




	for i, c := range "golang"{
		fmt.Println(i,c)
	}
}
