package genny

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/samsarahq/go/oops"
)

var GqlToGoScalarTypeMap = map[string]string{
	"Int":    "int32",
	"Float":  "float64",
	"String": "string",
	"Bool":   "bool",
}

var (
	doubleSpaceRe      = regexp.MustCompile(`  `)
	gqlTypeRe          = regexp.MustCompile(`(?s)type [^\{]+ \{[^\}]+\}`)
	gqlTypeSignatureRe = regexp.MustCompile(`(type [^\{]+) (\{)`)
	gqlIDFieldRe       = regexp.MustCompile(`.*_id: .*\n`)
	gqlFieldNameRe     = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]*:`)
	gqlScalarRe        = regexp.MustCompile(`Int|Float|String|Bool`)
	gqlListRe          = regexp.MustCompile(`\[(.*)\]`)
)

func GenerateType(typeStr string) (string, error) {
	valid := gqlTypeRe.MatchString(typeStr)
	if !valid {
		return "", oops.Errorf("invalid gql type string")
	}
	fmt.Println(typeStr)

	typeStr = gqlTypeSignatureRe.ReplaceAllString(typeStr, "$1 struct $2")
	typeStr = gqlIDFieldRe.ReplaceAllString(typeStr, "")
	typeStr = gqlScalarRe.ReplaceAllStringFunc(typeStr, func(s string) string { return GqlToGoScalarTypeMap[s] })
	typeStr = gqlListRe.ReplaceAllString(typeStr, "[]$1")
	typeStr = gqlFieldNameRe.ReplaceAllStringFunc(typeStr, func(s string) string { return strings.TrimRight(snakeToCamel(s), ":") })
	typeStr = doubleSpaceRe.ReplaceAllString(typeStr, "\t")

	return typeStr, nil
}
