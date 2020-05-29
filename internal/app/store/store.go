package store

import (
	"database/sql"

	_ "github.com/lib/pq" // ...
)

// Store ...
type Store struct {
	config         *Config
	db             *sql.DB
	userPepository *UserPepository
}

// New ...
func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Open ...
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close
func (s *Store) Close() {
	// ....

	s.db.Close()

}

func (s *Store) User() *UserPepository {
	if s.userPepository != nil {
		return s.userPepository
	}

	s.userPepository = &UserPepository{
		store: s,
	}
	return s.userPepository
}
