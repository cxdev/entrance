package main

import (
	"encoding/json"
	entrance "entrance/backend"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RouteManager struct {
	app *entrance.App
}

type CommonResponse struct {
	Data interface{}
}

func (routeManager *RouteManager) ResponseJson(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(CommonResponse{data})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (routeManager *RouteManager) PONG(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	routeManager.ResponseJson(w, "PONG")
}

func (routeManager *RouteManager) ListCommands(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	commands := routeManager.app.Commands(nil)
	routeManager.ResponseJson(w, commands)
}

func (routeManager *RouteManager) CheckCommand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	command := routeManager.app.Command(id)
	routeManager.ResponseJson(w, command)
}

func (routeManager *RouteManager) ExecuteCommand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if r.Body == nil {
		http.Error(w, "No request body", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	arguments := string(body)
	jobID := routeManager.app.AddJob(id, arguments)
	routeManager.ResponseJson(w, jobID)
}

func (routeManager *RouteManager) ListJobs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jobs := routeManager.app.Jobs(nil)
	routeManager.ResponseJson(w, jobs)
}

func (routeManager *RouteManager) CheckJob(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	idStr := params.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	job := routeManager.app.Job(id)
	routeManager.ResponseJson(w, job)
}

func (routeManager *RouteManager) SetupRoutes(router *httprouter.Router) {
	router.GET("/ping", routeManager.PONG)
	router.GET("/command", routeManager.ListCommands)
	router.GET("/command/:id/execute", routeManager.ExecuteCommand)
	router.GET("/command/:id/info", routeManager.CheckCommand)
	router.GET("/job", routeManager.ListJobs)
	router.GET("/job/:id/info", routeManager.CheckJob)
}
