package poemstore_test

import (
	"errors"
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

func TestGetAllPoems(t *testing.T) {
	anIndex := &mockIndex{}

	poemStore := NewStore(anIndex)
	list, err := poemStore.GetAllPoems()
	if err != nil {
		t.Errorf("Expected nil error, got %s", err.Error())
	}
	if list == nil {
		t.Errorf("Expected non-nil list, got nil")
	}
}

type mockIndex struct {
	errorOnLoad error
}

var _ IndexT = &mockIndex{}

func (mIndex *mockIndex) LoadIndex() error {
	return mIndex.errorOnLoad
}
