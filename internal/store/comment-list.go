package store

import "github.com/johnearl92/xendit-ta.git/internal/model"

// NewAccountList AccountList provider
func NewAccountList() *AccountList {
	list := &AccountList{}
	list.Init()
	return list
}

// AccountList list definition
type AccountList struct {
	BaseList
	items []*model.Account
}

// Model model to use for db
func (p *AccountList) Model() model.Model {
	return &model.Account{}
}

// ItemsPtr pointer items in the list
func (p *AccountList) ItemsPtr() interface{} {
	return &p.items
}

// Items items in the list
func (p *AccountList) Items() []*model.Account {
	return p.items
}
