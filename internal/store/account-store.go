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

// fFindByOrNum ind latest receipt by the given OR number
func (p *accountStore) FindByOrNum(orNum string) (*model.Account, error) {
	log.Debug("FindByOrNum Invoke")
	db := p.DB.Order("version").Where("or_number = ?", orNum)
	var account model.Account

	if err := db.Last(&account).Error; err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debug("FindByOrNum End")
	return &account, nil
}

func (p *accountStore) Get(id string, opts GetOpts) (*model.Account, error) {
	db := p.DB.Where("id = ?", id)
	receipt, err := p.Find(db, &model.Account{}, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return receipt.(*model.Account), nil
}
