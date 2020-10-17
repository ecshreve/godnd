// +build ignore

package genny

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var typeRe = regexp.MustCompile(`(?s)type [^\{]+ \{[^\}]+\}`)
var typeDeclRe = regexp.MustCompile(`(type [^\{]+) (\{)`)
var idRe = regexp.MustCompile(`.*_id: .*\n`)
var fieldsRe = regexp.MustCompile(`(?s)\{\s*(.*)\s*\n\}`)
var fieldsSplitRe = regexp.MustCompile(`\n[ ]*`)
var scalarRe = regexp.MustCompile(`Int|Float|String|Bool`)
var listRe = regexp.MustCompile(`\[(.*)\]`)
var fieldNameRe = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]*:`)
var doubleSpaceRe = regexp.MustCompile(`  `)
var apiReferenceRe = regexp.MustCompile(`.*\{\s+Index string\s+Name string\s+Url string\s+\}.*`)
var goTypeNameRe = regexp.MustCompile(`type (.*) struct`)

func (p *Parser) GetRawGoTypeStrings() {
	schemaFile, _ := os.Open("/Users/ericshreve/github.com/godnd/static/schema.graphql")
	b, _ := ioutil.ReadAll(schemaFile)

	allGqlTypes := typeRe.FindAllString(string(b), -1)
	var allGoTypes []string

	for _, typeStr := range allGqlTypes {
		typeStr = typeDeclRe.ReplaceAllString(typeStr, "$1 struct $2")
		typeStr = idRe.ReplaceAllString(typeStr, "")
		typeStr = scalarRe.ReplaceAllStringFunc(typeStr, func(s string) string { return GqlToGoScalarTypeMap[s] })
		typeStr = listRe.ReplaceAllString(typeStr, "[]$1")
		typeStr = fieldNameRe.ReplaceAllStringFunc(typeStr, func(s string) string { return strings.TrimRight(snakeToCamel(s), ":") })
		typeStr = doubleSpaceRe.ReplaceAllString(typeStr, "\t")
		typeStr = typeStr + "\n"

		if IsCommonType(typeStr) || IsGqlInternalType(typeStr) {
			continue
		}

		typeName := goTypeNameRe.FindStringSubmatch(typeStr)[1]
		embeddedTypeNameRe := regexp.MustCompile(fmt.Sprintf(`%s(\S)`, typeName))
		typeStr = embeddedTypeNameRe.ReplaceAllString(typeStr, "$1")

		embeddedChoiceRe := regexp.MustCompile(`\[\]\S+_choice`)
		typeStr = embeddedChoiceRe.ReplaceAllString(typeStr, "[]Choice")

		embeddedForeignTypeName := regexp.MustCompile(`([a-zA-Z]+_[a-zA-Z]+)+`)
		typeStr = embeddedForeignTypeName.ReplaceAllStringFunc(typeStr, snakeToCamel)

		goTypeName := goTypeNameRe.FindStringSubmatch(typeStr)[1]
		typeStr = fmt.Sprintf("// %s is a %s\n", goTypeName, goTypeName) + typeStr
		typeStr = strings.ReplaceAll(typeStr, "Url", "URL")
		allGoTypes = append(allGoTypes, typeStr)
	}

	p.RawGoTypeStrings = allGoTypes
}
