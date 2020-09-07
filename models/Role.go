package models

type Role struct {
	BaseModel
	Name              string              `gorm:"size:100;not null;unique" json:"name"`
	ResourceOperation []ResourceOperation `gorm:"many2many:role_resourcesoperations;"`
}
