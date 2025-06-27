package main

import (
	"fmt"
	"maps"
)

func main() {

	// m := make(map[string]string)
    // n:= make(map[string]int)

    // n["age"]= 30
	// n["height"]= 180
	// m["name"] = "ratnakar"
    // m["area"]= "bangalore"
	// delete(n, "height")
	// clear(m)
	// clear(n)
    // fmt.Println(n["age"], n["height"])
	// fmt.Println(m["name"], m["area"])










	// m :=map[string]int{"price":100,"iphone":1000}

	// k, ok := m["prie"]
	// fmt.Println(k)
	// if ok {
	// 	fmt.Println("all ok")
	// } else{
	// 	fmt.Println("not ok")
	// }
	// fmt.Println(m)




	m1 :=map[string]int{"price":100,"phones":10}
	m2 :=map[string]int{"price":100,"phones":10}

	fmt.Println(maps.Equal(m1,m2))
}