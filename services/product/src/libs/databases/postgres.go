package databases

import (
	_ "github.com/lib/pq" // here
	"github.com/jmoiron/sqlx"
	"log"
)

// Postgres connection struct
type Postgres struct {
	Connection *sqlx.DB
	ConnStr    string
	Tx         *sqlx.Tx
}

// NewPostgres return struct
func NewPostgres(connStr string) *Postgres {
	if len(connStr) == 0 {
		log.Println("[ALERT] Cannot connect to database. No database URL was given")
	}

	return &Postgres{
		ConnStr: connStr,
	}
}

// Connect creates connection with database
func (db *Postgres) Connect() (*sqlx.DB, error) {
	var err error
	if db.Connection == nil {
		db.Connection, err = db.getConnection()
	}
	return db.Connection, err
}

// GetConnection return connection only
func (db *Postgres) GetConnection() *sqlx.DB {
	conn, err := db.Connect()
	if err != nil {
		log.Println("Failed connecting to the database", err.Error())
	}
	return conn
}

func (db *Postgres) getConnection() (*sqlx.DB, error) {
	conn, err := sqlx.Connect("postgres", db.ConnStr)
	if err != nil {
		log.Println("Failed connecting to the database:", db.ConnStr, err.Error())
		return conn, err
	}

	return conn, err
}
