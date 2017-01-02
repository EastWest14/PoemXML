package poemlist

import (
	"github.com/EastWest14/gAssert"
)

const (
	INTRO_DESCRIPTION = "List of poems:"
)

type PoemList []string

func New() *PoemList {
	return &PoemList{}
}

func (pl *PoemList) AddPoemIds(poemIds []string) {
	gAssert.Assert(pl != nil, "Attempts to add poem ids to a nil poem list.")

	if len(poemIds) == 0 {
		return
	}

	*pl = append(*pl, poemIds...)
}

func (pl *PoemList) String() string {
	return INTRO_DESCRIPTION + "0 poems"
}
