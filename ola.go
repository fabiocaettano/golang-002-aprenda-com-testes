package main

import "fmt"

func Ola(nome string) string {
	return "Olá, " + nome
}

func main() {
	nome := "fabio"
	fmt.Println(Ola(nome))
}
