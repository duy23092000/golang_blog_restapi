package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Project struct {
	gorm.Model
	Title    string `gorm:"unique" json:"title"`
	Archived bool   `json:"archived"`
}

type Blog struct {
	gorm.Model
	Title       string `gorm:"unique" json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Slug        string `json:"slug"`
}

func (p *Project) Archive() {
	p.Archived = true
}

func (p *Project) Restore() {
	p.Archived = false
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Blog{})
	return db
}
