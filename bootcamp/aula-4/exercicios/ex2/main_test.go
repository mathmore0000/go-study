package media

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
	Exercício 2 - Calcular a média
	A escola relatou que as operações para calcular a média não estão sendo realizadas corretamente, portanto, agora somos obrigados a realizar os testes correspondentes:

	Calcular a média das notas dos alunos.
*/

func Test_getMedia(t *testing.T) {
	t.Run("Testando média sem negativos", func(t *testing.T) {
		gotMedia := getMedia(1, 2, 3, 4, 5)
		var expected float32 = 3
		require.Equal(t, expected, gotMedia)
	})
	t.Run("Testando média com alguns negativos", func(t *testing.T) {
		gotMedia := getMedia(1, 2, 3, 4, 5, -5, -3)
		var expected float32 = 3
		require.Equal(t, expected, gotMedia)
	})
	t.Run("Testando média com somente negativos", func(t *testing.T) {
		gotMedia := getMedia(-1, -2, -3, -4-5, -6)
		var expected float32 = 0
		require.Equal(t, expected, gotMedia)
	})
}
