package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)
type MovieMovies struct{
	gorm.Model
	Title string `gorm:"column:title;not_null"`
	ReleaseDate time.Time `gorm:"column:releasedate;not_null"`
    TypeId uint `gorm:"column:typeid;not_null"`
    LanguageId uint `gorm:"column:languageid;not_null"`
    Price float64 `gorm:"column:price;not_null"`
	GenreId uint `gorm:"column:genreid;not_null"`
	DirectorId uint `gorm:"column:directorid;not_null"`
	Actors []ActorWithMovies `gorm:"foreignkey:movieid"`
	Director MovieDirectors `gorm:"foreignkey:directorid"`
	MoveType MovieType `gorm:"foreignkey:typeid"`
	Language MovieLanguages `gorm:"foreignkey:languageid"`
    Genere MovieGenres `gorm:"foreignkey:genreid"`
    Available bool `gorm:"column:available;not_null;default:false"`
}
func (c MovieMovies) TableName() string {
	return "movie_movies"
}