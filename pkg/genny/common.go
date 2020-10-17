package genny

import (
	"regexp"
)

type ApiReference struct {
	Index string
	Name  string
	Url   string
}

type ClassApiReference struct {
	Index string
	Class string
	Url   string
}

type Choice struct {
	Choose int32
	From   []ApiReference
	Type   string
}

type Cost struct {
	Quantity int32
	Unit     string
}

var commonTypeNameToRegex = map[string]*regexp.Regexp{
	"ApiReference":      regexp.MustCompile(`.*\{\s+Index string\s+Name string\s+Url string\s+\}.*`),
	"ClassApiReference": regexp.MustCompile(`.*\{\s+Index string\s+Class string\s+Url string\s+\}.*`),
	"Choice":            regexp.MustCompile(`.*\{\s+Choose int32\s+From \[\].*From\s+Type string\s+\}.*`),
	//"Cost":              regexp.MustCompile(``),
}

func IsCommonType(s string) bool {
	for _, re := range commonTypeNameToRegex {
		if re.MatchString(s) {
			return true
		}
	}
	return false
}

var gqlInternalTypeRegexes = []*regexp.Regexp{
	regexp.MustCompile(`Payload`),
	regexp.MustCompile(`Query`),
	regexp.MustCompile(`Mutation`),
}

func IsGqlInternalType(s string) bool {
	for _, re := range gqlInternalTypeRegexes {
		if re.MatchString(s) {
			return true
		}
	}
	return false
}
