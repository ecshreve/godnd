package main

import (
	"context"
	"log"

	"github.com/ecshreve/godnd/pkg/api"
	"github.com/kr/pretty"
)

func main() {
	cc := api.NewClient()
	ctx := context.Background()
	r, err := cc.ConditionByIndex(ctx, "blinded")
	//r, err := cc.ConditionAll(context.Background())
	//r, err := cc.AbilityScoreByIndex(ctx, "str")
	//r, err := cc.AbilityScoreAll(ctx)

	if err != nil {
		log.Fatal(err)
	}

	pretty.Print(r)
}
