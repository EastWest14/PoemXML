package poemlist

import (
	"testing"
)

func (pl *PoemList) PoemList() (poemIds []string) {
	if pl == nil {
		panic("Nil poem list.")
	}
	return []string(*pl)
}

func TestPoemListEqual(t *testing.T) {
	threeElementList := PoemList([]string{"ID_1", "ID_2", "ID_3"})

	cases := []struct {
		pl1         *PoemList
		pl2         *PoemList
		expectEqual bool
	}{
		{nil, nil, true},
		{nil, New(), false},
		{New(), nil, false},
		{New(), New(), true},
		{&threeElementList, &threeElementList, true},
		{&threeElementList, nil, false},
		{New(), &threeElementList, false},
	}

	for i, aCase := range cases {
		equal := aCase.pl1.Equal(aCase.pl2)
		if equal && !aCase.expectEqual {
			t.Errorf("Error in case %d. Expected not equal, got equal", i)
		}
		if !equal && aCase.expectEqual {
			t.Errorf("Error in case %d. Expected equal, got not equal", i)
		}
	}
}
