package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	// DataBase file descriptor
	db *sql.DB
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}
func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	log.Println("DB connection success")
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}
