package Greet

import (
	"fmt"
)

func ExampleHello()  {
	gretting, err := Hello("Jon")	
	if err != nil {
		panic(err)
	}
	fmt.Println(gretting)
	
}