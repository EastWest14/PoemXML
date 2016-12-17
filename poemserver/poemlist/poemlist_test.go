package poemlist_test

import (
	. "poemXML/poemserver/poemlist"
	"strings"
	"testing"
)

func TestNewPoemList(t *testing.T) {
	aPoemList := New()
	if aPoemList == nil {
		t.Error("New PoemList initialized to zero")
	}
}

func TestPoemListString(t *testing.T) {
	//TODO: real test

	aPoemList := New()
	descriptionString := aPoemList.String()

	if !strings.Contains(descriptionString, INTRO_DESCRIPTION) {
		t.Errorf("Expected string to contain substring '%s'. Got: '%s'", INTRO_DESCRIPTION, descriptionString)
	}
}
