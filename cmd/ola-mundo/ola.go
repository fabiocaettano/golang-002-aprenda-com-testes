package main

import "fmt"

func Ola(idioma string) string {

	prefixo := "Olá "
	mundo := "Mundo"

	switch idioma {
	case "frances":
		mundo = "Le Monde"
	case "espanhol":
		mundo = "Mundo"
	case "ingles":
		mundo = "World"
	}

	switch idioma {
	case "frances":
		prefixo = "Bonjour "
	case "espanhol":
		prefixo = "Hola "
	case "ingles":
		prefixo = "Hello "
	}

	return prefixo + mundo
}

func main() {
	idioma := "espanhol"
	fmt.Println(Ola(idioma))
}
