package genny_test

import (
	"testing"

	"github.com/ecshreve/godnd/pkg/genny"
	"github.com/stretchr/testify/assert"
)

func TestGenerateType(t *testing.T) {
	testcases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "AbilityScore",
			input:    genny.TestGqlTypeString["AbilityScore"],
			expected: genny.TestGoTypeString["AbilityScore"],
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual, _ := genny.GenerateType(testcase.input)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
