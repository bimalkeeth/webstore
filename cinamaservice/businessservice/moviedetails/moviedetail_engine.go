package moviedetails

import (
  dat "cinamaservice/common/contacts/data"
      "cinamaservice/common/enums"
      "cinamaservice/dataservice/sqlservice"
  com "cinamaservice/common"
      "log"
  ent "cinamaservice/dataservice/sqlservice/entities"
      "fmt"
      "github.com/jinzhu/gorm"
  conf "cinamaservice/common/contacts/configuration"

)
var config conf.Configuration

type MovieDetails interface {
  GetAllMoviesAvailable()([]dat.MovieVM,error)
  GetMovieByMovieId(id uint)(dat.MovieVM,error)
}
type MovieService struct {}

func init(){
    config= com.GetDbConfig()
}

func(m *MovieService)GetMovieByMovieId(id uint)(dat.MovieVM,error){

    db,err,_:=sqlservice.DbConn()
    if err !=nil {
        log.Fatal(com.GetErrorDescription("ERR003"))
    }
    moviedata:=ent.MovieMovies{}
     err=db.Preload("Language").
            Preload("Director").
            Preload("MoveType").
            Preload("Genere").
            Preload("Actors.Actor.Gender").
            Preload("Actors.Actor").
            Preload("Actors").First(&moviedata, id).Error
              if err != nil{
                  log.Fatal(com.GetErrorDescription("ERR004"))
              }
     return composeMovieWith(moviedata),nil
}
//-------------------------------------------------
//Get All the movies based on connection type
//-------------------------------------------------
func(m *MovieService) GetAllMoviesAvailable()([]dat.MovieVM,error){
      if enums.DbOperation(config.ConnType)==enums.SqlConnection{
        return getDetailFromDatabase()
      }
      return nil, nil
}
//-------------------------------------------------
//Get all the Movies
//-------------------------------------------------
func getDetailFromDatabase()([]dat.MovieVM,error){

    db,err,_:=sqlservice.DbConn()
      if err !=nil {
        log.Fatal(com.GetErrorDescription("ERR001"))
      }
    var movies []ent.MovieMovies
    var movievm []dat.MovieVM
   err=db.Preload("Language").
          Preload("Director").
          Preload("MoveType").
          Preload("Genere").
          Preload("Actors").
          Preload("Actors.Actor").
          Order("price asc").Find(&movies).Error
        if err != nil{
            log.Fatal(com.GetErrorDescription("ERR004"))
        }
        for i := range movies {
            results := GetMovie(movies[i],db)
            for ch :=range results{
                movievm=append(movievm,ch)
            }
            fmt.Println(movies[i].Language.Language)
        }
        return movievm,nil
}
//-------------------------------------------------------------------
//Set up related actors for the movie
//-------------------------------------------------------------------
func GetMovie(moviefrom ent.MovieMovies,db *gorm.DB)<-chan dat.MovieVM{
    chanMovie:= make(chan	dat.MovieVM)
    go func(moviefromin ent.MovieMovies,dbin *gorm.DB){
         for _,act:=range moviefromin.Actors{
            actor:=ent.MovieActors{}
            dbin.Model(act).Association("Actor").Find(&actor)
            act.Actor=actor
            gender:=ent.MovieGender{}
            dbin.Model(actor).Association("Gender").Find(&gender)

        }
        chanMovie <- composeMovieWith(moviefromin)
        close(chanMovie)
    }(moviefrom,db)
    return chanMovie
}
//----------------------------------------------------------------
//Compose the Move model into VM Model
//----------------------------------------------------------------
func composeMovieWith(moviefrom ent.MovieMovies) dat.MovieVM{
    movitem:=dat.MovieVM{}

    movitem.Id=moviefrom.ID
    movitem.Title=moviefrom.Title
    movitem.TypeId=moviefrom.TypeId
    movitem.Price=moviefrom.Price
    movitem.ReleaseDate=moviefrom.ReleaseDate
    movitem.DirectorId=moviefrom.DirectorId
    movitem.LanguageId=moviefrom.LanguageId
    movitem.GenreId=moviefrom.GenreId

    movitem.Genere=dat.MovieGenresVM{Id:moviefrom.Genere.ID,Genre:moviefrom.Genere.Genre}

    movitem.Director=dat.MovieDirectorVM{Id:moviefrom.Director.ID,
                                            Name:moviefrom.Director.Name,
                                            LastName:moviefrom.Director.LastName,
                                            FirstName:moviefrom.Director.FirstName,
                                            GenderId:moviefrom.Director.GenderId}

    movitem.Language=dat.MovieLanguagesVM{Id:moviefrom.Language.ID,Language:moviefrom.Language.Language}
    movitem.MoveType=dat.MovieTypeVM{Id:moviefrom.MoveType.ID,Name:moviefrom.MoveType.Name}
    movitem.Genere=dat.MovieGenresVM{Id:moviefrom.ID,Genre:moviefrom.Genere.Genre}

    var actorsvm []dat.MovieActorVM
    for _,act:=range moviefrom.Actors{
        actorsvm=append(actorsvm,dat.MovieActorVM{Id:act.Actor.ID,Name:act.Actor.Name,FirstName:act.Actor.FirstName,LastName:act.Actor.LastName,GenderId:act.Actor.GenderId,Gender:act.Actor.Gender.GenderType})
    }
    movitem.Actors=actorsvm
    return movitem
}