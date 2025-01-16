package positioner_test

import (
	"testdoubles/positioner"
	"testing"
)

func TestPositionerDefault_GetLinearDistance(t *testing.T) {
	type fields struct {
		positioner *positioner.PositionerDefault
	}
	type args struct {
		from *positioner.Position
		to   *positioner.Position
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{
			name: "coordenadas positivas",
			fields: fields{
				positioner: positioner.NewPositionerDefault(),
			},
			args: args{
				from: &positioner.Position{
					X: 0,
					Y: 0,
					Z: 0,
				},
				to: &positioner.Position{
					X: 1,
					Y: 1,
					Z: 1,
				},
			},
			want: 1.7320508075688772,
		},
		{
			name: "coordenadas negativas",
			fields: fields{
				positioner: positioner.NewPositionerDefault(),
			},
			args: args{
				from: &positioner.Position{
					X: -1,
					Y: -1,
					Z: -1,
				},
				to: &positioner.Position{
					X: -2,
					Y: -2,
					Z: -2,
				},
			},
			want: 1.7320508075688772,
		},
		{
			name: "coordenadas sem decimais",
			fields: fields{
				positioner: positioner.NewPositionerDefault(),
			},
			args: args{
				from: &positioner.Position{
					X: 0,
					Y: 0,
					Z: 0,
				},
				to: &positioner.Position{
					X: 0,
					Y: 0,
					Z: 0,
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields.positioner
			if got := p.GetLinearDistance(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("PositionerDefault.GetLinearDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
