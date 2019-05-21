package main

import (
	"context"
	"entrance/backend/execute"
	"entrance/backend/router"
	"entrance/backend/storage"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Context struct {
	pidFile string
	dbFile  string
	logDir  string
}

type EntranceApp struct {
	storage.StorageService
	execute.ExecuteService
}

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "", "set configuration file")
	flag.StringVar(&configPath, "conf", "", "set configuration file")
	flag.Parse()
}

func createContext(rootDir string) *Context {
	pidFile := path.Join(rootDir, "pid")
	dbFile := path.Join(rootDir, "db")
	logDir := path.Join(rootDir, "logs")

	return &Context{pidFile, dbFile, logDir}
}

func prepareDir(context *Context) error {
	return os.MkdirAll(context.logDir, os.ModePerm)
}

func checkPid(pidFile string) error {
	file, err := os.Stat(pidFile)
	if os.IsNotExist(err) {
		return nil
	}
	if file.IsDir() == false {
		content, err := ioutil.ReadFile(pidFile)
		if err == nil {
			// TODO: check pid if still running
			return errors.New("Entrance is running on pid: " + string(content))
		}
	}

	return errors.New("The path of pid file is invalid")
}

func writePid(pidFile string) error {
	pid := []byte(strconv.Itoa(os.Getpid()))
	return ioutil.WriteFile(pidFile, pid, 0555)
}
func removePid(pidFile string) error {
	return os.Remove(pidFile)
}

func createEntranceApp(context *Context) (*EntranceApp, error) {
	db, err := gorm.Open("sqlite3", context.dbFile)
	if err != nil {
		return nil, err
	}

	storageService, err := storage.NewStorageService(db)
	if err != nil {
		return nil, err
	}

	executeService := execute.NewExecuteService(context.logDir)

	return &EntranceApp{*storageService, *executeService}, nil
}

func startHttpServer(appConfig *AppConfig, entranceApp *EntranceApp) *http.Server {
	address := ":" + strconv.Itoa(appConfig.Port)
	handler := router.NewRouter(entranceApp)
	server := &http.Server{Addr: address, Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	log.Println("Server start, on port", address)

	return server
}

func setupShutdownProcess(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Println("Failed to gracefully shutdown:", err)
	}
}

func main() {

	// Read config
	appConfig, err := loadAppConfig(configPath)
	handleError(err)

	// context setup
	context := createContext(appConfig.RootDir)
	err = prepareDir(context)
	handleError(err)

	// check if running: Yes to abort this start and message
	err = checkPid(context.pidFile)
	handleError(err)

	// init db
	entranceApp, err := createEntranceApp(context)
	handleError(err)

	// Add cmd? do it

	// WritePid
	err = writePid(context.pidFile)
	handleError(err)

	// Start
	server := startHttpServer(appConfig, entranceApp)

	// When shutdown
	setupShutdownProcess(server)
	removePid(context.pidFile)
	log.Println("Server shutdown")
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
