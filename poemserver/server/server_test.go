package server_test

import (
	//"io/ioutil"
	//"net/http"
	//"net/http/httptest"
	//"net/url"
	. "poemXML/poemserver/server"
	"testing"
)

func TestNewServer(t *testing.T) {
	aServer := NewServer()
	if aServer == nil {
		t.Error("Server initialized to nil")
	}
}

/*
func TestHandleDefaultRequest(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(HandleDefaultRequest))
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

func TestFirstPartOfPath(t *testing.T) {
	cases := []struct {
		urlString         string
		expectedFirstPart string
	}{
		{"http://localhost:8080", ""},
		{"http://localhost:8080/", ""},
		{"http://localhost:8080/poems", "poems"},
		{"http://localhost:8080/poems/123", "poems"},
		{"http://domain.com/carrots/r?a=123", "carrots"},
	}

	for i, aCase := range cases {
		urlForm, err := url.Parse(aCase.urlString)
		if err != nil {
			t.Errorf("Error parsing URL in case %d: %s", i, err.Error())
		}
		firstPart := firstPartOfPath(urlForm)

		if firstPart != aCase.expectedFirstPart {
			t.Errorf("Error in case %d. Expected %s, got %s", i, aCase.expectedFirstPart, firstPart)
		}
	}
}*/
