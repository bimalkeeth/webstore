package servicemanager

import (
	"context"
	"github.com/micro/go-log"
	cin "cinamaservice/proto/cinema"
	bus "cinamaservice/businessservice/moviedetails"
	com "cinamaservice/common"
	pro "cinamaservice/proto/cinema"
	dat "cinamaservice/common/contacts/data"
)
type CinemaRequest struct{}
func (e *CinemaRequest) GetAllMoviesSingleStream(ctx context.Context, req *cin.RequestAllMovies, rsp cin.CinamaWorldService_GetAllMoviesSingleStreamStream) error {
	log.Log("Received Example.Call request")
	movieService:=new(bus.MovieService)
	dataList,err:=movieService.GetAllMoviesAvailable()
	if err!=nil{
		log.Fatal(com.GetErrorDescription("ERR004"))
	}
	var recordCount=len(dataList)
	for _,mov:= range dataList{

		movie:=new(pro.MovieSimpleProto)
		movie.Id=int32(mov.Id)
		movie.Genre=mov.Genere.Genre
		movie.Language=mov.Language.Language
		movie.MovieType=mov.MoveType.Name
		movie.Director=mov.Director.Name
		movie.ReleaseDate=mov.ReleaseDate.String()
		movie.Price=float32(mov.Price)
		movie.Title=mov.Title
		movie.Provider="cinemaworld"
		var actorList =setMovieActorData(mov.Actors)
		movie.Actors=actorList
		err:=rsp.Send(&pro.ResponseAllMovies{Movie:movie,RecordCount:int32(recordCount)})
		if err!=nil{
			log.Fatal(com.GetErrorDescription("ERR004"))
		}
	}

	return nil
}

func setMovieActorData(actors []dat.MovieActorVM) []*pro.ActorProto{
	var actorList []*pro.ActorProto
	for _,actor:=range actors{

		movActor:= new(pro.ActorProto)
		movActor.Id=int32(actor.Id)
		movActor.FirstName=actor.FirstName
		movActor.LastName=actor.LastName
		movActor.Gender=actor.Gender
		movActor.Name=actor.Name
		actorList=append(actorList,movActor)
	}
	return actorList
}


func (e *CinemaRequest) GetMovieByMovieId(ctx context.Context, req *cin.RequestMovie, rsp *cin.ResponseMovie) error {
	log.Log("Received Example.Call request")
	movieService:=new(bus.MovieService)
	dataMovie,err:=movieService.GetMovieByMovieId(uint(req.Id))
	if err!=nil{
		log.Fatal(com.GetErrorDescription("ERR004"))
	}
	movie:=new(pro.MovieProto)
	movie.Id=int32(dataMovie.Id)
	movie.Title=dataMovie.Title
	movie.Price=float32(dataMovie.Price)
	movie.LanguageId=int32(dataMovie.LanguageId)
	movie.Language=&pro.MovieLanguagesProto{Language:dataMovie.Language.Language,Id:int32(dataMovie.LanguageId)}
	movie.ReleaseDate=dataMovie.ReleaseDate.String()
	movie.Genre=&pro.MovieGenresProto{Id:int32(dataMovie.Genere.Id),Genre:dataMovie.Genere.Genre}
	movie.MovieType=&pro.MovieTypeProto{Id:int32(dataMovie.TypeId),Name:dataMovie.MoveType.Name}
	movie.MovieDirector=&pro.DirectorProto{Name:dataMovie.Director.Name,Id:int32(dataMovie.Director.Id),FirstName:dataMovie.Director.FirstName,LastName:dataMovie.Director.LastName}
	movie.Actors=setMovieActorData(dataMovie.Actors)
	rsp.Movie=movie
	return nil
}

func (e *CinemaRequest) GetAllMoviesStream(ctx context.Context, req *cin.RequestAllMovies, stream cin.CinamaWorldService_GetAllMoviesStreamStream) error {

	movieService:=new(bus.MovieService)
	dataList,err:=movieService.GetAllMoviesAvailable()
	if err!=nil{
		log.Fatal(com.GetErrorDescription("ERR004"))
	}
    var recordCount=len(dataList)
	for _,mov:= range dataList{

		movie:=new(pro.MovieSimpleProto)
		movie.Id=int32(mov.Id)
		movie.Genre=mov.Genere.Genre
		movie.Language=mov.Language.Language
		movie.MovieType=mov.MoveType.Name
		movie.Director=mov.Director.Name
		movie.ReleaseDate=mov.ReleaseDate.String()
		movie.Price=float32(mov.Price)
		movie.Title=mov.Title
		movie.Provider="cinemaworld"
		var actorList =setMovieActorData(mov.Actors)
		movie.Actors=actorList

		err:=stream.Send(&pro.ResponseMovieStream{Movie:movie,RecordCount:int32(recordCount)})
		if err!=nil{
			log.Fatal(com.GetErrorDescription("ERR004"))
		}
	}
	return nil
}