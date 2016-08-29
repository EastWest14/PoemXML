package indexparse

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestNewIndex(t *testing.T) {
	dCreated := time.Date(2016, 8, 28, 12, 0, 0, 0, time.UTC)
	dModified := time.Date(2016, 8, 28, 12, 30, 0, 0, time.UTC)
	elements := []*IndexedElement{}
	const (
		poemId   = 1
		poemPath = "example/path"
	)
	elements = append(elements, NewIndexedElement(poemId, poemPath))
	ind := NewIndex(dCreated, dModified, elements)
	if ind.IndexVersion != INDEX_VERSION {
		t.Errorf("Index initialized incorrectly. Expected index version: %s, got %s.", INDEX_VERSION, ind.IndexVersion)
	}
	if ind.DateCreated != dCreated {
		t.Errorf("Index initialized incorrectly. Expected date created: %s, got %s.", dCreated.String(), ind.DateCreated.String())
	}
	if ind.DateModified != dModified {
		t.Errorf("Index initialized incorrectly. Expected date modified: %s, got %s.", dModified.String(), ind.DateModified.String())
	}
	if len(ind.Elements) != 1 {
		t.Fatal("Index initialized incorrectly. Expected %d element(s), got %d element(s).", 1, len(ind.Elements))
	}
	if ind.Elements[0].PoemId != poemId {
		t.Errorf("Indexed element initialized incorrectly. Expected poem id: %d, got %d.", poemId, ind.Elements[0].PoemId)
	}
	if ind.Elements[0].PoemPath != poemPath {
		t.Errorf("Indexed element initialized incorrectly. Expected poem path: %s, got %s.", poemPath, ind.Elements[0].PoemPath)
	}
}

func TestConvertIndexToXMLString(t *testing.T) {
	dCreated := time.Date(2016, 8, 28, 12, 0, 0, 0, time.UTC)
	dModified := time.Date(2016, 8, 28, 12, 30, 0, 0, time.UTC)
	const (
		ID_ONE   = 1
		ID_TWO   = 2
		PATH_ONE = "example/path"
		PATH_TWO = "example/path2"
	)
	elements := []*IndexedElement{}
	elements = append(elements, NewIndexedElement(ID_ONE, PATH_ONE))
	elements = append(elements, NewIndexedElement(ID_TWO, PATH_TWO))
	ind := NewIndex(dCreated, dModified, elements)
	indexXMLString := ind.ConvertToXMLString()
	fmt.Println(indexXMLString)

	indParsed := IndexParsed{}
	if err := xml.Unmarshal([]byte(indexXMLString), &indParsed); err != nil {
		t.Fatal(err.Error())
	}
	if indParsed.IndexVersion != INDEX_VERSION {
		t.Errorf("Index encoded incorrectly. Expected index version: %s, got %s.", INDEX_VERSION, indParsed.IndexVersion)
	}
	if indParsed.DateCreated != dCreated.String() {
		t.Errorf("Index encoded incorrectly. Expected date created: %s, got %s.", dCreated.String(), indParsed.DateCreated)
	}
	if indParsed.DateModified != dModified.String() {
		t.Errorf("Index encoded incorrectly. Expected date moified: %s, got %s.", dModified.String(), indParsed.DateModified)
	}

	//validate elements
	if len(indParsed.Elements) != 2 {
		t.Fatalf("Index encoded incorrectly. Expected 2 elements, got %d elements.", len(indParsed.Elements))
	}
	if indParsed.Elements[0].PoemId != strconv.Itoa(ID_ONE) {
		t.Errorf("Index encoded incorrectly. Expected poem_id[0]=%s, got %s", strconv.Itoa(ID_ONE), indParsed.Elements[0].PoemId)
	}
	if indParsed.Elements[1].PoemId != strconv.Itoa(ID_TWO) {
		t.Errorf("Index encoded incorrectly. Expected poem_id[1]=%s, got %s", strconv.Itoa(ID_TWO), indParsed.Elements[1].PoemId)
	}
	if indParsed.Elements[0].Path != PATH_ONE {
		t.Errorf("Index encoded incorrectly. Expected path[0]=%s, got %s", PATH_ONE, indParsed.Elements[0].Path)
	}
	if indParsed.Elements[1].Path != PATH_TWO {
		t.Errorf("Index encoded incorrectly. Expected path[1]=%s, got %s", PATH_TWO, indParsed.Elements[1].Path)
	}
}
