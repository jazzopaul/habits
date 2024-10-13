package habits_hello_world

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jazzopaul/habits/habits"
	"github.com/jazzopaul/habits/hello_world"
)

type (
	Controller struct {
		habitsSvc     *habits.Service
		helloWorldSvc *hello_world.Service
	}
)

func NewController(
	habitsSvc *habits.Service,
	helloWorldSvc *hello_world.Service,
) *Controller {

	fmt.Println("New Hello World Controller")
	return &Controller{
		habitsSvc:     habitsSvc,
		helloWorldSvc: helloWorldSvc,
	}
}

func (c *Controller) Dispatch(r chi.Router) {
	r.Use(corsMiddleware)
	r.Route("/hello-world", func(r chi.Router) {
		r.Get("/", c.helloWorldHandler)
		r.Post("/submit", c.submitHandler)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *Controller) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}

func (c *Controller) submitHandler(w http.ResponseWriter, r *http.Request) {

	data := struct {
		Name string `json:"name"`
	}{}

	// Decode JSON from the request body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Print the data to the terminal
	fmt.Printf("Received data: %+v\n", data)

	// Send a response back to the frontend
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Form data received"))
}
