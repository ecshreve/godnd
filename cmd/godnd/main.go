package main

import (
	"context"
	"fmt"
	"io"

	api "github.com/ecshreve/godnd/internal/api/generated"
)

func main() {
	cl, err := api.NewClient("http://localhost:3000")
	if err != nil {
		panic(err)
	}

	res, err := cl.GetApiAbilityScoresIndex(context.Background(), api.GetApiAbilityScoresIndexParamsIndex(""))
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
