package model

import "github.com/johnearl92/xendit-ta.git/internal/model/errors"

// Comment model definition
type Comment struct {
	BaseModel
	Message        string `gorm:"type:varchar(350)"`
	IsDeleted      bool
	OrganizationID string
}

// GetFST returns field FST value.
func (p *Comment) Delete() {
	p.IsDeleted = true
}

// CommentReq comment request data structure
type CommentReq struct {
	Comment string `json:"comment"`
}

// CommentWrapper parameters model
//
// swagger:parameters CommentReq
type CommentWrapper struct {
	// in: body
	// required: true
	CommentReq CommentReq
}

func (p *CommentReq) ValidateComment() errors.JSONErrors {
	var error errors.JSONErrors
	if len(p.Comment) > 350 {
		error = errors.New()

		error.Add("400",
			map[string]string{"pointer": "/data/receipt/payment/payment_breakdown/payment_type"},
			"Invalid payment_type",
			"Comment should be less than 350 characters")
	}
	return error
}

// CommentResponse comments response
type CommentResponse struct {
	Comments []string `json:"comments"`
}

// CommentResWrapper wrapper struct for CommentRes
//
// swagger:response CommentResponse
type CommentResWrapper struct {
	// in: body
	// required: true
	CommentResponse CommentResponse
}
