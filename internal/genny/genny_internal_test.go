package genny

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFieldType(t *testing.T) {
	testcases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "simple string scalar",
			input:    "string",
			expected: "string",
		},
		{
			desc:     "simple integer scalar",
			input:    "integer",
			expected: "int32",
		},
		{
			desc:     "list of scalar",
			input:    "list string",
			expected: "[]string",
		},
		{
			desc:     "simple api reference",
			input:    ` <a href="#apireference">APIReference</a> (<a href="#skills">Skills</a>)`,
			expected: "APIReference",
		},
		{
			desc:     "list api reference",
			input:    ` list <a href="#apireference">APIReference</a> (<a href="#skills">Skills</a>)`,
			expected: "[]APIReference",
		},
		{
			desc:     "list api reference with whitespace",
			input:    `\n\t\t list <a href="#apireference">APIReference</a> (<a href="#skills">Skills</a>)`,
			expected: "[]APIReference",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := parseFieldType(testcase.input)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
