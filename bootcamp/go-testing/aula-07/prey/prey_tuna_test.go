package prey_test

/*
	Exercício 2
	Realizar testes unitários da implementação tuna de Prey,
		nos métodos GetSpeed e GetPosition.
			No mínimo, cubra um caso em que a velocidade e a posição tenham seu valor padrão e outro caso em que tenham um valor diferente do default.
*/

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"
)

func TestGetPositionWithDefaultValues(t *testing.T) {
	tuna := prey.PreyStub{
		GetSpeedFunc: func() (speed float64) {
			return 0
		},
		GetPositionFunc: func() (position *positioner.Position) {
			return &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}
		},
	}
	if tuna.GetPosition().X != 0 || tuna.GetPosition().Y != 0 || tuna.GetPosition().Z != 0 {
		t.Errorf("Tuna position should be (0, 0, 0)")
	}
}

func TestGetSpeedWithDefaultValues(t *testing.T) {
	tuna := prey.PreyStub{
		GetSpeedFunc: func() (speed float64) {
			return 0
		},
		GetPositionFunc: func() (position *positioner.Position) {
			return &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}
		},
	}
	if tuna.GetSpeed() != 0 {
		t.Errorf("Tuna speed should be 0")
	}
}

func TestGetSpeedWithCustomValues(t *testing.T) {
	tuna := prey.PreyStub{
		GetSpeedFunc: func() (speed float64) {
			return 100
		},
		GetPositionFunc: func() (position *positioner.Position) {
			return &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}
		},
	}
	if tuna.GetSpeed() != 100 {
		t.Errorf("Tuna speed should be 100")
	}
}

func TestGetPositionWithCustomValues(t *testing.T) {
	tuna := prey.PreyStub{
		GetSpeedFunc: func() (speed float64) {
			return 0
		},
		GetPositionFunc: func() (position *positioner.Position) {
			return &positioner.Position{
				X: 100,
				Y: 100,
				Z: 100,
			}
		},
	}
	if tuna.GetPosition().X != 100 || tuna.GetPosition().Y != 100 || tuna.GetPosition().Z != 100 {
		t.Errorf("Tuna position should be (100, 100, 100)")
	}
}
