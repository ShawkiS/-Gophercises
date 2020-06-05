package main

import (
	"fmt"
	"github.com/Gophercises/math"
)


func main()  {
	add := math.Add(1,3)

	if add != 4 {
		msg := fmt.Sprintf("Fail wanted 4 but %d", add)
		panic(msg)
	}
	
	sum := math.Sum([]int {1,3,5})	
	if sum != 9 {
		msg := fmt.Sprintf("Fail wanted 9 but %d", sum)
		panic(msg)
	}

	fmt.Printf("Success")
}