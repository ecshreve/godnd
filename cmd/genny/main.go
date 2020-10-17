package main

import (
	"fmt"
	"log"

	"github.com/ecshreve/godnd/internal/genny"
)

func main() {
	c, err := genny.NewConverter()
	if err != nil {
		log.Fatalf("exited unexpectedly\n\n%v\n", err)
	}

	fmt.Printf("\nXXXXX\n\n")
	for _, supportedType := range genny.SupportedTypes {
		gqlTypeStr := c.GetGqlTypeFromSchema(supportedType)
		fmt.Println(gqlTypeStr)
		fmt.Println("~~~~~")
		goTypeStr, err := c.ConvertGqlToGoType(gqlTypeStr)
		if err != nil {
			log.Fatalf("error converting gql to json\n\n%v\n", err)
		}
		fmt.Println(goTypeStr)
		fmt.Printf("\nXXXXX\n\n")
	}
}
