package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Model model interface for the BaseModel
type Model interface {
	GetID() string
	GetVersion() int
	GetCreated() time.Time
	GetUpdated() time.Time
	Validate() error
}

// BaseModel base Model implementation
type BaseModel struct {
	ID      string    `json:"-" gorm:"primary_key;type:varchar(36) not null"`
	Version int       `json:"-" gorm:"type:int not null"`
	Created time.Time `json:"-" gorm:"type:timestamp not null"`
	Updated time.Time `json:"-" gorm:"type:timestamp not null"`
}

// BeforeCreate runs before model is inserted
func (m *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	if err := scope.SetColumn("ID", uuid.New().String()); err != nil {
		return err
	}

	if err := scope.SetColumn("Version", 0); err != nil {
		return err
	}

	if err := scope.SetColumn("Created", time.Now().UTC()); err != nil {
		return err
	}

	if err := scope.SetColumn("Updated", time.Now().UTC()); err != nil {
		return err
	}

	return nil
}

// BeforeDelete runs before deleting a row
func (m *BaseModel) BeforeDelete(scope *gorm.Scope) error {
	return nil
}

// BeforeUpdate runs before updating a row
func (m *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("Version", m.Version+1); err != nil {
		return err
	}

	if err := scope.SetColumn("Updated", time.Now().UTC()); err != nil {
		return err
	}

	return nil
}

// GetID ID getter
func (m *BaseModel) GetID() string {
	return m.ID
}

// GetVersion Version getter
func (m *BaseModel) GetVersion() int {
	return m.Version
}

// GetCreated Created getter
func (m *BaseModel) GetCreated() time.Time {
	return m.Created
}

// GetUpdated Updated getter
func (m *BaseModel) GetUpdated() time.Time {
	return m.Updated
}

// Validate Validation
func (m *BaseModel) Validate() error {
	return nil
}
