package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	. "poemXML/poemserver/handlers"
	"poemXML/poemserver/poemlist"
	"poemXML/poemserver/poemstore"
	"testing"
)

func TestNewHandlersInstance(t *testing.T) {
	handlers := NewHandlersInstance()
	if handlers == nil {
		t.Errorf("Error creating handlers - nil returned")
	}
}

func TestSetPoemStore(t *testing.T) {
	handlers := NewHandlersInstance()

	handlers.SetPoemStore(nil)
	if handlers.GetPoemStore() != nil {
		t.Errorf("Expecting poem store to be nil, got %v", handlers.GetPoemStore())
	}

	poemStore := poemstore.NewStore()
	handlers.SetPoemStore(poemStore)
	if handlers.GetPoemStore() == nil {
		t.Errorf("Expecting poem store to be non-nil, got nil")
	}
}

func TestHandleDefaultRequest(t *testing.T) {
	h := NewHandlersInstance()
	mockServer := httptest.NewServer(http.HandlerFunc(h.DefaultHandler))
	defer mockServer.Close()

	resp, err := http.Get(mockServer.URL)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	if string(body) != STANDARD_RESPONSE {
		t.Errorf("Expected response %s, got %s", STANDARD_RESPONSE, string(body))
	}
}

type mockPoemStore struct {
	listOfPoems *poemlist.PoemList
}

var _ PoemStoreT = &mockPoemStore{}

func (mockS *mockPoemStore) GetAllPoems() *poemlist.PoemList {
	return mockS.listOfPoems
}

func TestPoemListHandler(t *testing.T) {
	h := NewHandlersInstance()

	poemStore := &mockPoemStore{listOfPoems: poemlist.New()}
	h.SetPoemStore(poemStore)

	mockServer := httptest.NewServer(http.HandlerFunc(h.PoemListHandler))
	defer mockServer.Close()

	resp, err := http.Get(mockServer.URL)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	expectedResponse := poemStore.GetAllPoems().String()

	if string(body) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(body))
	}
}
