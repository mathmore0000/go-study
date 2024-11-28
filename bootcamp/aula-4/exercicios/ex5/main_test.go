package qtdRacao

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getCachorroComida(t *testing.T) {
	t.Run("Teste comida de cachorro positiva", func(t *testing.T) {

		var gotComida float64 = getCachorroComida(4)
		var expected float64 = 40

		require.Equal(t, expected, gotComida)
	})
	t.Run("Teste comida de cachorro negativa", func(t *testing.T) {

		var gotComida float64 = getCachorroComida(-4)
		var expected float64 = 0

		require.Equal(t, expected, gotComida)
	})
}

func Test_getTarantulaComida(t *testing.T) {
	t.Run("Teste comida de tarantula positiva", func(t *testing.T) {

		var gotComida float64 = getTarantulaComida(13)
		var expected float64 = 1.95

		require.Equal(t, expected, gotComida)
	})
	t.Run("Teste comida de tarantula negativa", func(t *testing.T) {

		var gotComida float64 = getTarantulaComida(-13)
		var expected float64 = 0

		require.Equal(t, expected, gotComida)
	})
}

func Test_getGatoComida(t *testing.T) {
	t.Run("Teste comida de gato positiva", func(t *testing.T) {

		var gotComida float64 = getGatoComida(5)
		var expected float64 = 25

		require.Equal(t, expected, gotComida)
	})
	t.Run("Teste comida de gato negativa", func(t *testing.T) {

		var gotComida float64 = getGatoComida(-5)
		var expected float64 = 0

		require.Equal(t, expected, gotComida)
	})
}

func Test_getHamsterComida(t *testing.T) {
	t.Run("Teste comida de hamster positiva", func(t *testing.T) {

		var gotComida float64 = getHamsterComida(44)
		var expected float64 = 11

		require.Equal(t, expected, gotComida)
	})
	t.Run("Teste comida de hamster negativa", func(t *testing.T) {

		var gotComida float64 = getHamsterComida(-44)
		var expected float64 = 0

		require.Equal(t, expected, gotComida)
	})
}
