package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	. "poemXML/poemserver/handlers"
	"testing"
)

func TestNewHandlersInstance(t *testing.T) {
	handlers := NewHandlersInstance()
	if handlers == nil {
		t.Errorf("Error creating handlers - nil returned")
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

func TestPoemListHandler(t *testing.T) {
	h := NewHandlersInstance()
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

	expectedResponse := "List of poems"

	if string(body) != expectedResponse {
		t.Errorf("Expected response %s, got %s", expectedResponse, string(body))
	}
}
