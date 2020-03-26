package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kradnoel/teka/internal/repository"
	"github.com/kradnoel/teka/internal/types"
	"github.com/kradnoel/teka/internal/utils"
)

// Handlers for project

var GetProjects = func(w http.ResponseWriter, r *http.Request) {

	data := repository.GetProjects()
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var CreateProject = func(w http.ResponseWriter, r *http.Request) {

	project := types.Project{}

	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		log.Panic(err)
	}
	project.Guid = utils.GenerateUUID()
	project.Completed = false

	//fmt.Println("Guid => " + project.Guid.String())
	//fmt.Println("Title => " + project.Title)
	//fmt.Println("Description => " + project.Description)

	data := repository.CreateProject(project)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var DeleteProject = func(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := repository.DeleteProject(id)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetProject = func(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := repository.GetProject(id)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetProjectStatus = func(w http.ResponseWriter, r *http.Request) {
	data := repository.GetProjectStatus()
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var SeedProject = func(w http.ResponseWriter, r *http.Request) {

	repository.Seed()
	data := repository.GetProjects()
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

// Handlers for task

var CreateTask = func(w http.ResponseWriter, r *http.Request) {

	formTask := types.FormTask{}

	err := json.NewDecoder(r.Body).Decode(&formTask)

	if err != nil {
		log.Panic(err)
	}

	task := types.Task{
		Guid:        utils.GenerateUUID(),
		Description: formTask.Description,
		ProjectId:   utils.ParseToUUID(formTask.ProjectId),
	}

	//fmt.Println("Guid => " + project.Guid.String())
	//fmt.Println("Title => " + project.Title)
	//fmt.Println("Description => " + project.Description)

	data := repository.CreateTask(task)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var DeleteTask = func(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := repository.DeleteTask(id)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

// Handlers for subtask

var SeedTask = func(w http.ResponseWriter, r *http.Request) {
	repository.SeedTasks()
	data := repository.GetTasks("c889b8b7-e134-450f-a9c0-0ce2e468d80e")
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetTask = func(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := repository.GetTasks(id)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var SeedSubTask = func(w http.ResponseWriter, r *http.Request) {
	repository.SeedSubTasks()
	data := repository.GetSubTasks("8c53e53e-0aa1-4465-8572-7aa1124284ce")
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}

var GetSubTask = func(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	data := repository.GetSubTasks(id)
	resp := utils.Message(true, "sucess")
	resp["data"] = data
	utils.Respond(w, resp)
}
