package genny_test

import (
	"testing"

	"github.com/ecshreve/godnd/internal/genny"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateType(t *testing.T) {
	c, err := genny.NewConverter()
	require.NoError(t, err)

	testcases := []struct {
		desc        string
		input       string
		expected    string
		expectError bool
	}{
		{
			desc:        "ability score",
			input:       genny.TestGqlTypeString["AbilityScore"],
			expected:    genny.TestGoTypeString["AbilityScore"],
			expectError: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual, err := c.ConvertGqlToGoType(testcase.input)
			assert.Equal(t, testcase.expectError, err != nil)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
