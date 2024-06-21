package page

import (
	"../document"
	"fmt"
)

const (
	Size = 5
)

// IPage Definindo uma interface com métodos abstratos
type IPage interface {
	AddDocument(document document.Document) error
	RemoveDocument(documentId document.Document) (int, error)
	ListAllDocuments() []*document.Document
}

type Page struct {
	Id        int
	Documents []*document.Document
}

func NewPage(id int) *Page {

	page := Page{Id: id}
	return &page
}

func (p *Page) AddDocument(doc document.Document) error {
	totalSize := 0
	for i := 0; i < len(p.Documents); i++ {
		totalSize += p.Documents[i].Did.Tam
	}

	if totalSize+doc.Did.Tam > Size {
		return fmt.Errorf("a pagina %d ja esta cheia", p.Id)
	}

	doc.Did.Seq = totalSize + 1
	doc.Did.PageId = p.Id
	p.Documents = append(p.Documents, &doc)

	return nil
}

func (p *Page) RemoveDocument(index int) {
	p.Documents = append(p.Documents[:index], p.Documents[index+1:]...) // Remocao, fazendo append de todos os itens anteriores com os posteriores, mas nao ele mesmo
	// Atualizacao da seq
	for j := 0; j < len(p.Documents); j++ {
		p.Documents[j].Did.Seq = j + 1
	}
}

func (p *Page) ListAllDocuments() []*document.Document {
	var DocumentList []*document.Document   // instancio uma lista de ponteiros para documentos
	for k := 0; k < len(p.Documents); k++ { //for que percorre os documentos da pagina
		DocumentList = append(DocumentList, p.Documents[k]) // adiciono na lista instanciada
	}

	return DocumentList
}
