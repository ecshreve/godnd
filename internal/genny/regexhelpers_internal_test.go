package genny

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllEndpoints(t *testing.T) {
	ctx := context.Background()

	endpoints, err := getAllEndpoints(ctx, validFullDocStr)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(endpoints))
}

func TestGetEndpointURL(t *testing.T) {
	ctx := context.Background()

	endpointURL, err := getEndpointURL(ctx, validEndpointStr)
	assert.NoError(t, err)
	assert.Equal(t, "api/foo/{index}", endpointURL)
}

func TestGetFieldTable(t *testing.T) {
	ctx := context.Background()

	fieldTable, err := getFieldTable(ctx, validEndpointStr)
	assert.NoError(t, err)
	assert.Equal(t, validFieldTableStr, fieldTable)
}

func TestGetTypeNameFromFieldTable(t *testing.T) {
	ctx := context.Background()

	fieldTypeName, err := getTypeNameFromFieldTable(ctx, validFieldTableStr)
	assert.NoError(t, err)
	assert.Equal(t, "Foo", fieldTypeName)
}

func TestGetTableRows(t *testing.T) {
	ctx := context.Background()

	rows, err := getTableRows(ctx, validFieldTableStr)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(rows))
}

func TestGetTableRowElements(t *testing.T) {
	ctx := context.Background()

	elements, err := getTableRowElements(ctx, validTableRow)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(elements))
}

func TestGetTableRowElementContents(t *testing.T) {
	ctx := context.Background()

	rowElements := []string{`<td align="left">index</td>`, `<td align="left">index field description</td>`, `<td align="left">string</td>`}
	expected := &tableRowElementContents{
		fieldNameRaw: "index",
		fieldTypeRaw: "string",
	}

	rowElementContents, err := getTableRowElementContents(ctx, rowElements)
	assert.NoError(t, err)
	assert.Equal(t, expected, rowElementContents)
}

func TestParseFieldType(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		desc         string
		fieldTypeRaw string
		expected     string
		expectError  bool
	}{
		{
			desc:         "simple scalar",
			fieldTypeRaw: "string",
			expected:     "string",
		},
		{
			desc:         "simple scalar list",
			fieldTypeRaw: "list string",
			expected:     "[]string",
		},
		{
			desc:         "simple object",
			fieldTypeRaw: "object",
			expected:     "map[string]interface{}",
		},
		{
			desc:         "simple APIReference",
			fieldTypeRaw: `<a href="#apireference">APIReference</a> (<a href="#skills">Skills</a>)`,
			expected:     "APIReference",
		},
		{
			desc:         "list of APIReference",
			fieldTypeRaw: `list <a href="#apireference">APIReference</a> (<a href="#skills">Skills</a>)`,
			expected:     "[]APIReference",
		},
		{
			desc:         "special url refrence string for class object",
			fieldTypeRaw: `\t\t\t\tstring (<a href="#starting-equipment">StartingEquipment</a>)\t\t\t`,
			expected:     "URLRefString",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual, err := parseFieldType(ctx, testcase.fieldTypeRaw)
			assert.Equal(t, testcase.expected, actual)

			if testcase.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestBuildGoFieldType(t *testing.T) {
	ctx := context.Background()

	testcases := []struct {
		desc         string
		fieldName    string
		fieldType    string
		fieldNameRaw string
		expected     string
	}{
		{
			desc:         "simple scalar",
			fieldName:    "Index",
			fieldType:    "int32",
			fieldNameRaw: "index",
			expected:     "Index int32 `json:\"index\"`",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := buildGoTypeField(ctx, testcase.fieldName, testcase.fieldNameRaw, testcase.fieldType)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
