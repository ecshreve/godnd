package main

import (
	"github.com/ecshreve/godnd/internal/genny"
)

func main() {
	toGen := []string{
		"ability-scores",
		"classes",
		"conditions",
		"damage-types",
		"equipment-categories",
		"equipment",
		// "features", type name collision with something from the `conditions` endpoint
		"languages",
		"magic-items",
		"magic-schools",
		// "monsters", this doc includes information on the filter query params, need to handle it
		"proficiencies",
		"races",
		// "rules",
		// "rule-sections",
		"skills",
		// "spellcasting",
		// "spells",
		// "starting-equipment",
		"subclasses",
		"subraces",
		"traits",
		"weapon-properties",
	}

	genny.GenerateTypes(toGen)
}
