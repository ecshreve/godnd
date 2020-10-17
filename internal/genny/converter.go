package genny

import (
	"fmt"
	"regexp"

	"github.com/samsarahq/go/oops"
)

type Action struct {
	Description  string
	Matcher      *regexp.Regexp
	Replacer     func(string, *regexp.Regexp) string
	ReplacerFunc *func(string) string
}

type Converter struct {
	GqlSchema string
	Replacers []Action
}

func NewConverter() (*Converter, error) {
	gqlSchema, err := getGqlSchema()
	if err != nil {
		return nil, oops.Wrapf(err, "error getting gql schema")
	}

	return &Converter{
		GqlSchema: *gqlSchema,
		Replacers: getReplacers(),
	}, nil
}

func (c *Converter) ConvertGqlToGoType(typeStr string) (string, error) {
	valid := gqlTypeRe.MatchString(typeStr)
	if !valid {
		return "", oops.Errorf("invalid gql type string\n%v", typeStr)
	}

	// Do all the simple conversions first.
	for _, a := range c.Replacers {
		typeStr = doActionOnString(a, typeStr)
	}

	// Strip out embedded field type names that are prepended with the struct type.
	goTypeNameRe := regexp.MustCompile(`type (.*) struct`)
	goTypeName := goTypeNameRe.FindStringSubmatch(typeStr)[1]
	embeddedTypeNameRe := regexp.MustCompile(fmt.Sprintf(`%s(\S)`, goTypeName))
	typeStr = embeddedTypeNameRe.ReplaceAllString(typeStr, "$1")

	return typeStr, nil
}
