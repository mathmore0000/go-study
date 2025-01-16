package simulator_test

import (
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testing"
)

func TestCanCatch(t *testing.T) {
	type fields struct {
		positioner *positioner.PositionerDefault
	}
	type args struct {
		preyPosition   *positioner.Position
		preySpeed      float64
		hunterPosition *positioner.Position
		hunterSpeed    float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "prey in range",
			fields: fields{
				positioner: positioner.NewPositionerDefault(),
			},
			args: args{
				hunterSpeed: 5,
				hunterPosition: &positioner.Position{
					X: 0,
					Y: 0,
					Z: 0,
				},
				preySpeed: 1,
				preyPosition: &positioner.Position{
					X: 1,
					Y: 1,
					Z: 1,
				},
			},
			want: true,
		},
		{
			name: "prey out of range",
			fields: fields{
				positioner: positioner.NewPositionerDefault(),
			},
			args: args{
				hunterSpeed: 1,
				hunterPosition: &positioner.Position{
					X: 0,
					Y: 0,
					Z: 0,
				},
				preySpeed: 1,
				preyPosition: &positioner.Position{
					X: 500,
					Y: 500,
					Z: 500,
				},
			},
			want: false,
		},
		{
			name: "prey faster but there is no time",
			fields: fields{
				positioner: positioner.NewPositionerDefault(),
			},
			args: args{
				hunterSpeed: 100,
				hunterPosition: &positioner.Position{
					X: 0,
					Y: 0,
					Z: 0,
				},
				preySpeed: 2,
				preyPosition: &positioner.Position{
					X: 1000,
					Y: 1000,
					Z: 1000,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := simulator.NewCatchSimulatorDefault(1, tt.fields.positioner)
			got := s.CanCatch(&simulator.Subject{
				Position: tt.args.hunterPosition,
				Speed:    tt.args.hunterSpeed,
			}, &simulator.Subject{
				Position: tt.args.preyPosition,
				Speed:    tt.args.preySpeed,
			})
			if got != tt.want {
				t.Errorf("CanCatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
