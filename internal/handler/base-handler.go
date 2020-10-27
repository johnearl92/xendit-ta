package handler

import (
	"github.com/gorilla/mux"
)

// Routable represents the base for handlers.
type Routable interface {
	Register(*mux.Router)
}
