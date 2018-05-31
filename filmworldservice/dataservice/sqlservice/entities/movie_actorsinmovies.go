package entities

import "github.com/jinzhu/gorm"

type ActorWithMovies struct{
	gorm.Model
	ActorId uint `gorm:"column:actorid;not_null"`
	MovieId uint `gorm:"column:movieid;not_null"`
    Movie MovieMovies `gorm:"foreignkey:movieid"`
    Actor MovieActors `gorm:"foreignkey:actorid"`
}
func (c ActorWithMovies) TableName() string {
	return "movie_actorsinmovies"
}