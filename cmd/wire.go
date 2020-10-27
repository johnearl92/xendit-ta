//+build wireinject

package cmd

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"github.com/johnearl92/xendit-ta.git/internal/db"
	"github.com/johnearl92/xendit-ta.git/internal/handler"
	"github.com/johnearl92/xendit-ta.git/internal/server"
	"github.com/johnearl92/xendit-ta.git/internal/service"
	"github.com/johnearl92/xendit-ta.git/internal/store"
	"github.com/spf13/viper"
)

// createServer dependency injection can be set here
func createServer() (*server.APIServer, error) {
	wire.Build(
		ProvideDBConfig,
		db.NewConn,
		store.NewAccountStore,
		service.NewXenditService,

		mux.NewRouter,
		handler.NewXenditHandler,
		ProvideHandlers,
		ProvideServerConfig,
		server.NewAPIServer,
	)
	return &server.APIServer{}, nil
}

// ProvideServerConfig server configuration
func ProvideServerConfig() *server.Config {
	return server.NewAPIServerConfig(
		viper.GetString("server.host"),
		viper.GetInt("server.port"),
		viper.GetString("server.spec"),
		viper.GetStringSlice("server.cors.allowedOrigins"),
		viper.GetStringSlice("server.cors.allowedHeaders"),
		viper.GetStringSlice("server.cors.allowedMethods"),
	)
}

// ProvideDBConfig db configuration
func ProvideDBConfig() *db.DBConfig {
	return db.NewDBConfig(
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.name"),
		viper.GetInt("db.pool.minOpen"),
		viper.GetInt("db.pool.maxOpen"),
		viper.GetBool("db.migrate"),
		viper.GetBool("db.logMode"),
	)
}

// ProvideHandlers handler injection
func ProvideHandlers(p *handler.XenditHandler) []handler.Routable {
	return []handler.Routable{p}
}
