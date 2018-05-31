package dataschema

import (
	  "filmworldservice/dataservice/sqlservice"
	  "log"
 com  "filmworldservice/common"
 ent  "filmworldservice/dataservice/sqlservice/entities"
	  "github.com/jinzhu/gorm"
	  "math/rand"
	  "time"
	  "math"
 mock "filmworldservice/common/mockdata"
)

type Schema interface {
   GenerateSchema()
   CreateSchema(db *gorm.DB)(bool)
   MigrateData(db *gorm.DB)(bool)
}
type GeneratedSchema bool

func (g *GeneratedSchema)GenerateSchema() bool {
	db,err,config:=sqlservice.DbConn()
	if err != nil {
		log.Fatal(com.GetErrorDescription("ERR001"))
	}
	if config.Create {
		CreateSchema(db)
	}
	if config.Migrate {
		MigrateData(db)
	}
	return true
}
func CreateSchema(db *gorm.DB)(bool) {

	db.SingularTable(true)
	db.DropTableIfExists(&ent.ActorWithMovies{})
	db.DropTableIfExists(&ent.MovieMovies{})
	db.DropTableIfExists(&ent.MovieActors{})
	db.DropTableIfExists(&ent.MovieLanguages{})
	db.DropTableIfExists(&ent.MovieType{})
	db.DropTableIfExists(&ent.MovieGenres{})
	db.DropTableIfExists(&ent.MovieGender{})
	db.DropTableIfExists(&ent.MovieDirectors{})

	if !db.HasTable(&ent.MovieType{}) {
		db.CreateTable(&ent.MovieType{})
		db.Model(&ent.MovieType{}).AddUniqueIndex("idx_movietypes_name", "name")
	}
	if !db.HasTable(&ent.MovieGenres{}) {
		db.CreateTable(&ent.MovieGenres{})
		db.Model(&ent.MovieGenres{}).AddUniqueIndex("idx_moviegeneres_genre", "genre")
	}
	if !db.HasTable(&ent.MovieGender{}) {
		db.CreateTable(&ent.MovieGender{})
		db.Model(&ent.MovieGender{}).AddUniqueIndex("idx_moviegender_gendertype", "gendertype")
	}
	if !db.HasTable(&ent.MovieDirectors{}) {
		db.CreateTable(&ent.MovieDirectors{})
		db.Model(&ent.MovieDirectors{}).AddUniqueIndex("idx_moviedirectors_name", "name")
	}
	if !db.HasTable(&ent.MovieLanguages{}) {
		db.CreateTable(&ent.MovieLanguages{})
		db.Model(&ent.MovieLanguages{}).AddUniqueIndex("idx_movie_languages_language", "language")
	}
	if !db.HasTable(&ent.MovieActors{}) {
		db.CreateTable(&ent.MovieActors{})
		db.Model(&ent.MovieActors{}).AddUniqueIndex("idx_movieactors_name", "name")
		db.Model(&ent.MovieActors{}).AddForeignKey("genderid", "movie_genders(id)", "RESTRICT", "RESTRICT")
	}
	if !db.HasTable(&ent.MovieMovies{}) {
		db.CreateTable(&ent.MovieMovies{})
		db.Model(&ent.MovieMovies{}).AddForeignKey("directorid", "movie_directors(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.MovieMovies{}).AddForeignKey("typeid", "movie_types(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.MovieMovies{}).AddForeignKey("languageid", "movie_languages(id)", "RESTRICT", "RESTRICT")
		db.Model(&ent.MovieMovies{}).AddForeignKey("genreid", "movie_generes(id)", "RESTRICT", "RESTRICT")
	}
	if !db.HasTable(&ent.ActorWithMovies{}) {
		db.CreateTable(&ent.ActorWithMovies{})
		db.Model(&ent.ActorWithMovies{}).AddForeignKey("actorid", "movie_actors(id)", "CASCADE", "RESTRICT")
		db.Model(&ent.ActorWithMovies{}).AddForeignKey("movieid", "movie_movies(id)", "CASCADE", "RESTRICT")
	}
	return true
}
func MigrateData(db *gorm.DB)(bool) {

	createMovieTypes(db)
	createMovieGender(db)
	createMovieGenre(db)
	createMovieLanguages(db)
	createMovieActors(db)
	createMovieDirectors(db)
	createMovies(db)
	createMoviesWithActors(db)
	return true
}

func createMovieTypes(db *gorm.DB){

	db.Create( &ent.MovieType{Name:"Classic"})
	db.Create( &ent.MovieType{Name:"Commercial"})
}
func createMovieGender(db *gorm.DB){

	db.Create( &ent.MovieGender{GenderType:"Male"})
	db.Create( &ent.MovieGender{GenderType:"Female"})
}

func createMovieGenre(db *gorm.DB){

	db.Create( &ent.MovieGenres{Genre:"Action"})
	db.Create( &ent.MovieGenres{Genre:"Adventure"})
	db.Create( &ent.MovieGenres{Genre:"Comedy"})
	db.Create( &ent.MovieGenres{Genre:"Crime"})
	db.Create( &ent.MovieGenres{Genre:"Drama"})
	db.Create( &ent.MovieGenres{Genre:"Fantasy"})
	db.Create( &ent.MovieGenres{Genre:"Historical"})
	db.Create( &ent.MovieGenres{Genre:"Horror"})
	db.Create( &ent.MovieGenres{Genre:"Magical realism"})
	db.Create( &ent.MovieGenres{Genre:"Mystery"})
	db.Create( &ent.MovieGenres{Genre:"Paranoid Fiction"})
	db.Create( &ent.MovieGenres{Genre:"Philosophical"})
	db.Create( &ent.MovieGenres{Genre:"Political"})
	db.Create( &ent.MovieGenres{Genre:"Romance"})
	db.Create( &ent.MovieGenres{Genre:"Saga"})
	db.Create( &ent.MovieGenres{Genre:"Satire"})
	db.Create( &ent.MovieGenres{Genre:"Science fiction"})
	db.Create( &ent.MovieGenres{Genre:"Social"})
	db.Create( &ent.MovieGenres{Genre:"Speculative"})
	db.Create( &ent.MovieGenres{Genre:"Thriller"})
}
func createMovieLanguages(db *gorm.DB){

	db.Create( &ent.MovieLanguages{Language:"English"})
	db.Create( &ent.MovieLanguages{Language:"German"})
	db.Create( &ent.MovieLanguages{Language:"Russian"})
	db.Create( &ent.MovieLanguages{Language:"French"})
	db.Create( &ent.MovieLanguages{Language:"Hindi"})
}
func createMovieActors(db *gorm.DB){

	var femalegender=ent.MovieGender{}
	db.Where("gendertype = ?", "Female").First(&femalegender)

	db.Create( &ent.MovieActors{Name:"Beverly Aadland",FirstName:"Beverly",LastName:"Aadland",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Christina Aguilera",FirstName:"Christina",LastName:"Aguilera",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Rowan Blanchard",FirstName:"Rowan",LastName:"Blanchard",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Alexis Bledel",FirstName:"Alexis",LastName:"Bledel",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Moon Bloodgood ",FirstName:"Moon",LastName:"Bloodgood",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Julie Benz",FirstName:"Julie",LastName:"Benz",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Elizabeth Berkley",FirstName:"Elizabeth",LastName:"Berkley",GenderId:femalegender.ID})
	db.Create( &ent.MovieActors{Name:"Jessica Biel",FirstName:"Jessica",LastName:"Biel",GenderId:femalegender.ID})

	var malegender=ent.MovieGender{}
	db.Where("gendertype = ?", "Male").First(&malegender)

	db.Create( &ent.MovieActors{Name:" Robert De Niro",FirstName:"Robert",LastName:"De Niro",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:" Jack Nicholson",FirstName:"Jack",LastName:"Nicholson",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:"Leonardo DiCaprio",FirstName:"Leonardo",LastName:"DiCaprio",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:"Johnny Depp",FirstName:"Johnny",LastName:"Depp",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:"Brad Pitt",FirstName:"Brad",LastName:"Pitt",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:"Sean Penn",FirstName:"Sean",LastName:"Penn",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:" Matt Damon",FirstName:"Matt",LastName:"Damon",GenderId:malegender.ID})
	db.Create( &ent.MovieActors{Name:"Will Smith",FirstName:"Will",LastName:"Smith",GenderId:malegender.ID})

}

func createMovieDirectors(db *gorm.DB){

	var femalegender=ent.MovieGender{}
	db.Where("gendertype = ?", "Female").First(&femalegender)

	db.Create( &ent.MovieDirectors{Name:"Lilly Wachowski",FirstName:"Lilly",LastName:"Wachowski",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Jennifer Abbott",FirstName:"Jennifer",LastName:"Abbott",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Maren Ade",FirstName:"Maren",LastName:"Ade",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Elizabeth Allen",FirstName:"Elizabeth",LastName:"Allen",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Kathy Bates",FirstName:"Kathy",LastName:"Bates",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Sadie Benning",FirstName:"Sadie",LastName:"Benning",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Cindy Baer",FirstName:"Cindy",LastName:"Baer",GenderId:femalegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Manon Barbeau",FirstName:"Manon",LastName:"Barbeau",GenderId:femalegender.ID})

	var malegender=ent.MovieGender{}
	db.Where("gendertype = ?", "Male").First(&malegender)

	db.Create( &ent.MovieDirectors{Name:"Steven Spielberg",FirstName:"Steven",LastName:"Spielberg",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Martin Scorsese",FirstName:"Martin",LastName:"Scorsese",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Ridley Scott",FirstName:"Ridley",LastName:"Scott",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Christopher Nolan",FirstName:"Christopher",LastName:"Nolan",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Peter Jackson",FirstName:"Peter",LastName:"Jackson",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"James Cameron",FirstName:"James",LastName:"Cameron",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Robert Zemeckis",FirstName:"Robert",LastName:"Zemeckis",GenderId:malegender.ID})
	db.Create( &ent.MovieDirectors{Name:"Oliver Stone",FirstName:"Oliver",LastName:"Stone",GenderId:malegender.ID})
}
func randBool(min int, max int) bool{
	return uint(rand.Intn(max-min) + min) % 2 == 0
}
func createMovies(db *gorm.DB){

	rand.Seed(time.Now().UnixNano())
	for  _,moviename :=range mock.MoviesNames{

		randomdirectors := random(1, 16)
		randomgener := random(1, 20)
		randomtype := random(1, 2)
		db.Debug().Create( &ent.MovieMovies{Title:moviename,LanguageId:1,DirectorId:randomdirectors,GenreId:randomgener,TypeId:randomtype, ReleaseDate:randate(), Price:randFloat(1.01,650.99),Available:randBool(1,999)})
	}
}

func random(min int, max int) uint {
	return uint(rand.Intn(max-min) + min)
}
func randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2018, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return  time.Unix(sec, 0)
}
func randFloat(min,max float64)float64{
	return math.Round(min + rand.Float64() * (max - min)*100)/100
}

func createMoviesWithActors(db *gorm.DB){
	rows, err := db.Model(&ent.MovieMovies{}).Select("id").Rows()
	defer rows.Close()
	if err !=nil{
		log.Fatal(com.GetErrorDescription("ERR002"))
	}
	rand.Seed(time.Now().UnixNano())
	for rows.Next(){
		var movieid uint
		rows.Scan(&movieid)
        for i:=1; i<16;i++{
			randomdirectors := random(1, 16)
			db.Debug().Create( &ent.ActorWithMovies{ActorId:randomdirectors,MovieId:movieid})
		}
	}
}