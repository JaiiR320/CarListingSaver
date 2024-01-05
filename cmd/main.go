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
	router.GET("/dashboard", webHandler.HandleShowDashboard)

	router.GET("/view/output.css", getCss)
	// router.File("/favicon.ico", "/favicon.ico")
	router.GET("/favicon.ico", getFavicon)
	router.Start(":3000")
}

func getFavicon(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "image/x-icon")
	c.Response().WriteHeader(200)
	return c.File("favicon.ico")
}

func getCss(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/css")
	c.Response().WriteHeader(200)
	return c.File("view/output.css")
}
