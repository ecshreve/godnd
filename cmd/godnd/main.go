package main

import (
	"context"
	"log"

	"github.com/ecshreve/godnd/pkg/api"
	"github.com/kr/pretty"
)

func main() {
	cc := api.NewClient()
	r, err := cc.ConditionsAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	pretty.Print(r)
}
