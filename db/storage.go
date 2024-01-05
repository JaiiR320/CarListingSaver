package db

import (
	"database/sql"
	"fmt"

	"github.com/JaiiR320/carlistingsaver/scraper"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateListing(listing *scraper.Listing) error
	GetListings() ([]*scraper.Listing, error)
	DropTables() error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=JairMeza1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createListingTable()
}

func (s *PostgresStore) createListingTable() error {
	query := `create table if not exists listing (
		id serial primary key,
		url varchar(255),
		title varchar(255),
		price serial,
		mileage serial
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateListing(l *scraper.Listing) error {
	query := `insert into listing 
	(url, title, price, mileage)
	values ($1, $2, $3, $4)`
	resp, err := s.db.Query(
		query,
		l.Url,
		l.Title,
		l.Price,
		l.Mileage)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) GetListings() ([]*scraper.Listing, error) {
	rows, err := s.db.Query("select * from listing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	listings := []*scraper.Listing{}
	for rows.Next() {
		l, err := scanIntoListing(rows)
		if err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	return listings, nil
}

func scanIntoListing(rows *sql.Rows) (*scraper.Listing, error) {
	l := new(scraper.Listing)
	id := 0
	err := rows.Scan(
		&id,
		&l.Url,
		&l.Title,
		&l.Price,
		&l.Mileage)

	return l, err
}

func (s *PostgresStore) DropTables() error {
	query := `drop table if exists listing`
	_, err := s.db.Exec(query)
	return err
}
