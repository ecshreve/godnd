package main

import (
	"github.com/ecshreve/godnd/internal/genny"
)

func main() {
	toGen := []string{
		"ability-scores",
		"skills",
		"proficiencies",
		"languages",
		"classes",
		"conditions",
		"damage-types",
		"magic-schools",
		"equipment-categories",
		"equipment",
		// "features", type name collision with something from the `conditions` endpoint
		"magic-items",
		// "monsters", this doc includes information on the filter query params, need to handle it
		"races",
		"subraces",
		"subclasses",
	}
	genny.GenerateTypes(toGen)
}
