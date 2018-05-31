package data

import "time"

type MovieVM struct{
	Id uint
	Title string
	ReleaseDate time.Time
	MovieType string
	Language string
	Price float64
	Genre string
	Director string
	Actors []MovieActorVM
	Available bool
	Provider string

}
