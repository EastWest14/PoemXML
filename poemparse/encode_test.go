package poemparse

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"testing"
)

func TestEncodePoemToXML(t *testing.T) {
	const (
		formatVersion = "0.0"
		id            = 1
		dateCreated   = "27/08/2016"
		dateModified  = "28/08/2016"
	)

	const (
		poemName        = "example"
		poemAuthor      = "Andrew V. Prosikhin"
		poemDateWritten = "August 25, 2016"
	)

	const (
		poemContent = "Hello, PoemXML!"
	)

	recMetadata := &RecordMetadata{FormatVersion: formatVersion, Id: id, DateCreated: dateCreated, DateModified: dateModified}
	pMetadata := &PoemMetadata{PoemName: poemName, PoemAuthor: poemAuthor, PoemDateWritten: poemDateWritten}
	pBody := &PoemBody{PoemContent: poemContent}
	xmlString := encodePoemToXML(recMetadata, pMetadata, pBody)
	fmt.Println(xmlString)
	decPoem := &DecodedPoem{}
	if err := xml.Unmarshal([]byte(xmlString), decPoem); err != nil {
		t.Errorf("Produced xml is not valid. Error: %s", err.Error())
	}
	fmt.Println("Format Version: " + decPoem.FormatVersion)
	fmt.Println("Id: " + decPoem.Id)
	fmt.Println("Date created: " + decPoem.DateCreated)
	fmt.Println("Date modified: " + decPoem.DateModified)
	fmt.Println("Body: " + decPoem.PoemBody.PoemContent)

	if decPoem.FormatVersion != formatVersion {
		t.Errorf("Format version unparsed incorrectly. Expected: %s, got %s", formatVersion, decPoem.FormatVersion)
	}
	if decPoem.Id != strconv.Itoa(id) {
		t.Errorf("Id unparsed incorrectly. Expected: %s, got %s", strconv.Itoa(id), decPoem.Id)
	}
	if decPoem.DateCreated != dateCreated {
		t.Errorf("Date created unparsed incorrectly. Expected: %s, got %s", dateCreated, decPoem.DateCreated)
	}
	if decPoem.DateModified != dateModified {
		t.Errorf("Date modified unparsed incorrectly. Expected: %s, got %s", dateModified, decPoem.DateModified)
	}

	if decPoem.PoemHeader.PoemName != poemName {
		t.Errorf("Poem header field unparsed incorrectly. Expected: %s, got %s", poemName, decPoem.PoemHeader.PoemName)
	}
	if decPoem.PoemHeader.PoemAuthor != poemAuthor {
		t.Errorf("Poem header field unparsed incorrectly. Expected: %s, got %s", poemAuthor, decPoem.PoemHeader.PoemAuthor)
	}
	if decPoem.PoemHeader.PoemDateWritten != poemDateWritten {
		t.Errorf("Poem header field unparsed incorrectly. Expected: %s, got %s", poemDateWritten, decPoem.PoemHeader.PoemDateWritten)
	}

	if decPoem.PoemBody.PoemContent != poemContent {
		t.Errorf("Poem content unparsed incorrectly. Expected: %s, got %s", poemContent, decPoem.PoemBody.PoemContent)
	}
}
