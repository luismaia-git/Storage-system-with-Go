package system

import (
	"../document"
	"../page"
	"fmt"
)

const (
	LimitPage = 20
)

type System struct {
	Pages []*page.Page
}

func (s *System) Scan() []*document.Document {

	var documentListScan []*document.Document
	for i := 0; i < len(s.Pages); i++ {
		allDocs := s.Pages[i].ListAllDocuments()
		documentListScan = append(documentListScan, allDocs...)
	}
	fmt.Println("===========================================================")
	return documentListScan
}

func (s *System) Seek(char string) (*document.DID, error) {
	docs := s.Scan()
	for j := 0; j < len(docs); j++ {
		doc := docs[j]
		if char == doc.Content {
			return &doc.Did, nil
		}
	}
	return nil, fmt.Errorf("documento '%s' nao encontrado\n", char)
}

func (s *System) Delete(char string) {

	doc, err := s.Seek(char)
	if doc == nil {
		fmt.Println(err)
	} else {
		s.Pages[doc.PageId-1].RemoveDocument(doc.Seq - 1) //-1 porque a posição no array é id-1, o mesmo vale para seq
		fmt.Printf("documento '%s' da pagina %d removido.\n", char, doc.PageId)
	}

}

func (s *System) Insert(char string) error {
	fmt.Println("===========================================================")
	newDoc, errorDocument := document.NewDocument(char) // Cria o novo documento

	if newDoc == nil {
		fmt.Println(errorDocument)
		return nil
	}

	if len(s.Pages) == 0 {
		s.Pages = append(s.Pages, page.NewPage(1))
	}

	for l := 0; l < len(s.Pages); l++ {

		currentPage := s.Pages[l]

		if err := currentPage.AddDocument(*newDoc); err == nil {
			fmt.Printf("Documento '%s' adicionado com sucesso na pagina %d.\n", char, currentPage.Id)
			// fmt.Println(newDoc)
			return nil
		} else {

			fmt.Printf("Erro ao adicionar documento '%s': %s\n", char, err)

		}
	}

	if len(s.Pages)+1 > LimitPage {
		fmt.Printf("Limite de paginas excedido\n")
		return nil
	} else {
		s.Pages = append(s.Pages, page.NewPage(s.Pages[len(s.Pages)-1].Id+1))
		lastPage := s.Pages[len(s.Pages)-1]
		err := lastPage.AddDocument(*newDoc)
		if err != nil {
			return err
		}
		fmt.Printf("Documento '%s' adicionado com sucesso na pagina %d.\n", char, lastPage.Id)
		return nil
	}

}
