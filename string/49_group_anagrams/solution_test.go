package group_anagrams

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGroupAnagrams(t *testing.T) {
	testCases := []struct {
		Strs     []string
		Expected [][]string
	}{
		{
			Strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			Expected: [][]string{
				{
					"bat",
				},
				{
					"nat", "tan",
				},
				{
					"ate", "eat", "tea",
				},
			},
		},
	}
	for _, testCase := range testCases {
		actual := groupAnagrams(testCase.Strs)
		assert.Equal(t, testCase.Expected, actual,
			fmt.Sprintf("groupAnagrams execute failed: Strs[%v],  expected[%v], actual[%v]",
				testCase.Strs, testCase.Expected, actual))
	}
}
