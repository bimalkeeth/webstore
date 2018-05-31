package data

import "time"

type MovieVM struct{
	Id uint
	Title string
	ReleaseDate time.Time
	TypeId uint
	LanguageId uint
	Price float64
	GenreId uint
	DirectorId uint
	Actors []MovieActorVM
	Director MovieDirectorVM
	MoveType MovieTypeVM
	Language MovieLanguagesVM
	Genere MovieGenresVM
	Available bool
}
