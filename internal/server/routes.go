package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var value = "25"

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", s.HelloWorldHandler)
	r.Get("/get", s.showValue)
	r.Get("/health", s.healthHandler)
	r.Get("/data", s.GetData)
	return r
}
func (s *Server) showValue(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["currentValue"] = value
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		log.Printf("error handling JSON marshal for ShowValue. Err: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResp)

}
func (s *Server) GetData(w http.ResponseWriter, r *http.Request) {

	value  = r.URL.Query().Get("sensor")
	fmt.Println("Input from the senors are :- ", value)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	//jsonResp, _ := json.Marshal(s.db.Health())
	jsonResp, _ := json.Marshal("Working")
	_, _ = w.Write(jsonResp)
}
