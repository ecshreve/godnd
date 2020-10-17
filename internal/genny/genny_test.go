package genny_test

import (
	"fmt"
	"testing"

	"github.com/ecshreve/godnd/internal/genny"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateType(t *testing.T) {
	c, err := genny.NewConverter()
	require.NoError(t, err)

	for _, supportedType := range genny.SupportedTypes {
		gennyTestHelper(supportedType, c, t)
	}
}

func gennyTestHelper(typeName string, c *genny.Converter, t *testing.T) {
	t.Run(fmt.Sprintf("generate go type from gql schema for: %s", typeName), func(t *testing.T) {
		input := genny.TestGqlTypeString[typeName]
		expected := genny.TestGoTypeString[typeName]

		actual, err := c.ConvertGqlToGoType(input)
		assert.NoError(t, err)
		assert.Equal(t, expected, actual)
	})
}
