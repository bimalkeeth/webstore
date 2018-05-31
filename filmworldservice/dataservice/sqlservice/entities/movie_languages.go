package entities

import "github.com/jinzhu/gorm"

type MovieLanguages struct{
	gorm.Model
	Language string `gorm:"column:language;not_null"`
	Movies []MovieMovies `gorm:"foreignkey:languageid"`

}
func (c MovieLanguages) TableName() string {
	return "movie_languages"
}