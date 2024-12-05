package media

/*
	Exercício 2 - Calcular a média

	Uma escola precisa calcular a média (por aluno) de suas notas.
	É solicitado que ela gere uma função na qual possam ser passados N números inteiros e que retorne a média.
	Não é possível inserir notas negativas.
*/

func getMedia(notas ...float32) (media float32) {
	media = 0
	var nota float32
	var lenNotas int = len(notas)
	for _, nota = range notas {
		if nota < 0 {
			lenNotas--
			continue
		}
		media += nota
	}
	if lenNotas == 0 {
		return 0
	}
	return media / float32(lenNotas)
}
