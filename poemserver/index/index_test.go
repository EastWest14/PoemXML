package index_test

import (
	. "poemXML/poemserver/index"
	"testing"
)

func TestNew(t *testing.T) {
	const filepath = "./Fake_filepath"
	anIndex := New(filepath)

	if anIndex == nil {
		t.Errorf("Failed to initialize index - it is nil")
	}

	extractedFilepath := anIndex.GetFilepath()
	if extractedFilepath != filepath {
		t.Errorf("Filepath of a new index set incorrectly. Expected %s, got %s", filepath, extractedFilepath)
	}
}

/*func TestConvertIndexToXMLString(t *testing.T) {
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
}*/
