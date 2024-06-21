package main

import (
	"./document"
	"./system"
	"fmt"
)

// Funcoes auxiliares para print
func printScan(all []*document.Document) {
	for _, doc := range all {
		fmt.Println(*doc)
	}
}

func printSeek(did *document.DID, err error) {
	if err == nil {
		fmt.Println(*did)
	} else {
		fmt.Print("Erro: ", err)
	}
}

func main() {

	s := system.System{}

	// teste insercao de tamanhos diferentes
	fmt.Println("===========================================================")
	fmt.Println("Teste insercao de tamanhos diferentes: ")
	err := s.Insert("malu")
	if err != nil {
		return
	}
	err = s.Insert("1")
	if err != nil {
		return
	}
	printScan(s.Scan())
	err = s.Insert("malu")
	if err != nil {
		return
	}
	printScan(s.Scan())
	printSeek(s.Seek("malu"))
	s.Delete("malu")
	printScan(s.Scan())

	/* s.Insert("ghij")
	s.Insert("klmno")
	s.Insert("abcdefghijk")
	fmt.Println("Teste Scan(): ")
	printScan(s.Scan())

	fmt.Println("===========================================================")
	fmt.Println("Teste do delete: ")
	s.Delete("abc")  // documento nao existente
	s.Delete("ghij") // documento existente
	fmt.Println("Teste Scan(): ")
	printScan(s.Scan())

	fmt.Println("===========================================================")
	fmt.Println("Insert: ")
	s.Insert("zvb")

	s.Insert("sb")
	fmt.Println("Teste Scan(): ")
	printScan(s.Scan())

	s.Delete("zvb")
	fmt.Println("Teste Scan(): ")
	printScan(s.Scan())

	//teste inserir um doc do mesmo tamanho do que foi removido
	fmt.Println("===========================================================")
	fmt.Println("Teste inserir um doc do mesmo tamanho do que foi removido: ")
	s.Insert("ghi")
	printScan((s.Scan()))

	//Teste Seek
	printSeek(s.Seek("de"))  // documento que nao existe
	printSeek(s.Seek("ghi")) // documento que existe

	// Teste pagina vazia
	fmt.Println("Esvaziando uma pagina e inserindo um novo documento")
	s.Insert("abcde")
	s.Delete("klmno")
	printScan(s.Scan())
	s.Insert("a")
	printScan(s.Scan())*/
}
