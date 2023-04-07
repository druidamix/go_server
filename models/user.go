package model

import (
	"time"

	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	ID                int    `gorm:"primaryKey;auto_increment;"`
	User              string `gorm:"type:varchar(100);"`
	Password          string `gorm:"type:varchar(100);"`
	Name              string `gorm:"type:varchar(100);"`
	Address           string `gorm:"type:varchar(200);"`
	Mail              string `gorm:"type:varchar(100);"`
	Lastupdate        time.Time
	Station_code_id   string `gorm:"type:varchar(100);"`
	First_login       int
	Token             string `gorm:"type:text"`
	Token_date        time.Time
	Token_date_string string
	Station_code      string
	Activated_at      time.Time
	Created_at        time.Time
	Updated_at        time.Time
	Deleted_at        time.Time
	Redundant_token   string `gorm:"type:TEXT;"`
}
