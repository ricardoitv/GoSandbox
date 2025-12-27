package v1

import (
	"log"
	"net/http"
)

type TodoApp struct{}

func NewTodoApp() *TodoApp {
	return &TodoApp{}
}

func (t *TodoApp) Run() {
	router := http.NewServeMux()
	router.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Got GET items/" + id))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Listening for connections")
	server.ListenAndServe()
}
