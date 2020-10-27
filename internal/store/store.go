package store

import (
	"errors"
	"github.com/johnearl92/xendit-ta.git/internal/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Store represents data access layers
type Store interface {
	Commit() error
	Rollback() error
	Run(func() error) error
	WithValue(func() (interface{}, error)) (interface{}, error)
}

// BaseStore base Store implementation
type BaseStore struct {
	DB *gorm.DB
}

// Apply apply options to gorm
func (b *BaseStore) Apply(opts ListOpts, db *gorm.DB) *gorm.DB {
	if opts == nil {
		return db
	}

	if opts.Offset() != nil {
		db = db.Offset(*opts.Offset())
	}

	if opts.Max() != nil {
		db = db.Limit(*opts.Max())
	}

	if len(opts.Sort()) > 0 {
		db = db.Order(opts.SortString())
	}

	return db
}

// Begin initiate connection
func (b *BaseStore) Begin() Store {
	return &BaseStore{
		DB: b.DB.Begin(),
	}
}

// Commit commit transaction to db
func (b *BaseStore) Commit() error {
	return b.DB.Commit().Error
}

// Rollback revert transaction
func (b *BaseStore) Rollback() error {
	return b.DB.Rollback().Error
}

// Run execute store
func (b *BaseStore) Run(closure func() error) error {
	err := closure()

	if err != nil {
		log.Error(err.Error())
		return b.DB.Rollback().Error
	}

	return b.DB.Commit().Error
}

// WithValue value
func (b *BaseStore) WithValue(closure func() (interface{}, error)) (interface{}, error) {
	value, err := closure()

	if err != nil {
		log.Error(err.Error())
		return nil, b.DB.Rollback().Error
	}

	if err := b.DB.Commit().Error; err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return value, nil
}

// Get getter
func (b *BaseStore) Get(id string, m interface{}, opts GetOpts) (interface{}, error) {
	return b.Find(b.DB.Where("id = ?", id), m, opts)
}

// Find find specific record
func (b *BaseStore) Find(gormDB *gorm.DB, m interface{}, opts GetOpts) (interface{}, error) {
	db := gormDB

	if opts != nil {
		for _, load := range opts.Preload() {
			db = db.Preload(load)
		}
	}

	err := db.First(m).Error

	switch err {
	case nil:
		return m, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.New("Not Found")
	default:
		return nil, err
	}
}

// All get all record
func (b *BaseStore) All(model interface{}, elems interface{}) error {
	return b.DB.Model(model).Find(elems).Error
}

// List list all related records
func (b *BaseStore) List(list List, opts ListOpts) error {
	return b.FindAll(b.DB, list, opts)
}

// FindAll list all related records
func (b *BaseStore) FindAll(gormDB *gorm.DB, list List, opts ListOpts) error {
	total := 0
	db := gormDB

	if opts != nil {
		for _, load := range opts.Preload() {
			db = db.Preload(load)
		}
	}

	if err := db.Model(list.Model()).Count(&total).Error; err != nil {
		log.Error(err.Error())
		return err
	}

	list.SetTotal(total)
	return b.Apply(opts, db).Find(list.ItemsPtr()).Error
}

// Delete removes specific record
func (b *BaseStore) Delete(m interface{}) error {
	db := b.DB.Model(m).
		Where("id = ?", m.(model.Model).GetID()).
		Delete(m)

	if db.RowsAffected == 0 {
		return errors.New("Concurrent Update")
	}

	return db.Error
}

// Update update a specific record
func (b *BaseStore) Update(m interface{}) error {
	db := b.DB.Model(m).
		Where("version = ?", m.(model.Model).GetVersion()).
		Updates(m)

	if db.RowsAffected == 0 {
		return errors.New("Concurrent Update")
	}

	return db.Error
}
