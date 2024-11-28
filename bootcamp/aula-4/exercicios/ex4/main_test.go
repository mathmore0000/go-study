package minMaxAvg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minFunc(t *testing.T) {
	t.Run("Teste min", func(t *testing.T) {

		gotSalario := minFunc(-3, 4, 3, 9, 10, 24, -23)
		var expected float32 = -23

		require.Equal(t, expected, gotSalario)
	})
}

func Test_averageFunc(t *testing.T) {
	t.Run("Teste avg", func(t *testing.T) {

		gotSalario := averageFunc(1, 2, 3, 4, 5)
		var expected float32 = 3

		require.Equal(t, expected, gotSalario)
	})
}

func Test_maxFunc(t *testing.T) {
	t.Run("Teste max", func(t *testing.T) {

		gotMax := maxFunc(2, 3, 4, 6, 8, 6, 4)
		var expected float32 = 8

		require.Equal(t, expected, gotMax)
	})
}
