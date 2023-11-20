package arraysslices

import "testing"

func Test(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	esperado := 15
	resultado := Soma(array)
	for i := 0; i < 5; i++ {
		resultado = resultado + array[i]
	}
	if resultado != esperado {
		t.Errorf("esperado %d , resultado %d", esperado, resultado)
	}
}
