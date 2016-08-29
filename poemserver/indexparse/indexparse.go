package indexparse

import (
	"encoding/xml"
	"fmt"
	"time"
)

const (
	INDEX_VERSION  = "0.0"
	XML_HEADER     = `<?xml version="1.0" encoding="UTF-8"?>`
	ROOT_OPENER    = `<Index:index xmlns:Index="https://github.com/EastWest14/PoemXML/index_schema" index_version="%s" date_created="%s" date_modified="%s">`
	ROOT_CLOSER    = `</Index:index>`
	ELEMENT_ENTITY = `		<Index:indexed_entity>
        <Index:poem_id>%d</Index:poem_id>
        <Index:path>%s</Index:path>
    </Index:indexed_entity>`
)

type Index struct {
	IndexVersion string
	DateCreated  time.Time
	DateModified time.Time
	Elements     []*IndexedElement
}

func NewIndex(dateCreated, dateModified time.Time, elements []*IndexedElement) *Index {
	return &Index{IndexVersion: INDEX_VERSION, DateCreated: dateCreated, DateModified: dateModified, Elements: elements}
}

func (in *Index) ConvertToXMLString() string {
	rootOpenerFilled := fmt.Sprintf(ROOT_OPENER, in.IndexVersion, in.DateCreated.String(), in.DateModified.String())
	var elementsRaw []byte
	if len(in.Elements) <= 0 {

	} else {
		for _, element := range in.Elements {
			elementsRaw = append(elementsRaw, []byte(fmt.Sprintf(ELEMENT_ENTITY, element.PoemId, element.PoemPath))...)
		}
	}
	return XML_HEADER + "\n" + rootOpenerFilled + "\n" + string(elementsRaw) + "\n" + ROOT_CLOSER
}

type IndexedElement struct {
	PoemId   int
	PoemPath string
}

func NewIndexedElement(poemId int, poemPath string) *IndexedElement {
	return &IndexedElement{PoemId: poemId, PoemPath: poemPath}
}

type IndexParsed struct {
	XMLName      xml.Name               `xml:"index"`
	IndexVersion string                 `xml:"index_version,attr"`
	DateCreated  string                 `xml:"date_created,attr"`
	DateModified string                 `xml:"date_modified,attr"`
	Elements     []IndexedElementParsed `xml:"indexed_entity"`
}

type IndexedElementParsed struct {
	XMLName xml.Name `xml:"indexed_entity"`
	PoemId  string   `xml:"poem_id"`
	Path    string   `xml:"path"`
}
