package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("genny")

	typePtr := flag.String("type", "all", "specify the type to generate")
	flag.Parse()

	fmt.Printf("%v\n", *typePtr)
}
