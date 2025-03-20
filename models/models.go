package models

import (
	"gorm.io/gorm"
)

type Song struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Group    string `json:"group"`
	SongName string `json:"song"`
	Lyrics   string `json:"lyrics"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Song{})
}
