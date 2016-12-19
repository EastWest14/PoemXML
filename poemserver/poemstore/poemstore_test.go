package poemstore_test

import (
	. "poemXML/poemserver/poemstore"
	"testing"
)

func TestNewPoemStore(t *testing.T) {
	poemStore := NewStore()
	if poemStore == nil {
		t.Error("Initialization of poem store failed. Got a nil store.")
	}
}

func TestGetAllPoems(t *testing.T) {
	//TODO: real test with mocked dependencies

	poemStore := NewStore()
	list, err := poemStore.GetAllPoems()
	if err != nil {
		t.Errorf("Expected nil error, got %s", err.Error())
	}
	if list == nil {
		t.Errorf("Expected non-nil list, got nil")
	}
}
