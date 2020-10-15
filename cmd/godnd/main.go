package main

import (
	"context"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/ecshreve/godnd/pkg/client"
)

func main() {
	ctx := context.Background()
	scs := spew.ConfigState{Indent: "\t"}

	c := client.NewClient()
	res, err := c.GetAllDamageTypes(ctx)
	if err != nil {
		log.Fatalf("bloop: %+v", err)
	}
	scs.Dump(res)
}
