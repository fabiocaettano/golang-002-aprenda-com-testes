package main

import (
	"fmt"
	"testing"
)

func Test_Mensagem_OlaNome(t *testing.T) {
	nome := "fabio"
	resultado := Ola(nome)
	fmt.Println(resultado)
	esperado := "Olá, fabio"
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func Test_Mensagem_OlaMundo(t *testing.T) {
	nome := ""
	resultado := Ola(nome)
	fmt.Println(resultado)
	esperado := "Olá, Mundo"
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func Test_Mensagem_Agrupado(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris")
		esperado := "Olá, Chris"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("'Mundo' como padrão para 'string' vazia", func(t *testing.T) {
		resultado := Ola("")
		esperado := "Olá, Mundo"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
