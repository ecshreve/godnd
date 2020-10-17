package genny

import (
	"io/ioutil"
	"os"
)

type Parser struct {
	TypeNames        []string
	RawGoTypeStrings []string
	GqlSchema        string
}

func NewParser() *Parser {
	schemaFile, _ := os.Open("/Users/ericshreve/github.com/godnd/static/schema.graphql")
	b, _ := ioutil.ReadAll(schemaFile)

	return &Parser{
		GqlSchema: string(b),
	}
}

// func (p *Parser) ParseGqlSchema() {
// 	p.GetRawGoTypeStrings()
// 	fmt.Println(len(p.RawGoTypeStrings))

// 	f, _ := os.Create("/Users/ericshreve/github.com/godnd/pkg/genny/generated_types.go")
// 	defer f.Close()

// 	f.WriteString("package genny\n\n")
// 	for _, t := range p.RawGoTypeStrings {
// 		f.WriteString(t)
// 		f.WriteString("\n")
// 	}
// }
