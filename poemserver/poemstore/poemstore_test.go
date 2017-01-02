package poemstore_test

import (
	"errors"
	"poemXML/poemserver/poemlist"
	. "poemXML/poemserver/poemstore"
	"testing"
)

func TestNewPoemStore(t *testing.T) {
	anIndex := &mockIndex{}

	poemStore := NewStore(anIndex)
	if poemStore == nil {
		t.Error("Initialization of poem store failed. Got a nil store.")
	}
}

func TestCheck(t *testing.T) {
	//No error case
	anIndex := &mockIndex{errorOnLoad: nil}
	poemStore := NewStore(anIndex)

	err := poemStore.Check()
	if err != nil {
		t.Errorf("Expected no error when checking index loading, got: %s", err.Error())
	}

	//Error case
	anIndex = &mockIndex{errorOnLoad: errors.New("Can't find the index")}
	poemStore = NewStore(anIndex)

	err = poemStore.Check()
	if err == nil {
		t.Errorf("Expected error when checking index loading, got no error")
	}
}

//TODO: convert to table-driven test
func TestGetAllPoems(t *testing.T) {
	//Loading error case
	indexLoadError := &mockIndex{errorOnLoad: errors.New("Failed to load index"), listOfIds: nil}
	poemStore := NewStore(indexLoadError)
	list, err := poemStore.GetAllPoems()
	if err == nil {
		t.Errorf("Expected non-nil error, got no error")
	} else if err.Type() != INDEX_UNAVAILABLE_ERROR {
		t.Errorf("Expected index unavailable error, got: %s", err.Error())
	}
	if list != nil {
		t.Errorf("Expected nil list, got: %v", list)
	}

	//Nil elements case
	indexNilElements := &mockIndex{errorOnLoad: nil, listOfIds: nil}
	poemStore = NewStore(indexNilElements)
	list, err = poemStore.GetAllPoems()
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if list != nil {
		t.Errorf("Expected nil list, got: %v", list)
	}

	//No elements case
	indexNoElements := &mockIndex{errorOnLoad: nil, listOfIds: []string{}}
	poemStore = NewStore(indexNoElements)
	list, err = poemStore.GetAllPoems()
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	if list != nil {
		t.Errorf("Expected nil list, got: %v", list)
	}

	//Three elements case
	indexThreeElements := &mockIndex{errorOnLoad: nil, listOfIds: []string{"ID_1", "ID_2", "ID_3"}}
	poemStore = NewStore(indexThreeElements)
	list, err = poemStore.GetAllPoems()
	if err != nil {
		t.Errorf("Expected no error, got: %s", err.Error())
	}
	poemListForComparrison := poemlist.New()
	poemListForComparrison.AddPoemIds([]string{"ID_1", "ID_2", "ID_3"})
	if !list.Equal(poemListForComparrison) {
		t.Errorf("Poem list filled up incorrectly")
	}
}

type mockIndex struct {
	errorOnLoad error
	listOfIds   []string
}

var _ IndexT = &mockIndex{}

func (mIndex *mockIndex) LoadIndex() error {
	return mIndex.errorOnLoad
}

func (mIndex *mockIndex) AllPoemIds() (ids []string) {
	return mIndex.listOfIds
}
