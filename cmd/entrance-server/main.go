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
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/yaml.v2"
)

type EntranceBackend struct {
	storage.StorageService
	execute.ExecuteService
}

type dbInitConfig struct {
	CreateNewDB  bool `yaml:"create_new_db"`
	InitCommands bool `yaml:"init_commands"`
}

type config struct {
	Version           string              `yaml:"version"`
	LogRootPath       string              `yaml:"log_root_path"`
	SqlitePath        string              `yaml:"sqlite_path"`
	WhenDbExistConfig *dbInitConfig       `yaml:"when_db_exist"`
	Commands          *[]entrance.Command `yaml:"commands"`
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

func getDBInitConfig(config *config) *dbInitConfig {
	_, err := os.Stat(config.SqlitePath)
	isDbExist := (err == nil)
	if isDbExist {
		return config.WhenDbExistConfig
	}
	return &dbInitConfig{false, true}
}

func removeAndBackup(path string) error {
	// RFC3339: "2006-01-02T15:04:05Z07:00"
	backupPostfix := time.Now().Format(time.RFC3339)
	newLocation := path + "." + backupPostfix
	err := os.Rename(path, newLocation)
	return err
}

func initCommands(commands *[]entrance.Command, entranceBackend *EntranceBackend) error {
	if commands != nil {
		for _, command := range *commands {
			// TODO: think more: create or update or need error?
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

	dbInitConfig := getDBInitConfig(config)

	if dbInitConfig.CreateNewDB {
		removeAndBackup(config.SqlitePath)
	}

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

	if dbInitConfig.InitCommands {
		err = initCommands(config.Commands, entranceApp)
		if err != nil {
			log.Fatal(err)
		}
	}

	router := router.NewRouter(entranceApp)

	log.Fatal(http.ListenAndServe(":9090", router))
}
