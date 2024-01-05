package db

import (
	"database/sql"

	"github.com/JaiiR320/carlistingsaver/types"
	_ "github.com/lib/pq"
)

// A storage interface for getting, creating, storing
// and deleting listings in a storage container
type Storage interface {
	CreateListing(listing *types.Listing) error
	GetListings() ([]*types.Listing, error)
	DeleteListing(id int) error
	DropTables() error
}

// A Postgres implementation of the Storage interface
type PostgresStore struct {
	db *sql.DB
}

// Creates a new instance of a postgres db
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

// Initializes the postgres db's tables
// If drop is true, then the tables will be dropped before being created
func (s *PostgresStore) Init(drop bool) error {
	if drop {
		s.DropTables()
	}
	return s.createListingTable()
}

// Creates the listing table
func (s *PostgresStore) createListingTable() error {
	query := `create table if not exists listing (
		id serial primary key,
		url varchar(255),
		title varchar(255),
		price serial,
		mileage serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

// Creates a new listing in the postgres db
// Takes a pointer to a listing to use for the insert
func (s *PostgresStore) CreateListing(l *types.Listing) error {
	query := `insert into listing 
	(url, title, price, mileage, created_at)
	values ($1, $2, $3, $4, $5)`
	_, err := s.db.Query(
		query,
		l.Url,
		l.Title,
		l.Price,
		l.Mileage,
		l.CreatedAt)

	return err
}

// Gets all listings from the postgres db
func (s *PostgresStore) GetListings() ([]*types.Listing, error) {
	rows, err := s.db.Query("select * from listing")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	listings := []*types.Listing{}
	for rows.Next() {
		l, err := scanIntoListing(rows)
		if err != nil {
			return nil, err
		}
		listings = append(listings, l)
	}
	return listings, nil
}

// Deletes a listing from the postgres db
func (s *PostgresStore) DeleteListing(id int) error {
	query := `delete from listing where id = $1`
	_, err := s.db.Exec(query, id)
	return err
}

// Drops all tables from the postgres db
func (s *PostgresStore) DropTables() error {
	query := `drop table if exists listing`
	_, err := s.db.Exec(query)
	return err
}

// Scans a row from the postgres db into a listing
func scanIntoListing(rows *sql.Rows) (*types.Listing, error) {
	l := new(types.Listing)

	err := rows.Scan(
		&l.Id,
		&l.Url,
		&l.Title,
		&l.Price,
		&l.Mileage,
		&l.CreatedAt)

	return l, err
}
