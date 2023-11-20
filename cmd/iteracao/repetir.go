package iteracao

const numeroRepeticoes = 5

func Repetir(letra string) string {
	var repeticoes string
	for i := 0; i < numeroRepeticoes; i++ {
		repeticoes = repeticoes + letra
	}
	return repeticoes
}
