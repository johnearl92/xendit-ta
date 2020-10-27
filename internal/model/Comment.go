package model

import "github.com/johnearl92/xendit-ta.git/internal/model/errors"

type Comment struct {
	BaseModel
	Message        string `gorm:"type:varchar(350)"`
	IsDeleted      bool
	OrganizationID string
}

// CommentReq parameters model
//
// swagger:parameters Comment
type CommentReq struct {
	// in: body
	// required: true
	Comment string `json:"comment"`
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
