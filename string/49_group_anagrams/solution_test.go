package group_anagrams

import (
	"reflect"
	"testing"
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
		if reflect.DeepEqual(testCase.Expected, actual) {
			t.Errorf("groupAnagrams execute failed: Strs[%v],  expected[%v], actual[%v]",
				testCase.Strs, testCase.Expected, actual)
		}
	}
}
