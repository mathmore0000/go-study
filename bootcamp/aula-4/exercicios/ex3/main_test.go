package salario

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSalarioCatA(t *testing.T) {
	t.Run("Teste horas positivas", func(t *testing.T) {

		gotSalario := getSalarioCatA(5)
		var expected float32 = 5000

		require.Equal(t, expected, gotSalario)
	})
	t.Run("Teste horas negativas", func(t *testing.T) {
		gotSalario := getSalarioCatA(-5)
		var expected float32 = 0

		require.Equal(t, expected, gotSalario)
	})
}

func Test_getSalarioCatB(t *testing.T) {
	t.Run("Teste horas positivas", func(t *testing.T) {

		gotSalario := getSalarioCatB(51)
		var expected float32 = 91800

		require.Equal(t, expected, gotSalario)
	})
	t.Run("Teste horas negativas", func(t *testing.T) {

		gotSalario := getSalarioCatB(-51)
		var expected float32 = 0

		require.Equal(t, expected, gotSalario)
	})
}

func Test_getSalarioCatC(t *testing.T) {
	t.Run("Teste horas positivas", func(t *testing.T) {

		gotSalario := getSalarioCatC(15)
		var expected float32 = 67500

		require.Equal(t, expected, gotSalario)
	})
	t.Run("Teste horas negativas", func(t *testing.T) {

		gotSalario := getSalarioCatC(-15)
		var expected float32 = 0

		require.Equal(t, expected, gotSalario)
	})
}
