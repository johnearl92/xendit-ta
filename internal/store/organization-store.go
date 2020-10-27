package store

import (
	"github.com/jinzhu/gorm"
	"github.com/johnearl92/xendit-ta.git/internal/model"
	log "github.com/sirupsen/logrus"
)

// NewOrganizationStore ...
func NewOrganizationStore(db *gorm.DB) OrganizationStore {
	return &organizationStore{BaseStore{
		DB: db,
	}}
}

type organizationStore struct {
	BaseStore
}

// OrganizationStore ...
type OrganizationStore interface {
	Create(organization *model.Organization, opts GetOpts) error
	Update(organization *model.Organization) error
	Get(id string, opts GetOpts) (*model.Organization, error)
}

func (p *organizationStore) Create(organization *model.Organization, opts GetOpts) error {
	db := p.DB.Create(organization)
	err := db.Error

	return err
}

func (p *organizationStore) Update(organization *model.Organization) error {
	db := p.BaseStore.Update(organization)
	return db
}

func (p *organizationStore) Get(id string, opts GetOpts) (*model.Organization, error) {
	db := p.DB.Where("id = ?", id)
	receipt, err := p.Find(db, &model.Comment{}, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return receipt.(*model.Organization), nil
}
