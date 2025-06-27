package main

import (
	"fmt"
	"slices"
)

func main() {
	// //uninitialized slice
	// // A slice is a dynamically-sized, flexible view into the elements of an array.
	// // Unlike arrays, slices do not have a fixed size and can grow or shrink as needed.
	// // Slices are more powerful and convenient than arrays in Go.
	// // Slices are built on top of arrays, and they provide a way to work with a
	// // portion of an array without needing to copy the data.
	// // A slice is a descriptor of an array segment. It consists of three components:
	// var nums []int
	// fmt.Println(nums) // Output: []

	// var numbs =make([]int, 5)
	// fmt.Println(numbs) // Output: [0 0 0 0 0]
	// // The length of the slice is 0, and the capacity is 5.
	// fmt.Println(len(numbs)) // Output: 5
	// fmt.Println(cap(numbs)) // Output: 5
	// // The length of a slice is the number of elements it contains, and the capacity is the
	// // number of elements that can be stored in the underlying array without needing to allocate
	// // a new array.
	// // Slices can be created using the `make` function, which takes the type of the slice,
	// // the length, and the capacity as arguments.
	// // For example, to create a slice of integers with a length of 5 and a capacity of 10,
	// slice := make([]int, 5, 10)
	// fmt.Println(slice) // Output: [0 0 0 0 0]
	// fmt.Println(len(slice)) // Output: 5
	// fmt.Println(cap(slice)) // Output: 10
	// // You can also create a slice using a composite literal, which is a shorthand way to
	// // create a slice with initial values.
	// slice2 := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice2) // Output: [1 2 3 4 5]
	// // You can access and modify elements of a slice using the index operator.
	// slice2[0] = 10
	// fmt.Println(slice2) // Output: [10 2 3 4 5]
	// // You can also append elements to a slice using the `append` function.
	// slice2 = append(slice2, 6, 7, 8)
	// fmt.Println(slice2) // Output: [10 2 3 4 5 6 7 8]
	// // The `append` function returns a new slice with the added elements, and it may
	// // allocate a new array if the capacity of the original slice is exceeded.
	// // If the capacity of the slice is exceeded, a new array is allocated, and the
	// // elements of the original slice are copied to the new array.
	// // You can also create a slice from an existing array using the slice operator.
	// arr := [5]int{1, 2, 3, 4, 5}
	// slice3 := arr[1:4] // Create a slice from the array, starting
	// // at index 1 and ending at index 4 (exclusive).
	// fmt.Println(slice3) // Output: [2 3 4]
	// // The slice operator allows you to specify a range of indices to create a slice from an array.
	// // You can also create a slice from another slice using the slice operator.
	// slice4 := slice2[2:5] // Create a slice from the existing slice
	// fmt.Println(slice4) // Output: [3 4 5]

	//  var nums = make([]int, 5)
	//  nums = append(nums, 6, 7, 8, 9, 10)
	//  fmt.Println(nums) // Output: [0 0 0 0 0 6 7 8 9 10]
	// //  fmt.Println(cap(nums))

	// // Slices can also be used to create multi-dimensional arrays.
	//  multiSlice := make([][]int, 3) // Create a slice of slices (2D slice)
	//  for i := range multiSlice {
	// 	 multiSlice[i] = make([]int, 3) // Initialize each inner slice
	//  }
	//  multiSlice[0][0] = 1
	//  multiSlice[0][1] = 2
	//  multiSlice[0][2] = 3
	//  multiSlice[1][0] = 4
	//  multiSlice[1][1] = 5
	//  multiSlice[1][2] = 6
	//  multiSlice[2][0] = 7
	//  multiSlice[2][1] = 8
	//  multiSlice[2][2] = 9
	//  fmt.Println(multiSlice) // Output: [[1 2 3] [4 5 6] [7 8 9]]
	// // Slices can be used to create a slice of structs.

	// var slice1 []int
	// slice1 = append(slice1, 1, 2, 3, 4, 5)

	var nums = []int{1, 2, 3, 4, 5}
	fmt.Println(nums[2:5])

	// slicepackege

	var numbs1 =[]int{1,2,3,4}
	var numbs2 = []int{1,2,3,4}
	fmt.Println(slices.Equal(numbs1, numbs2)) // Output: true
}
