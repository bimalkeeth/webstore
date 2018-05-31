package entities

import "github.com/jinzhu/gorm"

type MovieGender struct{
	gorm.Model
	GenderType string `gorm:"column:gendertype;not_null"`
    Actors []MovieActors `gorm:"foreignkey:genderid"`
	Directors []MovieDirectors `gorm:"foreignkey:genderid"`
}

func (c MovieGender) TableName() string {
	return "movie_genders"
}