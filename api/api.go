package api

import (
	"encoding/json"
	"log"

	"github.com/JaiiR320/carlistingsaver/db"
	"github.com/JaiiR320/carlistingsaver/scraper"
	"github.com/JaiiR320/carlistingsaver/types"
	"github.com/labstack/echo/v4"
)

type Server struct {
	listenAddr string
	db         db.Storage
}

func NewServer(listenAddr string) *Server {
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	return &Server{
		listenAddr: listenAddr,
		db:         store,
	}
}

func (s *Server) Run() {
	router := echo.New()
	router.POST("/api/listing", s.HandleMakeListing)
	router.GET("/api/listing", s.HandleGetListings)
	router.DELETE("/api/admin/delete", s.HandleDropTables)
	router.Start(":3000")
}

func (s *Server) HandleMakeListing(c echo.Context) error {
	log.Println("POST for making a listing was called")
	var req types.MakeListingRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return err
	}
	listing := scraper.GetListing(req.Url)
	log.Printf("Listing: %+v\n created from URL\n", listing)
	if err := s.db.CreateListing(&listing); err != nil {
		log.Fatal(err)
	}
	log.Println("Added listing to database")

	return writeJSON(c, 200, map[string]scraper.Listing{"Created": listing})
}

func (s *Server) HandleGetListings(c echo.Context) error {
	log.Println("GET for getting listings was called")
	listings, err := s.db.GetListings()
	if err != nil {
		log.Fatal(err)
	}
	return writeJSON(c, 200, listings)
}

func (s *Server) HandleDropTables(c echo.Context) error {
	log.Println("DELETE for dropping tables was called")
	if err := s.db.DropTables(); err != nil {
		log.Fatal(err)
	}
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	s.db = store

	return writeJSON(c, 200, map[string]string{"Dropped": "Tables"})
}

func writeJSON(c echo.Context, status int, v any) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(status)
	return json.NewEncoder(c.Response()).Encode(v)
}
