package platform

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

// Open returns a DB reference for a data source.
func Open(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

var createTableStatements = []string{
	`CREATE TABLE IF NOT EXISTS commands (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			process_type INTEGER,
			command_segments TEXT,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	`CREATE TABLE IF NOT EXISTS jobs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			status INTEGER,
			command_id INTEGER,
			system_cmd TEXT,
			arguments TEXT,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
}

func CreateDB(dataSourceName string) (*DB, error) {
	// var err error
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = createTable(db); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// createTable creates the table, and if necessary, the database.
func createTable(conn *sql.DB) error {
	for _, stmt := range createTableStatements {
		_, err := conn.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
