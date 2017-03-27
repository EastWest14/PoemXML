package poemstore_test

import (
	. "poemXML/poemserver/poemstore"
	"testing"
)

func TestNewPoemStore(t *testing.T) {
	poemStore := NewStore(nil, nil)
	if poemStore == nil {
		t.Error("Initialization of poem store failed. Got a nil store.")
	}
}

func TestCheck(t *testing.T) {

}

func TestGetAllPoems(t *testing.T) {

}
