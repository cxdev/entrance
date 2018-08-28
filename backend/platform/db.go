package platform

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

type Query struct {
	Query string
	Args  []interface{}
}

type QueryPlan struct {
	db      *DB
	queries []*Query
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
			name TEXT,
			process_type INTEGER,
			command_segments TEXT,
			created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
	`CREATE TABLE IF NOT EXISTS jobs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			status INTEGER,
			command_id INTEGER,
			arguments TEXT,
			system_cmd TEXT,
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

func (db *DB) NewQueryPlan() *QueryPlan {
	return &QueryPlan{db, []*Query{}}
}

func (db *DB) Exec(query string, args ...interface{}) (res sql.Result, err error) {
	return db.NewQueryPlan().AppendQuery(query, args).Exec()
}

func (qp *QueryPlan) AppendQuery(query string, args ...interface{}) *QueryPlan {
	qp.queries = append(qp.queries, &Query{query, args})
	return qp
}

func (qp *QueryPlan) Exec() (res sql.Result, err error) {
	tx, _ := qp.db.Begin()
	defer tx.Rollback()

	for _, query := range qp.queries {
		if res, err = tx.Exec(query.Query, query.Args); err != nil {
			return
		}
	}

	err = tx.Commit()
	return
}
