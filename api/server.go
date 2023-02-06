package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)


type Item struct {
	UD    uuid.UUID `json:"id"`
	Name string     `json:"name"`
}

type Server struct {
	*mux.Router

	shoppingItems []Item
}

func newServer() *Server {
	s := &Server{
		Router:         mux.NewRouter(),
		shoppingItems: []item{},
	}
	return s
}

func (s * Server) createShoppingItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.UD = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}