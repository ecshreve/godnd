package main

import (
	"context"
	"log"

	"github.com/ecshreve/godnd/pkg/api"
	"github.com/kr/pretty"
)

func main() {
	cc := api.NewClient()
	//r, err := cc.ConditionByIndex(context.Background(), "blinded")
	r, err := cc.ConditionAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	pretty.Print(r)
	pretty.Print("hi")
}
