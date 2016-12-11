package poemstore

import (
	//"poemXML/poemserver/poem"
	"poemXML/poemserver/poemlist"
)

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (pStore *Store) GetAllPoems() (poemlist *poemlist.PoemList) {
	return nil
}
