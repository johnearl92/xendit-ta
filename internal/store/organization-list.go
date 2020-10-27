package store

import "github.com/johnearl92/xendit-ta.git/internal/model"

// NewOrganizationList OrganizationList provider
func NewOrganizationList() *OrganizationList {
	list := &OrganizationList{}
	list.Init()
	return list
}

// OrganizationList list definition
type OrganizationList struct {
	BaseList
	items []*model.Organization
}

// Model model to use for db
func (p *OrganizationList) Model() model.Model {
	return &model.Organization{}
}

// ItemsPtr pointer items in the list
func (p *OrganizationList) ItemsPtr() interface{} {
	return &p.items
}

// Items items in the list
func (p *OrganizationList) Items() []*model.Organization {
	return p.items
}
