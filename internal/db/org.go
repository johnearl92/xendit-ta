package db

import (
	"github.com/jinzhu/gorm"
	"github.com/johnearl92/xendit-ta.git/internal/model"
	log "github.com/sirupsen/logrus"
)

func RunOrgMigration(db *gorm.DB) error {
	log.Info("Migrating Predefinec Organizations")

	org := &model.Organization{
		Name: "xendit",
	}

	err := db.FirstOrCreate(org).Error

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
