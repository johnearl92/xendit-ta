package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/johnearl92/xendit-ta.git/internal/model"
	log "github.com/sirupsen/logrus"
)

func RunMigration(db *gorm.DB) error {
	log.Info("Migrating Predefinec Organizations")

	org := &model.Organization{
		Name: "xendit",
	}

	err := db.FirstOrCreate(org).Error

	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Infof("Migrating Predefineed Members")
	memberNames := [5]string{"john", "doe", "smith", "elastic", "invisible"}
	if db.Find(&model.Account{}).RowsAffected == 0 {
		for i := 0; i < 5; i++ {
			log.Infof("Creating account for %s", memberNames[i])
			err := db.Create(&model.Account{
				Username:       memberNames[i],
				FollowedNum:    int32(i),
				FollowersNum:   int32(i),
				OrganizationID: org.ID,
				AvatarURL:      fmt.Sprintf("http://%s.com", memberNames[i]),
			}).Error

			if err != nil {
				log.Error(err.Error())
				return err
			}
		}
	}

	return nil
}
