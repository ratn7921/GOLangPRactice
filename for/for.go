package main

//for is only construct in go for looping
import "fmt"

func main() {
	//while loop ex under for syntax:-
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i=i+1
	}

	// //classic for loop

	for i := 0; i <= 100; i++ {
		//	break
		if i%2 == 0 {
			fmt.Println(i, "is an even number")

			fmt.Print()
		} else {
			fmt.Println(i, "this is an odd number ")
		}

	}
	
	
//Go lang update 1.22 range feature ex:-


	for i:= range 30{
		if i%2==0{
			fmt.Println(i)
		}
	}
}
