package handlers_test

import (
	"github.com/EastWest14/errorcode"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	. "poemXML/poemserver/handlers"
	"poemXML/poemserver/poemlist"
	"poemXML/poemserver/poemstore"
	"strings"
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
	listOfPoems            *poemlist.PoemList
	errorOnAllPoemsRequest *errorcode.Errorcode
}

var _ PoemStoreT = &mockPoemStore{}

func (mockS *mockPoemStore) GetAllPoems() (list *poemlist.PoemList, err *errorcode.Errorcode) {
	if mockS.errorOnAllPoemsRequest == nil {
		return mockS.listOfPoems, nil
	}

	return nil, mockS.errorOnAllPoemsRequest
}

func TestPoemListHandler(t *testing.T) {
	h := NewHandlersInstance()

	//Success case
	poemStore := &mockPoemStore{listOfPoems: poemlist.New(), errorOnAllPoemsRequest: nil}
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

	listOfPoems, _ := poemStore.GetAllPoems()
	expectedResponse := listOfPoems.String()

	if string(body) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(body))
	}

	//Index unavailbale error case
	poemStore.errorOnAllPoemsRequest = errorcode.New(poemstore.INDEX_UNAVAILABLE_ERROR, "HELLO!")
	resp, err = http.Get(mockServer.URL)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	bodyString := strings.Trim(string(body), " \n")
	statusCode := resp.StatusCode
	if statusCode != http.StatusInternalServerError || DATA_ACCESS_FAILED_USER_MESSAGE != bodyString {
		t.Errorf("Expected error '%d' and body '%s', got: '%d' and body '%s'", http.StatusInternalServerError, DATA_ACCESS_FAILED_USER_MESSAGE, statusCode, bodyString)
	}

	//Unknown error case
	poemStore.errorOnAllPoemsRequest = errorcode.New("TOTALLY_UNKNOWN_ERROR", "LALALA")
	resp, err = http.Get(mockServer.URL)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}
	bodyString = strings.Trim(string(body), " \n")
	statusCode = resp.StatusCode

	if statusCode != http.StatusInternalServerError || UNKNOWN_ERROR_USER_MESSAGE != bodyString {
		t.Errorf("Expected error '%d' and body '%s', got: '%d' and body '%s'", http.StatusInternalServerError, UNKNOWN_ERROR_USER_MESSAGE, statusCode, bodyString)
	}

}
