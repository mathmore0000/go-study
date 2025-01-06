package repository

import (
	"app/internal"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

func (r *VehicleMap) FindAllByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		if strings.ToLower(value.Color) == strings.ToLower(color) && value.FabricationYear == year {
			v[key] = value
		}
	}

	return
}

func (r *VehicleMap) Create(v *internal.Vehicle) (nv internal.Vehicle, err error) {
	nv = *v
	nv.Id = len(r.db) + 1
	r.db[nv.Id] = nv
	return nv, nil
}

func (r *VehicleMap) ExistsByRegistration(v *internal.Vehicle) bool {
	for _, vehicle := range r.db {
		if vehicle.Registration == v.Registration {
			return true
		}
	}
	return false
}
