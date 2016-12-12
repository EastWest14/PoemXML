package server_test

import (
	"poemXML/poemserver/handlers"
	. "poemXML/poemserver/server"
	"testing"
)

func TestNewServer(t *testing.T) {
	aServer := NewServer()
	if aServer == nil {
		t.Error("Server initialized to nil")
	}
}

func TestSetHandlers(t *testing.T) {
	aServer := NewServer()
	h := handlers.NewHandlersInstance()
	aServer.SetHandlers(h)
	if aServer.GetHandlers() == nil {
		t.Errorf("Error in setting handlers, handlers are nil")
	}
}

/*func firstPartOfPath(u *url.URL) string {
	urlPath := u.Path
	pathParts := strings.Split(urlPath, "/")

	if len(pathParts) > 1 {
		return pathParts[1]
	}

	return ""
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
