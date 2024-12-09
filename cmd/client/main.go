package main

import (
	"context"
	"fmt"

	client "github.com/ecshreve/godnd/internal/api"
	"github.com/ecshreve/godnd/internal/api/models"
	"github.com/kr/pretty"
)

func main() {
	// Depends on a running api server at http://localhost:3003
	cl, err := client.NewClientWithResponses("http://localhost:3003")
	if err != nil {
		panic(err)
	}

	res, err := cl.GetApiEndpointWithResponse(context.Background(), models.GetApiEndpointParamsEndpoint("ability-scores"))
	if err != nil {
		panic(err)
	}

	fmt.Println(res.JSON200)
	pretty.Print(res.JSON200)
}
