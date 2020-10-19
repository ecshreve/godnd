package genny

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/samsarahq/go/oops"
)

func GenerateTypes() {
	apiTypeStr, _ := parseResourceFile("ability-scores")

	fmt.Println(apiTypeStr)
}

func parseResourceFile(resourceName string) (string, error) {
	docFilePath := fmt.Sprintf("%s/doc-resource-%s.ejs", DocFilePath, resourceName)
	docFile, err := os.Open(docFilePath)
	if err != nil {
		return "", oops.Wrapf(err, "error opening doc file for resource: %s with path: %s", resourceName, docFilePath)
	}

	b, err := ioutil.ReadAll(docFile)
	if err != nil {
		return "", oops.Wrapf(err, "error reading doc file for resource: %s", resourceName)
	}
	rawStr := string(b)

	x := parseEndpoints(rawStr)
	fmt.Println(x)
	return "", nil
}

func parseEndpoints(fullDoc string) map[string]string {
	fullEndpointRe := regexp.MustCompile(`(?s)<h3>.*?</h3>.*?<table>(\s|.)*?</table>`)
	endpointRefRe := regexp.MustCompile(`<h3>.*? (.*?)</h3>`)
	endpoints := fullEndpointRe.FindAllString(fullDoc, -1)

	endpointRefToTypeStr := make(map[string]string)
	for _, endpoint := range endpoints {
		endpointRef := endpointRefRe.FindStringSubmatch(endpoint)[1]
		endpointTypeStr := parseEndpoint(endpoint)
		endpointRefToTypeStr[endpointRef] = endpointTypeStr
	}
	return endpointRefToTypeStr
}

func parseEndpoint(endpointStr string) string {
	typeTableRe := regexp.MustCompile(`(?s)<h4>(.*?)</h4>\s*<table>(\s|.)*?</table>`)
	typeTable := typeTableRe.FindString(endpointStr)
	return parseTable(typeTable)
}

func parseTable(table string) string {
	tableTitleRe := regexp.MustCompile(`<h4>(.*?)</h4>`)
	tableTitle := tableTitleRe.FindStringSubmatch(table)[1]
	typeName := spaceSepToCamel(tableTitle)

	tableRowRe := regexp.MustCompile(`(?s)<tr>\s*(<td.*?>.*?</td>\s*){3}\s*</tr>`)
	tableRows := tableRowRe.FindAllString(table, -1)

	var parsedRows []string
	for _, row := range tableRows {
		parsedRows = append(parsedRows, parseTableRow(row))
	}
	fields := strings.Join(parsedRows, "\n")

	typeStr := fmt.Sprintf("type %s struct {\n%s\n}", typeName, fields)
	formattedTypeStr, _ := format.Source([]byte(typeStr))
	return string(formattedTypeStr)
}

func parseTableRow(row string) string {
	rowElementRe := regexp.MustCompile(`(?s)<td.*?>.*?</td>`)
	rowElements := rowElementRe.FindAllString(row, -1)

	elemContentsRe := regexp.MustCompile(`(?s)<td.*?>(.*?)</td>`)
	fieldNameRaw := elemContentsRe.FindStringSubmatch(rowElements[0])[1]
	fieldName := snakeToCamel(fieldNameRaw)

	fieldTypeRaw := elemContentsRe.FindStringSubmatch(rowElements[2])[1]
	fieldType := parseFieldType(fieldTypeRaw)

	parsed := fmt.Sprintf("%s %s `json:\"%s\"`", fieldName, fieldType, fieldNameRaw)

	return parsed
}

func parseFieldType(elem string) string {
	fieldType := ""
	if regexp.MustCompile(`APIReference`).MatchString(elem) {
		fieldType = fmt.Sprintf("APIReference")
	} else if regexp.MustCompile(`Choice`).MatchString(elem) {
		fieldType = fmt.Sprintf("Choice")
	} else {
		scalar := regexp.MustCompile(`\S+$`).FindString(elem)
		fieldType = fmt.Sprintf("%s", apiScalarToGoScalar[scalar])
	}

	if regexp.MustCompile(`list`).MatchString(elem) {
		fieldType = fmt.Sprintf("[]%s", fieldType)
	}

	return fieldType
}
