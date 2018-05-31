package entities

import "github.com/jinzhu/gorm"

type MovieActors struct{
	gorm.Model
	Name string `gorm:"column:name;not_null"`
	GenderId uint `gorm:"column:genderid;not_null"`
	FirstName string `gorm:"column:firstname"`
	LastName string `gorm:"column:lastname"`
	Gender MovieGender `gorm:"foreignkey:genderid"`
}
func (c MovieActors) TableName() string {
	return "movie_actors"
}