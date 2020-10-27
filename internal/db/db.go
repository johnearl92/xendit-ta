// Package this contains DB related files
package db

import (
	"fmt"
	"github.com/johnearl92/xendit-ta.git/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// DBConfig configuration type for DB
type DBConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Pool     Pool
	Migrate  bool
	LogMode  bool
}

// Pool configuration type for the pooling
type Pool struct {
	MinOpen int
	MaxOpen int
}

// NewDBConfig provides DB definition
func NewDBConfig(host string, port int, username string, password string, name string,
	minOpen int, maxOpen int, migrate bool, logMode bool) *DBConfig {
	return &DBConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Name:     name,
		Pool: Pool{
			MinOpen: minOpen,
			MaxOpen: maxOpen,
		},
		Migrate: migrate,
		LogMode: logMode,
	}
}

// NewConn Creates a database connection
func NewConn(config *DBConfig) (*gorm.DB, error) {
	log.Infoln("Creating new DB connection")
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Name,
	))

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Infoln("Connection created...")

	log.Infof("host=%s port=%d dbname=%s connection successful",
		config.Host,
		config.Port,
		config.Name)

	db.LogMode(config.LogMode)
	db.SingularTable(true)

	log.Info("DB LogMode=", config.LogMode)

	// Run migration
	log.Info("DB Migration=", config.Migrate)
	if config.Migrate {
		models := []interface{}{
			&model.Account{},
		}

		if err := db.AutoMigrate(models...).Error; err != nil {
			log.Error(err.Error())
			return nil, err
		}

	}

	return db, nil
}
