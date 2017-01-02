package poemstore

import (
	"github.com/EastWest14/errorcode"
	"github.com/EastWest14/gAssert"
	"poemXML/poemserver/index"
	"poemXML/poemserver/poem"
	"poemXML/poemserver/poemlist"
)

const (
	INDEX_UNAVAILABLE_ERROR = "INDEX_UNAVAILABLE"
)

type Store struct {
	storeIndex IndexT
}

func NewStore(storeIndex IndexT) *Store {
	return &Store{storeIndex: storeIndex}
}

//Check determines if the store is functionable
func (pStore *Store) Check() error {
	storeIndex := pStore.storeIndex
	gAssert.Assert(storeIndex != nil, "Index is nil")

	err := storeIndex.LoadIndex()
	return err
}

func (pStore *Store) GetAllPoems() (plist *poemlist.PoemList, err *errorcode.Errorcode) {
	storeIndex := pStore.storeIndex
	indLoadErr := storeIndex.LoadIndex()
	if indLoadErr != nil {
		return nil, errorcode.New(INDEX_UNAVAILABLE_ERROR, indLoadErr.Error())
	}

	listOfIds := storeIndex.AllPoemIds()
	if len(listOfIds) == 0 {
		return nil, nil
	}

	pList := poemlist.New()
	pList.AddPoemIds(listOfIds)

	return pList, nil
}

func (pStore *Store) GetPoemByID() (p *poem.Record, err *errorcode.Errorcode) {
	return nil, nil
}

type IndexT interface {
	LoadIndex() error
	AllPoemIds() (ids []string)
}

var _ IndexT = index.New("")
