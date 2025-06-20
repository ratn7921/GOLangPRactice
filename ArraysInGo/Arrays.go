package main

import "fmt"

func main() {
	var nums [10]int

	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4
	nums[4] = 5
	nums[5] = 6	
	nums[6] = 7
	nums[7] = 8
	nums[8] = 9
	nums[9] = 10


	fmt.Println(nums)
	fmt.Println(len(nums))


var vals [4]bool
vals[2]= true
fmt.Println(vals)



var name [3]string
name[0] = "John"
fmt.Println(name)


nums2 := [5]int{1, 2, 3, 4, 5}
fmt.Println(nums2)

example :=[3][3]int{{1,3,5},{2,3,4}}
fmt.Println(example)
}