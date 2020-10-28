package store

import (
	"github.com/jinzhu/gorm"
	"github.com/johnearl92/xendit-ta.git/internal/model"
	log "github.com/sirupsen/logrus"
)

// NewCommentStore ...
func NewCommentStore(db *gorm.DB) CommentStore {
	return &commentStore{BaseStore{
		DB: db,
	}}
}

type commentStore struct {
	BaseStore
}

// CommentStore ...
type CommentStore interface {
	Create(comment *model.Comment, opts GetOpts) error
	Update(comment *model.Comment) error
	Get(id string, opts GetOpts) (*model.Comment, error)
	ListByOrgID(orgID string, opts ListOpts) (*CommentList, error)
}

func (p *commentStore) Create(comment *model.Comment, opts GetOpts) error {
	db := p.DB.Create(comment)
	err := db.Error

	return err
}

func (p *commentStore) Update(comment *model.Comment) error {
	db := p.BaseStore.Update(comment)
	return db
}

func (p *commentStore) Get(id string, opts GetOpts) (*model.Comment, error) {
	db := p.DB.Where("id = ?", id)
	comment, err := p.Find(db, &model.Comment{}, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return comment.(*model.Comment), nil
}

func (p *commentStore) ListByOrgID(orgID string, opts ListOpts) (*CommentList, error) {
	list := NewCommentList()
	db := p.DB.Where("organization_id = ? AND is_deleted = false", orgID)
	err := p.BaseStore.FindAll(db, list, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return list, nil
}
