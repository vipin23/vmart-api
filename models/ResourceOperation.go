package models

import (
	"time"

	"github.com/vipin23/vmart-api/api_server"
)

type ResourceOperation struct {
	BaseModel
	Path        string  `gorm:"size:1024;not null;index" json:"path"`
	HTTPMethod  string  `gorm:"size:10;not null;index" json:"http_method"`
	Role        []*Role `gorm:"many2many:role_resourcesoperations;"`
	Description string  `gorm:"size:2048;not null;index" json:"description"`
}

func (r *ResourceOperation) Prepare() {
	r.ID = 0
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

func (u *ResourceOperation) Save() (*ResourceOperation, error) {
	var err error
	err = api_server.DefaultServer.DB.Debug().Create(&u).Error
	if err != nil {
		return &ResourceOperation{}, err
	}
	return u, nil
}

func (r *ResourceOperation) Upsert() (*ResourceOperation, error) {
	existing, err := r.FindByMethodAndPath(r.HTTPMethod, r.Path)
	if err != nil {
		return r.Save()
	}
	return r.Update(existing.ID)
}

func (r *ResourceOperation) FindByMethodAndPath(method, path string) (*ResourceOperation, error) {
	existing := ResourceOperation{}
	err := api_server.DefaultServer.DB.Debug().Where("http_method = ? AND path = ?", r.HTTPMethod, r.Path).First(&existing).Error
	if err != nil {
		return &ResourceOperation{}, err
	}
	return &existing, err
}

func (r *ResourceOperation) Update(uid uint64) (*ResourceOperation, error) {
	db := api_server.DefaultServer.DB.Debug().Model(&ResourceOperation{}).Where("id = ?", uid).
		Take(&ResourceOperation{}).UpdateColumns(
		map[string]interface{}{
			"description": r.Description,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &ResourceOperation{}, db.Error
	}
	// This is the display the updated user
	err := api_server.DefaultServer.DB.Debug().Model(&ResourceOperation{}).Where("id = ?", uid).Take(&r).Error
	if err != nil {
		return &ResourceOperation{}, err
	}
	return r, nil
}
