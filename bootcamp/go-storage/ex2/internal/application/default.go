package application

import (
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-sql-driver/mysql"
)

// ConfigDefault is a struct that represents the default application configuration
type ConfigDefault struct {
	// Database is the database configuration
	Database mysql.Config
	// Address is the address of the application
	Address string
}

// NewDefault returns a new default application
func NewDefault(cfg *ConfigDefault) *Default {
	// default
	cfgDefault := &ConfigDefault{
		Address: ":8080",
	}
	if cfg != nil {
		cfgDefault.Database = cfg.Database
		if cfg.Address != "" {
			cfgDefault.Address = cfg.Address
		}
	}

	return &Default{
		cfgDb: cfgDefault.Database,
		addr:  cfgDefault.Address,
	}
}

// Default is a struct that represents the default application
type Default struct {
	// cfgDb is the database configuration
	cfgDb mysql.Config
	// addr is the address of the application
	addr string
}

// Run runs the default application
func (d *Default) Run() (err error) {
	// dependencies
	// - database: connection
	db, err := sql.Open("mysql", d.cfgDb.FormatDSN())
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return
	}

	// - repository: products
	rp := repository.NewProductsMySQL(db)
	rw := repository.NewWarehousesMySQL(db)

	// - handler: products
	hp := handler.NewProductsDefault(rp)
	hw := handler.NewWarehousesDefault(rw)

	// - router: chi
	rt := chi.NewRouter()
	// - router: middlewares
	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)
	// - router: routes
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", hp.GetAll())
		// - GET /products
		r.Get("/{id}", hp.GetOne())
		// - POST /products
		r.Post("/", hp.Create())
		// - PUT /products/{id}
		r.Patch("/{id}", hp.Update())
		// - DELETE /products/{id}
		r.Delete("/{id}", hp.Delete())
	})

	rt.Route("/warehouses", func(r chi.Router) {
		r.Get("/", hw.GetAll())
		r.Get("/{id}", hw.GetOne())
		r.Get("/reportProducts", hw.GetProductsCount())
		r.Post("/", hw.Create())
	})

	// run
	err = http.ListenAndServe(d.addr, rt)
	if err != nil {
		return
	}
	return
}
