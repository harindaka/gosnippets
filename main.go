package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	var exampleToRun string
	if len(os.Args) == 2 {
		exampleToRun = os.Args[1]
	}

	switch exampleToRun {
	case "hello_world":
		helloWorld()

	case "random_number_generator":
		generateRandomNumbers()

	default:
		fmt.Println("Usage: gosnippets (snippet name)")
		fmt.Println("eg.: gosnippets hello_world")
	}

}
