// Package service Official-Receipt API.
package service

import (
	"github.com/johnearl92/xendit-ta.git/internal/model"
	"github.com/johnearl92/xendit-ta.git/internal/store"
	"strings"
)

// xenditService EOR service definition
type xenditService struct {
	acctStore    store.AccountStore
	commentStore store.CommentStore
	orgStore     store.OrganizationStore
}

// NewXenditService xenditService provider
func NewXenditService(acctStore store.AccountStore, commentStore store.CommentStore, orgStore store.OrganizationStore) XenditService {
	return &xenditService{
		acctStore:    acctStore,
		commentStore: commentStore,
		orgStore:     orgStore,
	}
}

// XenditService Receiptservice interface
type XenditService interface {
	CreateAccount(account *model.Account, opts store.GetOpts) error
	UpdateAccount(account *model.Account) error
	GetAccount(id string, opts store.GetOpts) (*model.Account, error)

	CreateOrganization(organization *model.Organization, opts store.GetOpts) error
	UpdateOrganization(organization *model.Organization) error
	GetOrganization(id string, opts store.GetOpts) (*model.Organization, error)
	FindByOrgName(name string, opts store.GetOpts) (*model.Organization, error)

	CreateComment(comment *model.Comment, opts store.GetOpts) error
	UpdateComment(comment *model.Comment) error
	GetComment(id string, opts store.GetOpts) (*model.Comment, error)
	FindCommentsByOrg(orgName string, opts store.ListOpts) (*store.CommentList, error)
}

// Create saves receipts in database
func (p *xenditService) CreateAccount(acoount *model.Account, opts store.GetOpts) error {
	return p.acctStore.Create(acoount, opts)
}

// Update updates the specific receipt
func (p *xenditService) UpdateAccount(account *model.Account) error {
	return p.acctStore.Update(account)
}

// Get find a specific receipt in database
func (p *xenditService) GetAccount(id string, opts store.GetOpts) (*model.Account, error) {
	return p.acctStore.Get(id, opts)
}

// Create saves receipts in database
func (p *xenditService) CreateOrganization(organization *model.Organization, opts store.GetOpts) error {
	return p.orgStore.Create(organization, opts)
}

// Update updates the specific receipt
func (p *xenditService) UpdateOrganization(organization *model.Organization) error {
	return p.orgStore.Update(organization)
}

// Get find a specific receipt in database
func (p *xenditService) GetOrganization(id string, opts store.GetOpts) (*model.Organization, error) {
	return p.orgStore.Get(id, opts)
}

// Get find a specific receipt in database
func (p *xenditService) FindByOrgName(name string, opts store.GetOpts) (*model.Organization, error) {
	return p.orgStore.FindByName(name, opts)
}

// Create saves receipts in database
func (p *xenditService) CreateComment(comment *model.Comment, opts store.GetOpts) error {
	return p.commentStore.Create(comment, opts)
}

// Update updates the specific receipt
func (p *xenditService) UpdateComment(comment *model.Comment) error {
	return p.commentStore.Update(comment)
}

// Get find a specific receipt in database
func (p *xenditService) GetComment(id string, opts store.GetOpts) (*model.Comment, error) {
	return p.commentStore.Get(id, opts)
}

// FindCommentsByOrg find all comments by org
func (p *xenditService) FindCommentsByOrg(orgName string, opts store.ListOpts) (*store.CommentList, error) {
	org, err := p.FindByOrgName(strings.ToLower(orgName), nil)
	if err != nil {
		return nil, err
	}
	return p.commentStore.ListByOrgID(org.ID, opts)
}
