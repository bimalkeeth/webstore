package entities

import "github.com/jinzhu/gorm"

type MovieDirectors struct{
	gorm.Model
	Name string `gorm:"column:name;not_null"`
	FirstName string `gorm:"column:firstname"`
	LastName string `gorm:"column:lastname"`
	GenderId uint `gorm:"column:genderid;not_null"`
	Gender MovieGender `gorm:"foreignkey:genderid"`
}

func (c MovieDirectors) TableName() string {
	return "movie_directors"
}