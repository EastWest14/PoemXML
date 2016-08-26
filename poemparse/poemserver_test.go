package poemparse

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestEncodePoemToXML(t *testing.T) {
	recMetadata := &RecordMetadata{FormatVersion: "0.0", Id: 1, DateCreated: "", DateModified: ""}
	pMetadata := &PoemMetadata{PoemName: "example", PoemAuthor: "Andrew V. Prosikhin", PoemDateWritten: "August 25, 2016"}
	pBody := &PoemBody{PoemContent: "Hello, PoexXML!"}
	xmlString := encodePoemToXML(recMetadata, pMetadata, pBody)
	fmt.Println(xmlString)
	dummy := &struct{}{}
	if err := xml.Unmarshal([]byte(xmlString), dummy); err != nil {
		t.Errorf("Produced xml is not valid. Error: %s", err.Error())
	}
}
