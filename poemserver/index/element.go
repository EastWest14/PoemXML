package index

import (
	"encoding/xml"
)

type indexElement struct {
	PoemId string
	Path   string
}

type _IndexedElement struct {
	XMLName xml.Name `xml:"indexed_entity"`
	PoemId  string   `xml:"poem_id"`
	Path    string   `xml:"path"`
}

func (ie _IndexedElement) convertToElement() *indexElement {
	elem := &indexElement{PoemId: ie.PoemId, Path: ie.Path}
	return elem
}
