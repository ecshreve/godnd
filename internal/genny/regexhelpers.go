package genny

import (
	"context"
	"fmt"
	"regexp"

	"github.com/samsarahq/go/oops"
)

func getAllEndpoints(ctx context.Context, docData string) ([]string, error) {
	fullEndpointRe := regexp.MustCompile(`(?s)<h3.*?>.*?</h3>.*?<table>(\s|.)*?</table>`)

	allRawEndpoints := fullEndpointRe.FindAllString(docData, -1)
	if len(allRawEndpoints) <= 0 {
		return nil, oops.Errorf("no endpoints found in file")
	}

	return allRawEndpoints, nil
}

func getEndpointURL(ctx context.Context, endpointStr string) (string, error) {
	endpointURLRe := regexp.MustCompile(`<h3.*?>.*? (.*?)</h3>`)

	endpointURLMatches := endpointURLRe.FindStringSubmatch(endpointStr)
	if len(endpointURLMatches) < 2 {
		return "", oops.Errorf("unable to get endpointURL from endpointStr: %s...[truncated]", endpointStr[:100])
	}

	// TODO: add additional validation on the endpoint url.
	return endpointURLMatches[1], nil
}

func getFieldTable(ctx context.Context, endpointStr string) (string, error) {
	fieldTableRe := regexp.MustCompile(`(?s)<h4>(.*?)</h4>\s*<table>(\s|.)*?</table>`)

	fieldTable := fieldTableRe.FindString(endpointStr)
	if fieldTable == "" {
		return "", oops.Errorf("unable to get fieldTable from endpointStr: %s...[truncated]", endpointStr[:100])
	}

	return fieldTable, nil
}

func getTypeNameFromFieldTable(ctx context.Context, fieldTableStr string) (string, error) {
	typeNameRe := regexp.MustCompile(`<h4>(.*?)</h4>\s*<table>`)

	typeNameMatches := typeNameRe.FindStringSubmatch(fieldTableStr)
	if len(typeNameMatches) < 2 {
		return "", oops.Errorf("unable to get typeName from fieldTable: %s...[truncated]", fieldTableStr[:100])
	}

	typeName := spaceSepToCamel(typeNameMatches[1])
	return typeName, nil
}

func getTableRows(ctx context.Context, fieldTableStr string) ([]string, error) {
	tableRowRe := regexp.MustCompile(`(?s)<tr>\s*(<td.*?>.*?</td>\s*){3}\s*</tr>`)

	tableRows := tableRowRe.FindAllString(fieldTableStr, -1)
	if len(tableRows) < 1 {
		return nil, oops.Errorf("unable to get tableRows from fieldTable: %s...[truncated]", fieldTableStr[:100])
	}

	return tableRows, nil
}

func getTableRowElements(ctx context.Context, tableRow string) ([]string, error) {
	rowElementRe := regexp.MustCompile(`(?s)<td.*?>.*?</td>`)

	rowElements := rowElementRe.FindAllString(tableRow, -1)
	if len(rowElements) != 3 {
		return nil, oops.Errorf("unable to get rowElements from tableRow: %s", tableRow)
	}

	return rowElements, nil
}

type tableRowElementContents struct {
	fieldNameRaw string
	fieldTypeRaw string
}

func getTableRowElementContents(ctx context.Context, rowElements []string) (*tableRowElementContents, error) {
	if len(rowElements) != 3 {
		return nil, oops.Errorf("unable to parse rowElements slice with length not equal to 3: %v", rowElements)
	}

	elementContentsRe := regexp.MustCompile(`(?s)<td.*?>(.*?)</td>`)

	fieldNameRawMatches := elementContentsRe.FindStringSubmatch(rowElements[0])
	if len(fieldNameRawMatches) < 2 {
		return nil, oops.Errorf("unable to get element contents from rowElement: %v", rowElements[0])
	}

	fieldTypeRawMatches := elementContentsRe.FindStringSubmatch(rowElements[2])
	if len(fieldTypeRawMatches) < 2 {
		return nil, oops.Errorf("unable to get element contents from rowElement: %v", rowElements[2])
	}

	return &tableRowElementContents{
		fieldNameRaw: fieldNameRawMatches[1],
		fieldTypeRaw: fieldTypeRawMatches[1],
	}, nil
}

func parseFieldType(ctx context.Context, fieldTypeRaw string) (string, error) {
	var parsedFieldType string
	apiCommonTypeNames := []string{"APIReference", "Choice", "Cost", "AbilityBonus"}

	// Check for the types defined in the Common Models section of the API docs.
	for _, commonTypeName := range apiCommonTypeNames {
		if regexp.MustCompile(commonTypeName).MatchString(fieldTypeRaw) {
			parsedFieldType = commonTypeName
			break
		}
	}

	if parsedFieldType == "" {
		if regexp.MustCompile(`(?s).*string.+`).MatchString(fieldTypeRaw) {
			parsedFieldType = "URLRefString"
		} else if regexp.MustCompile("object").MatchString(fieldTypeRaw) {
			parsedFieldType = "map[string]interface{}"
		} else if scalar := regexp.MustCompile(`[[:alpha:]]+$`).FindString(fieldTypeRaw); scalar != "" {
			parsedFieldType = apiScalarToGoScalar[scalar]
		} else {
			parsedFieldType = "interface{}"
		}
	}

	// Check if the element contains "list" and if it does pre-pend the parsedFieldType
	// with brackets to indicate a slice.
	if regexp.MustCompile("list").MatchString(fieldTypeRaw) {
		parsedFieldType = fmt.Sprintf("[]%s", parsedFieldType)
	}

	return parsedFieldType, nil
}

func buildGoTypeField(ctx context.Context, fieldName, fieldNameRaw, fieldType string) string {
	return fmt.Sprintf("%s %s `json:\"%s\"`", fieldName, fieldType, fieldNameRaw)
}
