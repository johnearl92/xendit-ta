package model

type Comment struct {
	BaseModel
	Message        string `gorm:"type:varchar(350)"`
	IsDeleted      bool
	OrganizationID string
}
