package iteracao

func Repetir(letra string) string {
	var repeticoes string
	for i := 0; i < 5; i++ {
		repeticoes = repeticoes + letra
	}
	return repeticoes
}
