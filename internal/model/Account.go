package model

type Account struct {
	BaseModel
	Username       string `gorm:"type:varchar(20)"`
	AvatarURL      string
	FollowersNum   int32
	FollowedNum    int32
	OrganizationID string
}

// AccountReq parameters model
//
// swagger:parameters Account
type AccountReq struct {
	// in: body
	// required: true
	Account Account
}

// AccountResponse comments response
type AccountResponse struct {
	Accounts []Account `json:"accounts"`
}

// AccountResWrapperv wrapper struct for AccountResponse
//
// swagger:response AccountResponse
type AccountResWrapper struct {
	// in: body
	// required: true
	AccountResponse AccountResponse
}
