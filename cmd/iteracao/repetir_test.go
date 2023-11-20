package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	resultado := Repetir("a")
	esperado := "aaaaa"
	if resultado != esperado {
		t.Errorf("esperado %s , resultado %s", esperado, resultado)
	}
}
