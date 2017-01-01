package utils_test

import (
	. "poemXML/utils"
	"testing"
)

func TestSliceOfStringEqual(t *testing.T) {
	cases := []struct {
		slice1      []string
		slice2      []string
		expectEqual bool
	}{
		{[]string{"abc", "def"}, []string{"abc", "def"}, true},
		{[]string{"abc", "def"}, nil, false},
		{nil, []string{"abc", "def"}, false},
		{[]string{}, nil, false},
		{nil, nil, true},
		{[]string{}, []string{}, true},
		{[]string{"abc", "def"}, []string{"abc"}, false},
		{[]string{"abc", "def"}, []string{"def", "abc"}, false},
		{[]string{" "}, []string{" "}, true},
	}

	for i, aCase := range cases {
		equal := SliceOfStringEqual(aCase.slice1, aCase.slice2)

		if equal && !aCase.expectEqual {
			t.Errorf("Error in case %d. Expected not equal, got equal", i)
		}
		if !equal && aCase.expectEqual {
			t.Errorf("Error in case %d. Expected equal, got not equal", i)
		}
	}
}
