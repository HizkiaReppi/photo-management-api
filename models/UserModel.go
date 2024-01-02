package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User adalah model yang merepresentasikan data pengguna.
type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null" json:"username" form:"username" valid:"required~Username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid Email"`
	Password  string     `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password minimum is 6 characters"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:",omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

// BeforeCreate digunakan untuk validasi sebelum data User dibuat.
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	return u.validateStruct()
}

// BeforeUpdate digunakan untuk validasi sebelum data User diperbarui.
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.validateStruct()
}

// validateStruct digunakan untuk melakukan validasi struct menggunakan govalidator.
func (u *User) validateStruct() error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
