package poemlist

import (
	"fmt"
	"github.com/EastWest14/gAssert"
	"poemXML/utils"
)

const (
	DESCRIPTION_TEMPLATE = `List of poems: %d poems.`
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
	gAssert.Assert(pl != nil, "Attempts to convert a nil poem list to string.")

	numberOfIds := len(*pl)
	return fmt.Sprintf(DESCRIPTION_TEMPLATE, numberOfIds) + "\n"
}

func (pl1 *PoemList) Equal(pl2 *PoemList) bool {
	if pl1 == nil || pl2 == nil {
		if pl1 == nil && pl2 == nil {
			return true
		}
		return false
	}

	return utils.SliceOfStringEqual([]string(*pl1), []string(*pl2))
}
