package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uuid.UUID `gorm:"type:char(36);primary_key;"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Username   string    `json:"username"`
	Email      string    `gorm:"unique"`
	Password   string    `json:"-"`
	IsVerified bool      `json:"isVerified"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
