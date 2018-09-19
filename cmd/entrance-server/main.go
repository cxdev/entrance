package main

import (
	"entrance/backend"
	"entrance/backend/exec"
	"entrance/backend/platform"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func initApp() *entrance.App {
	var TestDB, _ = platform.CreateDB("file::memory:?mode=memory")
	commandService := entrance.CommandService{TestDB}
	jobService := entrance.JobService{TestDB, exec.ExecContextBuilder{"/tmp/entrance"}}
	return &entrance.App{commandService, jobService}
}

func main() {
	entranceApp := initApp()
	routerManager := RouteManager{entranceApp}

	router := httprouter.New()
	routerManager.SetupRoutes(router)

	log.Fatal(http.ListenAndServe(":9090", router))
}
