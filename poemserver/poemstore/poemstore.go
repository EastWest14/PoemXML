package poemstore

import (
	"github.com/EastWest14/errorcode"
	//"poemXML/poemserver/database"
	"database/sql"
	"poemXML/poemserver/poem"
	"poemXML/poemserver/poemlist"
)

const (
	INDEX_UNAVAILABLE_ERROR = "INDEX_UNAVAILABLE"
)

type Store struct {
	storeIndex IndexT
}

func NewStore(storeIndex IndexT, database *sql.DB) *Store {
	return &Store{storeIndex: storeIndex}
}

//Check determines if the store is functionable
func (pStore *Store) Check() error {
	return nil
}

func (pStore *Store) GetAllPoems() (plist *poemlist.PoemList, err *errorcode.Errorcode) {
	return nil, nil
}

func (pStore *Store) GetPoemByID() (p *poem.Record, err *errorcode.Errorcode) {
	return nil, nil
}

type IndexT interface {
	LoadIndex() error
	AllPoemIds() (ids []string)
}
