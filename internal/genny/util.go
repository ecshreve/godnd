package genny

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/samsarahq/go/oops"
)

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

var gqlToGoScalarTypeMap = map[string]string{
	"Int":    "int32",
	"Float":  "float64",
	"String": "string",
	"Bool":   "bool",
}

var gqlTypeRe = regexp.MustCompile(`(?s)type [^\{]+ \{[^\}]+\}`)

func getGqlSchema() (*string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, oops.Wrapf(err, "error getting current working directory")
	}

	schemaFilePath := fmt.Sprintf("%s/static/schema.graphql", wd)
	schemaFile, err := os.Open(schemaFilePath)
	if err != nil {
		return nil, oops.Wrapf(err, "error opening schema file at path: %s", schemaFilePath)
	}

	b, err := ioutil.ReadAll(schemaFile)
	if err != nil {
		return nil, oops.Wrapf(err, "error reading schema file")
	}

	str := string(b)
	return &str, nil
}

func (c *Converter) GetGqlTypeFromSchema(typeName string) string {
	gqlTypeRe := regexp.MustCompile(fmt.Sprintf(`(?s)type %s \{[^\}]+\}`, typeName))
	fmt.Println(gqlTypeRe)
	return gqlTypeRe.FindString(c.GqlSchema)
}

func getReplacers() []Action {
	actions := []Action{
		Action{
			Description: "gql to go type signature -- `type <TYPENAME> {` -> `type TYPENAME struct {",
			Matcher:     regexp.MustCompile(`(type [^\{]+) (\{)`),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllString(s, "$1 struct $2") },
		},
		Action{
			Description: "remove auto generated _id field -- `  _id: <SOMEVALUE>` -> ``",
			Matcher:     regexp.MustCompile(`.*_id: .*\n`),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllString(s, "") },
		},
		Action{
			Description: "convert gql scalar to go scalar -- `Int` -> `int32`",
			Matcher:     regexp.MustCompile(`Int|Float|String|Bool`),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllStringFunc(s, gqlScalarToGo) },
		},
		Action{
			Description: "convert gql list to go slice -- `[TYPE]` -> `[]TYPE`",
			Matcher:     regexp.MustCompile(`\[(.*)\]`),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllString(s, "[]$1") },
		},
		Action{
			Description: "convert gql field name to go field name -- `foo_bar` -> `FooBar`",
			Matcher:     regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]*:`),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllStringFunc(s, gqlFieldNameToGo) },
		},
		Action{
			Description: "uppercase url field names -- `Url` -> `URL`",
			Matcher:     regexp.MustCompile(`\s([uU][rR][lL])\s`),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllString(s, " URL ") },
		},
		Action{
			Description: "replace double spaces with tabs",
			Matcher:     regexp.MustCompile(`  `),
			Replacer:    func(s string, re *regexp.Regexp) string { return re.ReplaceAllString(s, "\t") },
		},
	}

	return actions
}

func doActionOnString(a Action, s string) string {
	return a.Replacer(s, a.Matcher)
}

func toStringPtr(s string) *string   { return &s }
func fromStringPtr(s *string) string { return *s }

func gqlScalarToGo(s string) string    { return gqlToGoScalarTypeMap[s] }
func gqlFieldNameToGo(s string) string { return strings.TrimRight(snakeToCamel(s), ":") }

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
