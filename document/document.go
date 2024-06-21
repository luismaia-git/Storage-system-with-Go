package document

import (
	"fmt"
)

type DID struct {
	PageId int
	Seq    int
	Tam    int
}

type Document struct {
	Did     DID
	Content string
}

func NewDocument(char string) (*Document, error) {
	size := len(char)
	if size < 1 || size > 5 {
		return nil, fmt.Errorf("[Erro] Documento '%s' deve ter entre 1 e 5 bytes", char)
	}

	did := DID{Tam: len(char)}
	doc := Document{Did: did, Content: char}

	return &doc, nil
}
