package main

import (
	entrance "entrance/backend"
	"entrance/backend/execute"
	"entrance/backend/router"
	"entrance/backend/storage"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/yaml.v2"
)

type EntranceBackend struct {
	storage.StorageService
	execute.ExecuteService
}

type config struct {
	Version     string             `yaml:"version"`
	LogRootPath string             `yaml:"log_root_path"`
	SqlitePath  string             `yaml:"sqlite_path"`
	Commands    []entrance.Command `yaml:"commands"`
}

func initConfig(configPath string) (*config, error) {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func initEntranceBackend(db *gorm.DB, logRoot string) (*EntranceBackend, error) {
	storageservice, err := storage.NewStorageService(db)
	if err != nil {
		return nil, err
	}
	executeservice := execute.NewExecuteService(logRoot)
	return &EntranceBackend{*storageservice, *executeservice}, nil
}

func initCommands(commands *[]entrance.Command, entranceBackend *EntranceBackend) error {
	if commands != nil {
		for _, command := range *commands {
			_, err := entranceBackend.CreateCommand(command.Name, command.CommandType, command.CommandSegments)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "conf", "", "set configuration file")
	flag.Parse()
	println(configPath)

	config, err := initConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(config)

	// db, _ := gorm.Open("sqlite3", "file::memory:?mode=memory")
	db, err := gorm.Open("sqlite3", config.SqlitePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	entranceApp, err := initEntranceBackend(db, config.LogRootPath)
	if err != nil {
		log.Fatal(err)
	}

	err = initCommands(&config.Commands, entranceApp)
	if err != nil {
		log.Fatal(err)
	}

	router := router.NewRouter(entranceApp)

	log.Fatal(http.ListenAndServe(":9090", router))
}
