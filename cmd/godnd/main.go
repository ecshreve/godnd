package main

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/ecshreve/godnd/pkg/client"
)

func main() {
	ctx := context.Background()
	scs := spew.ConfigState{Indent: "\t"}

	c := client.NewClient()
	res, _ := c.GetClassByIndex(ctx, "sorcerer")
	scs.Dump(res)
}
