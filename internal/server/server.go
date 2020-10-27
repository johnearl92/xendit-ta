// Package this contains server related files
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/johnearl92/xendit-ta.git/internal/handler"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// APIServer server definition
type APIServer struct {
	*mux.Router
	*Config
	handlers []handler.Routable
	// filters  []mux.MiddlewareFunc
}

// Config server config definition
type Config struct {
	Host           string
	Port           int
	Spec           string
	AllowedHeaders []string
	AllowedOrigins []string
	AllowedMethods []string
}

// NewAPIServerConfig provides server config implementation
func NewAPIServerConfig(host string, port int, spec string, allowedOrigins, allowedHeaders, allowedMethods []string) *Config {
	return &Config{
		Host:           host,
		Port:           port,
		Spec:           spec,
		AllowedOrigins: allowedOrigins,
		AllowedHeaders: allowedHeaders,
		AllowedMethods: allowedMethods,
	}
}

// NewAPIServer provides server implementation
func NewAPIServer(apiserverConfig *Config, router *mux.Router, handlers []handler.Routable) *APIServer {
	return &APIServer{
		Config:   apiserverConfig,
		Router:   router,
		handlers: handlers,
	}
}

func (s *Config) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Run runs the server and it's services
func (a *APIServer) Run() error {
	a.serveSwagger()
	a.registerRoutes()
	a.registerFilter()
	address := a.Config.String()
	log.WithField("address", address).Info("API Server started")
	return http.ListenAndServe(address, a.Router)
}

// serveSwagger register swagger
func (a *APIServer) serveSwagger() {
	a.Router.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui", http.FileServer(http.Dir("swagger"))))
	log.Info("/swagger-ui/#/ registered")
}

// registerRoutes register endpoints
func (a *APIServer) registerRoutes() {
	for _, routable := range a.handlers {
		routable.Register(a.Router)
	}
}

// registerFilter register Filters
func (a *APIServer) registerFilter() {
	corsFilter := handlers.CORS(
		handlers.AllowedOrigins(a.AllowedOrigins),
		handlers.AllowedHeaders(a.AllowedHeaders),
		handlers.AllowedMethods(a.AllowedMethods),
	)
	a.Use(corsFilter)
}
