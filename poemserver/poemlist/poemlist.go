package poemlist

import ()

const (
	INTRO_DESCRIPTION = "List of poems:"
)

type PoemList struct {
}

func NewPoemList() *PoemList {
	return &PoemList{}
}

func (pl *PoemList) String() string {
	return INTRO_DESCRIPTION + "0 poems"
}
