package poemparse

import (
	"fmt"
)

const (
	XML_HEADER  = `<?xml version="1.0" encoding="UTF-8"?>`
	ROOT_OPENER = `,,,,,,,<Poem:poem xmlns:Poem="https://github.com/EastWest14/PoemXML/poem_schema" format_version="%s" id="%d" date_created="%s" date_modified="%s">`
	ROOT_CLOSER = `</Poem:poem>`

	POEM_HEADER = `	<Poem:header>
 		<Poem:poem_name> %s </Poem:poem_name>
 		<Poem:poem_author> %s </Poem:poem_author>
 		<Poem:poem_date_written> %s </Poem:poem_date_written>
	</Poem:header>`

	POEM_BODY = `	<Poem:body>
 		<Poem:poem_content>
			%s
 		</Poem:poem_content>
 	</Poem:body>`
)

type RecordMetadata struct {
	FormatVersion string
	Id            int
	DateCreated   string
	DateModified  string
}

type PoemMetadata struct {
	PoemName        string
	PoemAuthor      string
	PoemDateWritten string
}

type PoemBody struct {
	PoemContent string
}

func encodePoemToXML(recMetadata *RecordMetadata, pMetadata *PoemMetadata, pBody *PoemBody) (xmlString string) {
	rootOpener := fmt.Sprintf(ROOT_OPENER, recMetadata.FormatVersion, recMetadata.Id, recMetadata.DateCreated, recMetadata.DateModified)
	header := fmt.Sprintf(POEM_HEADER, pMetadata.PoemName, pMetadata.PoemAuthor, pMetadata.PoemDateWritten)
	body := fmt.Sprintf(POEM_BODY, pBody.PoemContent)
	return XML_HEADER + "\n" + rootOpener + "\n" + header + "\n" + body + "\n" + ROOT_CLOSER
}
