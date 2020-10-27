// Package service Official-Receipt API.
package service

import (
	"github.com/johnearl92/xendit-ta.git/internal/model"
	"github.com/johnearl92/xendit-ta.git/internal/store"
)

// XenditServiceImpl EOR service definition
type XenditServiceImpl struct {
	store store.AccountStore
}

// NewXenditService XenditServiceImpl provider
func NewXenditService(store store.AccountStore) *XenditServiceImpl {
	return &XenditServiceImpl{
		store: store,
	}
}

// XenditService Receiptservice interface
type XenditService interface {
	Create(receipt *model.Account, opts store.GetOpts) error
	Update(receipt *model.Account) error
	Get(id string, opts store.GetOpts) (*model.Account, error)
}

// Create saves receipts in database
func (p *XenditServiceImpl) Create(acoount *model.Account, opts store.GetOpts) error {
	return p.store.Create(acoount, opts)
}

// Update updates the specific receipt
func (p *XenditServiceImpl) Update(account *model.Account) error {
	return p.store.Update(account)
}

// Get find a specific receipt in database
func (p *XenditServiceImpl) Get(id string, opts store.GetOpts) (*model.Account, error) {
	return p.store.Get(id, opts)
}
