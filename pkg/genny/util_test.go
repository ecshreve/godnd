package genny_test

import (
	"testing"

	"github.com/ecshreve/godnd/pkg/genny"
	"github.com/stretchr/testify/assert"
)

func TestIsApiReference(t *testing.T) {
	testcases := []struct {
		desc     string
		input    string
		expected bool
	}{
		{
			desc: "is",
			input: `type TraitSubrace struct {
				Index string
				Name string
				Url string
			}`,
			expected: true,
		},
		{
			desc: "no",
			input: `type AbilityScore struct {
				Desc []string
				FullName string
				Index string
				Name string
				Skills []AbilityScoreSkill
				Url string
			}`,
			expected: false,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			actual := genny.IsCommonType(testcase.input)
			assert.Equal(t, testcase.expected, actual)
		})
	}
}
