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
	}
	genny.GenerateTypes(toGen)
}
