package apiroutemanager

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	 com "moviapiservice/common"
	"strconv"
	"log"
	"github.com/micro/go-micro"
     cin "moviapiservice/proto/cinema"
     fil "moviapiservice/proto/filmworld"
	 con "golang.org/x/net/context"
	 dat "moviapiservice/common/contacts/data"
	 hel "moviapiservice/helper"
)

var service micro.Service
func init(){
	service = micro.NewService(micro.Name("cinamaclient.client"))
	service.Init()
	err:=com.GetErrors()
	if err !=nil{
		log.Fatal(com.GetErrorDescription("ERR004"))
	}
}
//-------------------------------------------------------
//Get All Movies order by Price
//-------------------------------------------------------
func GetAllCinemaMovies(w http.ResponseWriter, r *http.Request) {

	var reader hel.MovieReadHelper= &hel.MovieRead{}
	cinemaResult:=reader.GetCinemaMovieSingleStream(service)
	var dataList []dat.MovieVM
	for _,filitem :=range cinemaResult{
		var item dat.MovieVM
			item= dat.MovieVM{
				Price:float64(filitem.Price),
				Id:uint(filitem.Id),
				Title:filitem.Title,
				Director:filitem.Director,
				MovieType:filitem.MovieType,
				Genre:filitem.Genre,
				Language:filitem.Language,
				Provider:filitem.Provider,
			}
			dataList=append(dataList,item)
		}
	json.NewEncoder(w).Encode(dataList)
}


//------------------------------------------------------
//Get Movie by Id
//------------------------------------------------------
func GetCinenmaMovie(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(com.GetErrorDescription("ERR003"))
	}
	cinema :=cin.NewCinamaWorldService("go.micro.srv.cinamaservice",service.Client())
	rsp, err :=cinema.GetMovieByMovieId(con.Background(),&cin.RequestMovie{Id:int32(id)})

	if err != nil{
		log.Fatal(com.GetErrorDescription("ERR004"))
	}
	json.NewEncoder(w).Encode(rsp.Movie)
}


//-------------------------------------------------------
//Get All Movies order by Price
//-------------------------------------------------------
func GetAllFilmWorldMovies(w http.ResponseWriter, r *http.Request) {

	var reader hel.MovieReadHelper= &hel.MovieRead{}
	cinemaResult:=reader.GetFilmMovieSingleStream(service)
	var dataList []dat.MovieVM
	for _,filitem :=range cinemaResult{
		var item dat.MovieVM
		item= dat.MovieVM{
			Price:float64(filitem.Price),
			Id:uint(filitem.Id),
			Title:filitem.Title,
			Director:filitem.Director,
			MovieType:filitem.MovieType,
			Genre:filitem.Genre,
			Language:filitem.Language,
			Provider:filitem.Provider,
		}
		dataList=append(dataList,item)
	}
	json.NewEncoder(w).Encode(dataList)
}
//------------------------------------------------------
//Get Movie by Id
//------------------------------------------------------
func GetFilmWorldMovie(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(com.GetErrorDescription("ERR003"))
	}
	film :=fil.NewFilmWorldService("go.micro.srv.filmworldservice",service.Client())
	rsp, err :=film.GetMovieByMovieId(con.Background(),&fil.RequestMovie{Id:int32(id)})
	if err != nil{
		log.Fatal(com.GetErrorDescription("ERR004"))
	}
	json.NewEncoder(w).Encode(rsp.Movie)
}
//-----------------------------------------------------
//Get movies with less price
//-----------------------------------------------------
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
  var reader hel.MovieReadHelper= &hel.MovieRead{}

	cinemaResult:=reader.GetCinemaMovie(service)
	filmResult:=reader.GetFilmMovie(service)
	var dataList []dat.MovieVM

	for _,elm := range cinemaResult{
		for _,filitem :=range filmResult{
			var item dat.MovieVM
			if filitem.Title == elm.Title && filitem.Price < elm.Price{

				item= dat.MovieVM{
					Price:float64(filitem.Price),
					Id:uint(filitem.Id),
					Title:filitem.Title,
					Director:filitem.Director,
					MovieType:filitem.MovieType,
					Genre:filitem.Genre,
					Language:filitem.Language,
					Provider:filitem.Provider,

				}
				dataList=append(dataList,item)
			}else if filitem.Title == elm.Title && filitem.Price > elm.Price{

				item= dat.MovieVM{
					Price:float64(elm.Price),
					Id:uint(elm.Id),
					Title:elm.Title,
					Director:elm.Director,
					MovieType:elm.MovieType,
					Genre:elm.Genre,
					Language:elm.Language,
					Provider:filitem.Provider,
				}
				dataList=append(dataList,item)
			}
		}
     }
	json.NewEncoder(w).Encode(dataList)
}
