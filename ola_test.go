package main

import "testing"

func TestOla(t *testing.T) {
	nome := "fabio"
	resultado := Ola(nome)
	esperado := "Olá, fabio"
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
