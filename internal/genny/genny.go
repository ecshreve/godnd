package genny

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/samsarahq/go/oops"
)

func GenerateTypes() {
	apiTypeStr, _ := getAPIGoTypeStr("ability-scores")
	fmt.Println(apiTypeStr)
}

func getAPIGoTypeStr(resourceName string) (string, error) {
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

	typeNameRe := regexp.MustCompile(`<h4>(.*)</h4>`)
	typeNameRaw := typeNameRe.FindStringSubmatch(rawStr)[1]
	typeName := regexp.MustCompile(`\s+`).ReplaceAllString(typeNameRaw, "")

	tableBodyRe := regexp.MustCompile(`(?s)<tbody>.*</tbody>`)
	tableRowRe := regexp.MustCompile(`<tr>\s*(<td.*>\s*.*\s*</td>\s*){3}</tr>`)
	tableBody := tableBodyRe.FindString(rawStr)
	tableRows := tableRowRe.FindAllString(tableBody, -1)

	var fields []string
	for _, row := range tableRows {
		fields = append(fields, getAPIGoTypeFieldStr(row))
	}
	fieldsStr := strings.Join(fields, "\n")

	typeStr := fmt.Sprintf("type %s struct {\n%s\n}", typeName, fieldsStr)
	fmt.Println(typeStr)
	return "", nil
}

func getAPIGoTypeFieldStr(field string) string {
	stripped := regexp.MustCompile(`(\n|\t|  )`).ReplaceAllString(field, "")
	elems := regexp.MustCompile(`</td>`).Split(stripped, 3)

	elemContentsRe := regexp.MustCompile(`<td[^>]*>(.*)$`)
	fieldNameRaw := elemContentsRe.FindStringSubmatch(elems[0])[1]
	fieldName := snakeToCamel(fieldNameRaw)

	fieldDesc := elemContentsRe.FindStringSubmatch(elems[1])[1]
	fieldTypeRaw := elemContentsRe.FindStringSubmatch(elems[2])[1]

	// TODO: add validation checks.
	fieldType := ""
	if regexp.MustCompile(`APIReference`).MatchString(fieldTypeRaw) {
		fieldType = fmt.Sprintf("APIReference")
	} else {
		scalar := regexp.MustCompile(`(\S+)</td>`).FindStringSubmatch(fieldTypeRaw)[1]
		fieldType = fmt.Sprintf("%s", scalar)
	}

	if regexp.MustCompile(`list`).MatchString(fieldTypeRaw) {
		fieldType = fmt.Sprintf("[]%s", fieldType)
	}

	fieldGoTypeStr := fmt.Sprintf("%s %s `json:\"%s\"`// %s", fieldName, fieldType, fieldNameRaw, fieldDesc)
	return fieldGoTypeStr

}
