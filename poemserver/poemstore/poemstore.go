package poemstore

import (
	"github.com/EastWest14/errorcode"
	"poemXML/poemserver/index"
	"poemXML/poemserver/poem"
	"poemXML/poemserver/poemlist"
)

const (
	INDEX_UNAVAILABLE_ERROR = "INDEX_UNAVAILABLE"
)

type Store struct {
	storeIndex *index.Index
}

func NewStore(storeIndex *index.Index) *Store {
	return &Store{storeIndex: storeIndex}
}

func (pStore *Store) GetAllPoems() (plist *poemlist.PoemList, err *errorcode.Errorcode) {
	//return nil, errorcode.New(INDEX_UNAVAILABLE_ERROR, "Can't find index file")
	return poemlist.New(), nil
}

func (pStore *Store) GetPoemByID() (p *poem.Record, err *errorcode.Errorcode) {
	return nil, nil
}
