package store

import (
	"github.com/jinzhu/gorm"
	"github.com/johnearl92/xendit-ta.git/internal/model"
	log "github.com/sirupsen/logrus"
)

// NewAccountStore ...
func NewAccountStore(db *gorm.DB) AccountStore {
	return &accountStore{BaseStore{
		DB: db,
	}}
}

type accountStore struct {
	BaseStore
}

// AccountStore ...
type AccountStore interface {
	Create(account *model.Account, opts GetOpts) error
	Update(account *model.Account) error
	Get(id string, opts GetOpts) (*model.Account, error)
	ListByOrgID(orgID string, opts ListOpts) (*AccountList, error)
}

func (p *accountStore) Create(account *model.Account, opts GetOpts) error {
	db := p.DB.Create(account)
	err := db.Error

	return err
}

func (p *accountStore) Update(account *model.Account) error {
	db := p.BaseStore.Update(account)
	return db
}

func (p *accountStore) Get(id string, opts GetOpts) (*model.Account, error) {
	db := p.DB.Where("id = ?", id)
	acct, err := p.Find(db, &model.Account{}, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return acct.(*model.Account), nil
}

func (p *accountStore) ListByOrgID(orgID string, opts ListOpts) (*AccountList, error) {
	list := NewAccountList()
	db := p.DB.Where("organization_id = ?", orgID)
	err := p.BaseStore.FindAll(db, list, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return list, nil
}
