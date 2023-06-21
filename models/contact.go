package models

import (
	"time"
)

type Contact struct{
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaKontak string `gorm:"type:varchar(300)" json:"name"`
	Gender string `gorm:"type:varchar(300)" json:"gender"`
	Phone string `gorm:"type:varchar(50)" json:"phone"`
	Email string `gorm:"type:varchar(50)" json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}