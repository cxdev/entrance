package main

import (
	"entrance/backend/execute"
	"entrance/backend/router"
	"entrance/backend/storage"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type EntranceBackend struct {
	storage.StorageService
	execute.ExecuteService
}

func createEntranceBackend(db *gorm.DB, logRoot string) (*EntranceBackend, error) {
	storageservice, err := storage.NewStorageService(db)
	if err != nil {
		return nil, err
	}
	executeservice := execute.NewExecuteService(logRoot)
	return &EntranceBackend{*storageservice, *executeservice}, nil
}

func main() {
	// db, _ := gorm.Open("sqlite3", "file::memory:?mode=memory")
	db, err := gorm.Open("sqlite3", "/tmp/entrance/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	entranceApp, err := createEntranceBackend(db, "/tmp/entrance")
	if err != nil {
		log.Fatal(err)
	}

	router := router.NewRouter(entranceApp)

	log.Fatal(http.ListenAndServe(":9090", router))
}
