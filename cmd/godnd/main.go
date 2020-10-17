package main

import (
	"github.com/ecshreve/godnd/pkg/genny"
)

func main() {
	parser := genny.Parser{}
	parser.ParseGqlSchema()
}
