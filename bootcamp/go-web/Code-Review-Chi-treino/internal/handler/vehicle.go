package handler

import (
	"app/internal"
	"app/pkg"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// VehicleRequestBody is a struct that represents a body for the creation of a vehicle in JSON format
type VehicleRequestBody struct {
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var vehicleRequest VehicleRequestBody
		if err := json.NewDecoder(r.Body).Decode(&vehicleRequest); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "Dados do veículo mal formatados ou incompletos.",
			})
			return
		}

		newVehicle := internal.Vehicle{
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           vehicleRequest.Brand,
				Model:           vehicleRequest.Model,
				Registration:    vehicleRequest.Registration,
				Color:           vehicleRequest.Color,
				FabricationYear: vehicleRequest.FabricationYear,
				Capacity:        vehicleRequest.Capacity,
				MaxSpeed:        vehicleRequest.MaxSpeed,
				FuelType:        vehicleRequest.FuelType,
				Transmission:    vehicleRequest.Transmission,
				Weight:          vehicleRequest.Weight,
				Dimensions: internal.Dimensions{
					Height: vehicleRequest.Height,
					Length: vehicleRequest.Length,
					Width:  vehicleRequest.Width,
				},
			},
		}
		nv, err := h.sv.Create(&newVehicle)
		if err != nil {
			if errors.Is(pkg.ErrRegistrationAlreadyExists, err) {
				response.JSON(w, http.StatusConflict, map[string]any{
					"message": err.Error(),
				})
				return
			}
			response.JSON(w, http.StatusConflict, map[string]any{
				"message": err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    nv,
		})
	}
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")
		year, err := strconv.Atoi(chi.URLParam(r, "year"))

		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{
				"message": "year invalid",
			})
			return
		}

		v, err := h.sv.FindAllByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		if len(data) == 0 {
			response.JSON(w, http.StatusNotFound, map[string]any{
				"message": "not found",
				"data":    data,
			})
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}