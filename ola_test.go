package main

import (
	"testing"
)

func TestOla1(t *testing.T) {
	idioma := "espanhol"
	resultado := Ola(idioma)
	esperado := "Hola Mundo"
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestOla2(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas em espanhol", func(t *testing.T) {
		resultado := Ola("espanhol")
		esperado := "Hola Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas em francês", func(t *testing.T) {
		resultado := Ola("frances")
		esperado := "Bonjour Le Monde"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas em ingles", func(t *testing.T) {
		resultado := Ola("ingles")
		esperado := "Hello World"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
