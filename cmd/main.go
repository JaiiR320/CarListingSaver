package main

import (
	"log"

	"github.com/JaiiR320/carlistingsaver/db"
	"github.com/JaiiR320/carlistingsaver/handler/api"
	"github.com/JaiiR320/carlistingsaver/handler/web"

	"github.com/labstack/echo/v4"
)

func main() {
	app := NewServer(":3000")
	app.Run()
}

// A server that handles the web and api routes
type Server struct {
	listenAddr string
	db         db.Storage
}

// Creates a new server with the given listen address and a postgres db
func NewServer(listenAddr string) *Server {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(false); err != nil {
		log.Fatal(err)
	}

	return &Server{
		listenAddr: listenAddr,
		db:         store,
	}
}

// Run starts the server with the given routes and initializations
func (s *Server) Run() {
	router := echo.New()

	// initialize handlers
	webHandler := web.WebHandler{}
	apiHandler := api.ApiHandler{Store: s.db}

	// API routes
	// public routes
	router.POST("/api/listing", apiHandler.HandleMakeListing)
	router.GET("/api/listing", apiHandler.HandleGetListings)
	router.DELETE("/api/listing/:id", apiHandler.HandleDeleteListing)

	// admin routes
	router.DELETE("/admin/drop", apiHandler.HandleDropTables)

	// HTML routes
	router.GET("/", webHandler.HandleShowHome)

	router.Start(":3000")
}
