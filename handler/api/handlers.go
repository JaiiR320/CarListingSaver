package api

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/JaiiR320/carlistingsaver/db"
	"github.com/JaiiR320/carlistingsaver/handler"
	"github.com/JaiiR320/carlistingsaver/scraper"
	"github.com/JaiiR320/carlistingsaver/types"
	"github.com/labstack/echo/v4"
)

type ApiHandler struct {
	Store db.Storage
}

func (h *ApiHandler) HandleMakeListing(c echo.Context) error {
	log.Println("POST for making a listing was called")
	var req types.MakeListingRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return err
	}
	listing := scraper.GetListing(req.Url)
	log.Printf("Listing: %+v\n created from URL\n", listing)
	if err := h.Store.CreateListing(&listing); err != nil {
		log.Fatal(err)
	}
	log.Println("Added listing to database")

	return handler.WriteJSON(c, 200, map[string]types.Listing{"Created": listing})
}

func (h *ApiHandler) HandleGetListings(c echo.Context) error {
	log.Println("GET for getting listings was called")
	listings, err := h.Store.GetListings()
	if err != nil {
		log.Fatal(err)
	}
	return handler.WriteJSON(c, 200, listings)
}

func (h *ApiHandler) HandleDeleteListing(c echo.Context) error {
	log.Println("DELETE for deleting a listing was called")
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	if err := h.Store.DeleteListing(idInt); err != nil {
		log.Fatal(err)
	}
	return handler.WriteJSON(c, 200, map[string]string{"Deleted": id})
}

func (h *ApiHandler) HandleDropTables(c echo.Context) error {
	log.Println("DELETE for dropping tables was called")
	store, err := db.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(true); err != nil {
		log.Fatal(err)
	}
	h.Store = store

	return handler.WriteJSON(c, 200, map[string]string{"Dropped": "Tables"})
}
