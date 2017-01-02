package index

import (
	"encoding/xml"
	"github.com/EastWest14/gAssert"
	"io/ioutil"
)

type Index struct {
	indexPath string
	Elements  []*indexElement
}

func New(indexPath string) *Index {
	return &Index{indexPath: indexPath, Elements: []*indexElement{}}
}

func (ind *Index) LoadIndex() error {
	data, err := ioutil.ReadFile(ind.indexPath)
	if err != nil {
		return err
	}

	err = ind.unmarshalXML(data)
	if err != nil {
		return err
	}

	return nil
}

//AllPoemIds returns the list of all the poem ids currently loaded in the index structure.
//The index is not re-loaded from a file.
func (ind *Index) AllPoemIds() (ids []string) {
	gAssert.Assert(ind != nil, "Index is nil, can't get all poem ids.")
	gAssert.Assert(ind.Elements != nil, "Index elements slice is nil, can't get all poem ids.")

	ids = []string{}
	for _, elem := range ind.Elements {
		ids = append(ids, elem.PoemId)
	}

	return ids
}

func (ind *Index) unmarshalXML(data []byte) error {
	gAssert.Assert(ind.Elements != nil, "Index element slice is nil, can't unmarshal XML index into it.")

	var indexParseStruct _Index
	err := xml.Unmarshal(data, &indexParseStruct)
	if err != nil {
		return err
	}

	ind.removeAllElements()
	for _, xmlElem := range indexParseStruct.Elements {
		elem := xmlElem.convertToElement()
		ind.Elements = append(ind.Elements, elem)
	}

	return nil
}

func (ind *Index) removeAllElements() {
	ind.Elements = []*indexElement{}
}

func GetPoemPathById(id int) (filepath string) {
	return ""
}

//XML Unmarshaling

type _Index struct {
	XMLName      xml.Name          `xml:"index"`
	IndexVersion string            `xml:"index_version,attr"`
	DateCreated  string            `xml:"date_created,attr"`
	DateModified string            `xml:"date_modified,attr"`
	Elements     []_IndexedElement `xml:"indexed_entity"`
}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//

/*const (
	INDEX_VERSION  = "0.0"
	XML_HEADER     = `<?xml version="1.0" encoding="UTF-8"?>`
	ROOT_OPENER    = `<Index:index xmlns:Index="https://github.com/EastWest14/PoemXML/index_schema" index_version="%s" date_created="%s" date_modified="%s">`
	ROOT_CLOSER    = `</Index:index>`
	ELEMENT_ENTITY = `		<Index:indexed_entity>
        <Index:poem_id>%d</Index:poem_id>
        <Index:path>%s</Index:path>
    </Index:indexed_entity>`
)*/

/*func (in *Index) ConvertToXMLString() string {
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
}*/