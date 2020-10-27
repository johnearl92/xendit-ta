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
	receipt, err := p.Find(db, &model.Comment{}, opts)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return receipt.(*model.Comment), nil
}
