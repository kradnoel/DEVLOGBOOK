package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type server struct{}

func getCors() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"localhost:3000"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	return cors
}

func New() server {
	return server{}
}

func (s *server) Run() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(getCors().Handler)

	// Projects Routes
	router.Post("/projects", CreateProject)
	router.Get("/projects", GetProjects)
	router.Get("/projects/{id}", GetProject)
	router.Delete("/projects/{id}", DeleteProject)

	// Tasks Routes
	router.Post("/tasks", CreateTask)
	router.Get("/tasks/{id}", GetTask)
	router.Delete("/tasks/{id}", DeleteTask)

	//Subtasks Routes
	router.Get("/subtasks/{id}", GetSubTask)

	// Status and Seed Routes
	router.Get("/status", GetProjectStatus)
	router.Get("/seed/projects", SeedProject)
	router.Get("/seed/tasks", SeedTask)
	router.Get("/seed/subtasks", SeedSubTask)

	fmt.Println("Server starting at port :8080")
	http.ListenAndServe(":8080", router)
}

func (s *server) Execute() {
	server := New()
	server.Run()
}
