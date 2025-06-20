package main

import (
	"fmt"
)

func main() {
	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
    // case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("five")
	// default:
	// 	fmt.Println("out of case")	

	// }


	// switch time.Now().Weekday(){
	// case time.Saturday, time.Friday:
	// 	fmt.Println("half day")
    // case time.Sunday:
	// 	fmt.Println("off day")		
	// default:
	// 	fmt.Println(" workday")
	
	// }
	


	// powerfull : the type switch


	whoAmi := func (i interface{})  {
		switch i:= i.(type){
			case int:
				fmt.Println("its an integer")
			case string:
				fmt.Println("its a string")
			case bool:
				fmt.Println("its a boolean")
			default:
				fmt.Println(" other", i)			
			
		}
	}
	whoAmi(5)
	whoAmi("hello")
	whoAmi(true)
	whoAmi(3.14)
	whoAmi([]int{1, 2, 3})
	whoAmi(map[string]int{"one": 1, "two": 2})

}