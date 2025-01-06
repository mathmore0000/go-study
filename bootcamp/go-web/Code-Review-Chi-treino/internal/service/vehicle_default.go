package service

import (
	"app/internal"
	"app/pkg"
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

func (s *VehicleDefault) FindAllByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAllByColorAndYear(color, year)
	return
}

func (s *VehicleDefault) Create(v *internal.Vehicle) (nv internal.Vehicle, err error) {
	if s.rp.ExistsByRegistration(v) {
		return *v, pkg.ErrRegistrationAlreadyExists
	}
	nv, err = s.rp.Create(v)

	return
}
