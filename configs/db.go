package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteConnection(name string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(name), &gorm.Config{})
}
