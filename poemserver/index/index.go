package index

import (
	"encoding/xml"
	"io/ioutil"
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
	indexPath string
	//IndexVersion string
	//DateCreated  time.Time
	//DateModified time.Time
	//Elements     []*IndexedElement
}

func New(indexPath string) *Index {
	return &Index{indexPath: indexPath}
}

func (ind *Index) LoadIndex() error {
	data, err := ioutil.ReadFile(ind.indexPath)
	if err != nil {
		return err
	}

	err = ind.UnmarshalXML(data)
	if err != nil {
		return err
	}

	return nil
}

func (ind *Index) UnmarshalXML(data []byte) error {
	var indexParseStruct _Index
	err := xml.Unmarshal(data, &indexParseStruct)

	if err != nil {
		return err
	}

	//fmt.Println(indexParseStruct)
	return nil
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

type _IndexedElement struct {
	XMLName xml.Name `xml:"indexed_entity"`
	PoemId  string   `xml:"poem_id"`
	Path    string   `xml:"path"`
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
