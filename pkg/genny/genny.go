package genny

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/kr/pretty"
)

var tst = `this is somethings

type AbilityScore {
  _id: ObjectId
  desc: [String]
  full_name: Int
  index: Float
  name: Bool
  skills: [AbilityScoreSkill]
  url: String
}

type AbilityScore {
	_id: ObjectId
	desc: [String]
	full_name: Int
	index: Float
	name: Bool
	skills: [AbilityScoreSkill]
	url: String
}

this is something else`

var GqlToGoScalarTypeMap = map[string]string{
	"Int":    "int32",
	"Float":  "float64",
	"String": "string",
	"Bool":   "bool",
}

var typeRe = regexp.MustCompile(`(?s)type [^\{]+ \{[^\}]+\}`)
var typeDeclRe = regexp.MustCompile(`(type [^\{]+) (\{)`)
var idRe = regexp.MustCompile(`.*_id: .*\n`)
var fieldsRe = regexp.MustCompile(`(?s)\{\s*(.*)\s*\n\}`)
var fieldsSplitRe = regexp.MustCompile(`\n[ ]*`)
var scalarRe = regexp.MustCompile(`Int|Float|String|Bool`)
var listRe = regexp.MustCompile(`\[(.*)\]`)
var fieldNameRe = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]*:`)
var doubleSpaceRe = regexp.MustCompile(`  `)

func DoIt() {
	schemaFile, _ := os.Open("/Users/ericshreve/github.com/godnd/static/schema.graphql")
	b, _ := ioutil.ReadAll(schemaFile)

	allGqlTypes := typeRe.FindAllString(string(b), -1)
	allGoTypes := make([]string, len(allGqlTypes))

	for i, typeStr := range allGqlTypes {
		typeStr = typeDeclRe.ReplaceAllString(typeStr, "$1 struct $2")
		typeStr = idRe.ReplaceAllString(typeStr, "")
		typeStr = scalarRe.ReplaceAllStringFunc(typeStr, func(s string) string { return GqlToGoScalarTypeMap[s] })
		typeStr = listRe.ReplaceAllString(typeStr, "[]$1")
		typeStr = fieldNameRe.ReplaceAllStringFunc(typeStr, func(s string) string { return strings.TrimRight(snakeToCamel(s), ":") })
		typeStr = doubleSpaceRe.ReplaceAllString(typeStr, "\t")
		typeStr = typeStr + "\n"

		allGoTypes[i] = typeStr
	}

	pretty.Printf("%v", allGoTypes)
}
