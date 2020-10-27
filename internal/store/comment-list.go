package store

import "github.com/johnearl92/xendit-ta.git/internal/model"

// NewCommentList CommentList provider
func NewCommentList() *CommentList {
	list := &CommentList{}
	list.Init()
	return list
}

// CommentList list definition
type CommentList struct {
	BaseList
	items []*model.Comment
}

// Model model to use for db
func (p *CommentList) Model() model.Model {
	return &model.Comment{}
}

// ItemsPtr pointer items in the list
func (p *CommentList) ItemsPtr() interface{} {
	return &p.items
}

// Items items in the list
func (p *CommentList) Items() []*model.Comment {
	return p.items
}
