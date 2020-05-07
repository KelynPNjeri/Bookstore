package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func OpenDBConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=blog_backend password=Kelyn@1998")
}