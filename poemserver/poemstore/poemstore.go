package poemstore

import (
	"github.com/EastWest14/errorcode"
	"poemXML/poemserver/poemlist"
)

const (
	INDEX_UNAVAILABLE_ERROR = "INDEX_UNAVAILABLE"
)

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (pStore *Store) GetAllPoems() (plist *poemlist.PoemList, err *errorcode.Errorcode) {
	return nil, errorcode.New(INDEX_UNAVAILABLE_ERROR, "Can't find index file")
	//return poemlist.New(), nil
}
