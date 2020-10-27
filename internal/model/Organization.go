package model

type Organization struct {
	BaseModel
	Name     string `gorm:"type:varchar(50)"`
	Members  []Account
	Comments []Comment
}
