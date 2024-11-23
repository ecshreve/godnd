package main

import (
	"context"
	"fmt"

	client "github.com/ecshreve/godnd/internal/client/generated"
	"github.com/kr/pretty"
)

func main() {
	// Depends on a running api server at http://localhost:3000
	cl, err := client.NewClientWithResponses("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	res, err := cl.GetAPIAbilityScoresIndexWithResponse(context.Background(), "cha")
	if err != nil {
		panic(err)
	}

	fmt.Println(res.JSON200)
	pretty.Print(res.JSON200)
}
