package entities

import "github.com/jinzhu/gorm"

type MovieType struct{
	gorm.Model
	Name string `gorm:"column:name;not_null"`
    Movies []MovieMovies `gorm:"foreignkey:typeid"`
}
func (c MovieType) TableName() string {
	return "movie_types"
}