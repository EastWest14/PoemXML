package poemstore

import (
	"poemXML/poemserver/poem"
)

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (pStore *Store) GetAllPoems() (poems []poem.Record) {
	return nil
}
