package main

import "fmt"

const prefixoCumprimento = "Olá, "

func Ola(nome string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixoCumprimento + nome
}

func main() {
	nome := ""
	fmt.Println(Ola(nome))
}
