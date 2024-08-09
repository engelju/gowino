package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{listenAddr: listenAddr}
}

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		println("Error encoding JSON response: ", err)
	}
}

func (s *APIServer) Start() error {
	r := http.NewServeMux()

	r.HandleFunc("GET /user", s.handleUser)
	r.HandleFunc("GET /user/{id}", s.handleGetUser)
	r.HandleFunc("POST /user/{id}/create", s.handleCreateUser)
	r.HandleFunc("DELETE /user/{id}/delete", s.handleDeleteUser)

	r.HandleFunc("GET /wines", s.handleWines)
	r.HandleFunc("GET /wines/{id}", s.handleGetWine)
	r.HandleFunc("POST /wines/create", s.handleCreateWine)
	r.HandleFunc("DELETE /wines/delete", s.handleDeleteWine)

	log.Print("Starting API server on ", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, r)
}

// User handling

func (s *APIServer) handleUser(w http.ResponseWriter, r *http.Request) {
	//WriteJSON(w, http.StatusOK, map[string]string{"message": "Hello, world!"})
	WriteJSON(w, http.StatusOK, NewUser("Test User", "test@example.org"))
}

func (s *APIServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid ID", "message": "ID must be an integer, ID provided was `" + r.PathValue("id") + "`"})
		//http.NotFound(w, r)
		return
	}
	WriteJSON(w, http.StatusOK, NewUserWithId(id, "Test User", "test@example.org"))
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
}

// handle wines

func (s *APIServer) handleWines(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, map[string]string{"message": "Hello, world!"})
}

func (s *APIServer) handleGetWine(w http.ResponseWriter, r *http.Request) {
}

func (s *APIServer) handleCreateWine(w http.ResponseWriter, r *http.Request) {
}

func (s *APIServer) handleDeleteWine(w http.ResponseWriter, r *http.Request) {
}
