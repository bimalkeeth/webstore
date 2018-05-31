package entities

import "github.com/jinzhu/gorm"

type MovieGenres struct{
	gorm.Model
	Genre string `gorm:"column:genre;not_null"`
	Movies []MovieMovies `gorm:"foreignkey:genreid"`

}
func (c MovieGenres) TableName() string {
	return "movie_generes"
}