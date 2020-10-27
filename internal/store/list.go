package store

import (
	"github.com/johnearl92/xendit-ta.git/internal/model"
)

// List contains list of models
type List interface {
	Total() int
	SetTotal(int)
	ItemsPtr() interface{}
	Model() model.Model
}

// BaseList base List implementation
type BaseList struct {
	total int
}

// Total getter
func (b *BaseList) Total() int {
	return b.total
}

// SetTotal setter
func (b *BaseList) SetTotal(arg int) {
	b.total = arg
}

// Init initializer
func (b *BaseList) Init() {
	b.SetTotal(0)
}
