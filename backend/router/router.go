package router

import (
	"encoding/json"
	entrance "entrance/backend"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Router struct {
	httprouter.Router
	app entrance.App
}

type CommonResponse struct {
	Data interface{}
}

type AddCommandReqBody struct {
	Name        string                    `json:"name"`
	CommandType entrance.CommandType      `json:"command_type"`
	Segments    []entrance.CommandSegment `json:"segments"`
}

func NewRouter(app entrance.App) *Router {
	router := Router{*httprouter.New(), app}
	router.SetupRoutes()
	return &router
}

func (router *Router) ResponseJson(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(CommonResponse{data})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (router *Router) PONG(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	router.ResponseJson(w, "PONG")
}

func (router *Router) AddCommand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if r == nil || r.Body == nil {
		http.Error(w, "No request body", http.StatusBadRequest)
		return
	}
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body AddCommandReqBody
	err = json.Unmarshal(bodyData, &body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cid, err := router.app.CreateCommand(body.Name, body.CommandType, body.Segments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	router.ResponseJson(w, cid)
}

func (router *Router) ListCommands(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	commands, err := router.app.Commands(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	router.ResponseJson(w, commands)
}

func (router *Router) CheckCommand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id64, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := uint(id64)
	command, err := router.app.Command(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	router.ResponseJson(w, command)
}

func (router *Router) ExecuteCommand(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id64, err := strconv.ParseUint(params.ByName("id"), 10, 64)
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
	id := uint(id64)
	var arguments entrance.Arguments
	err = json.Unmarshal(body, &arguments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	job, err := router.app.CreateJob(id, &arguments)

	// error handle for execute
	go router.app.Execute(job)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	router.ResponseJson(w, job.ID)
}

func (router *Router) ListJobs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jobs, err := router.app.Jobs(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	router.ResponseJson(w, jobs)
}

func (router *Router) CheckJob(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id64, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := uint(id64)
	job, err := router.app.Job(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	router.ResponseJson(w, job)
}

func (router *Router) ReadJobResult(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id64, err := strconv.ParseUint(params.ByName("id"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := uint(id64)
	job, err := router.app.Job(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resultData, err := router.app.GetResult(job)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	router.ResponseJson(w, resultData)
}

func (router *Router) SetupRoutes() {
	router.GET("/ping", router.PONG)
	router.POST("/admin/command/add", router.AddCommand)
	router.GET("/command", router.ListCommands)
	router.POST("/command/:id/execute", router.ExecuteCommand)
	router.GET("/command/:id/info", router.CheckCommand)
	router.GET("/job", router.ListJobs)
	router.GET("/job/:id/info", router.CheckJob)
	router.GET("/job/:id/result", router.ReadJobResult)
}
