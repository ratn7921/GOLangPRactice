package main

import (
	"fmt"
	// "reflect"
)

var person int

func main() {
 
	result , err := GenrateRamdomkey(20)

	fmt.Println("Enter person age")
	fmt.Scan(&person)
	// person :=

	if person >= 18 {
		fmt.Println("person is adult")
	} else if person <= 12 {
		fmt.Println("person is kid")
	} else {
		fmt.Println("person is not an adult")
	}

	var role = "admin"
	
	var haspermissions bool
	if person > 30 {
		haspermissions = true
	} else {
		haspermissions = false
		return 
	}

	if role == "admin" || haspermissions   {
	if err != nil {
		fmt.Println("Error",err)
		return
	}
   fmt.Println("admin key:-",result)
}



}

