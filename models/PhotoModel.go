package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo adalah model yang merepresentasikan data foto.
type Photo struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title" form:"title" valid:"required~Photo Title is Required"`
	Caption   string    `json:"caption" form:"caption"`
	PhotoURL  string    `json:"photo_url" form:"photo_url" valid:"required~Photo URL is Required"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// BeforeCreate digunakan untuk validasi sebelum data Photo dibuat.
func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	return p.validateStruct()
}

// BeforeUpdate digunakan untuk validasi sebelum data Photo diperbarui.
func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	return p.validateStruct()
}

// validateStruct digunakan untuk melakukan validasi struct menggunakan govalidator.
func (p *Photo) validateStruct() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
