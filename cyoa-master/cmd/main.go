package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"../cyoa"
)



func main()  {
	filename := flag.String("file", "gopher.json", "The JSON file that has the stories")
	flag.Parse()
	fmt.Printf("Using the story file in %s.\n", *filename)
	
	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	d := json.NewDecoder(f)

	var story cyoa
	if err:= d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v", story)
	 
}