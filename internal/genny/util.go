package genny

import (
	"bytes"
	"strings"
)

// DocFilePath is the file path to the directory containing the ejs files that
// hold documentation information.
//
// TODO: this should be updated to automatically pull the latest version of these
// files form the main repo.
const DocFilePath = "/Users/ericshreve/github.com/godnd/resources/partials"

var resourceToEndpoint = map[string]string{
	"ability-scores":       "/api/ability-scores",
	"classes":              "/api/classes",
	"conditions":           "/api/conditions",
	"damage-types":         "/api/damage-types",
	"equipment-categories": "/api/equipment-categories",
	"equipment":            "/api/equipment",
	"features":             "/api/features",
	"languages":            "/api/languages",
	"magic-items":          "/api/magic-items",
	"magic-schools":        "/api/magic-schools",
	"monsters":             "/api/monsters",
	"proficiencies":        "/api/proficiencies",
	"races":                "/api/races",
	"rules":                "/api/rules",
	"rule-sections":        "/api/rule-sections",
	"skills":               "/api/skills",
	"spellcasting":         "/api/spellcasting",
	"spells":               "/api/spells",
	"starting-equipment":   "/api/starting-equipment",
	"subclasses":           "/api/subclasses",
	"subraces":             "/api/subraces",
	"traits":               "/api/traits",
	"weapon-properties":    "/api/weapon-properties",
}

func snakeToCamel(s string) string {
	b := []byte(strings.TrimSpace(s))
	buffer := make([]byte, 0, len(s))

	tokens := bytes.Split(b, []byte("_"))
	for _, tok := range tokens {
		buffer = append(buffer, bytes.ToUpper(tok[:1])...)
		buffer = append(buffer, tok[1:]...)
	}

	return string(buffer)
}

func spaceSepToCamel(s string) string {
	b := []byte(strings.TrimSpace(s))
	buffer := make([]byte, 0, len(s))

	tokens := bytes.Split(b, []byte(" "))
	for _, tok := range tokens {
		buffer = append(buffer, bytes.ToUpper(tok[:1])...)
		buffer = append(buffer, tok[1:]...)
	}

	return string(buffer)
}
