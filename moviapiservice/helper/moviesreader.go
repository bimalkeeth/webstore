package helper

import (
	"fmt"
	"github.com/micro/go-micro"
	cin "moviapiservice/proto/cinema"
	fil "moviapiservice/proto/filmworld"
	con "golang.org/x/net/context"
)
type MovieReadHelper interface {
	GetCinemaMovie(service micro.Service)[]*cin.MovieSimpleProto
	GetFilmMovie(service micro.Service)[]*fil.MovieSimpleProto
    GetCinemaMovieSingleStream(service micro.Service)[]*cin.MovieSimpleProto
	GetFilmMovieSingleStream(service micro.Service)[]*fil.MovieSimpleProto
}
type MovieRead struct {

}

func (mr *MovieRead)GetCinemaMovie(service micro.Service)[]*cin.MovieSimpleProto{
	cinema :=cin.NewCinamaWorldService("go.micro.srv.cinamaservice",service.Client())
	stream, err :=cinema.GetAllMoviesStream(con.Background(),&cin.RequestAllMovies{})
	if err != nil {
		fmt.Println(err)
	}
	var dataList []*cin.MovieSimpleProto
	rsp, err := stream.Recv()
	dataList=append(dataList,rsp.Movie)
	if err != nil {
		fmt.Println(err)
	}
	for i:=0;i< int(rsp.RecordCount);i++{
		rsp2, err2 := stream.Recv()
		if err2 != nil && (rsp2==nil|| rsp2.Movie==nil) {
			continue
		}
		dataList=append(dataList,rsp2.Movie)
	}
	defer stream.Close()
	return dataList
}
func (mr *MovieRead) GetFilmMovie(service micro.Service)[]*fil.MovieSimpleProto{
	film :=fil.NewFilmWorldService("go.micro.srv.filmworldservice",service.Client())
	stream, err :=film.GetAllMoviesStream(con.Background(),&fil.RequestAllMovies{})
	if err != nil {
		fmt.Println(err)
	}
	var dataList []*fil.MovieSimpleProto
	rsp, err := stream.Recv()
	dataList=append(dataList,rsp.Movie)
	if err != nil {
		fmt.Println(err)
	}
	for i:=0;i< int(rsp.RecordCount);i++{
		rsp2, err2 := stream.Recv()
		if err2 != nil && (rsp2==nil|| rsp2.Movie==nil){
			continue
		}
		dataList=append(dataList,rsp2.Movie)
	}
	defer stream.Close()
	return dataList
}


func (mr *MovieRead)GetCinemaMovieSingleStream(service micro.Service)[]*cin.MovieSimpleProto{
	cinema :=cin.NewCinamaWorldService("go.micro.srv.cinamaservice",service.Client())
	stream, err :=cinema.GetAllMoviesStream(con.Background(),&cin.RequestAllMovies{})
	if err != nil {
		fmt.Println(err)
	}
	var dataList []*cin.MovieSimpleProto
	rsp, err := stream.Recv()
	dataList=append(dataList,rsp.Movie)
	if err != nil {
		fmt.Println(err)
	}
	for i:=0;i< int(rsp.RecordCount);i++{
		rsp2, err2 := stream.Recv()
		if err2 != nil && (rsp2==nil|| rsp2.Movie==nil) {
			continue
		}
		dataList=append(dataList,rsp2.Movie)
	}
	defer stream.Close()
	return dataList
}
func (mr *MovieRead) GetFilmMovieSingleStream(service micro.Service)[]*fil.MovieSimpleProto{
	film :=fil.NewFilmWorldService("go.micro.srv.filmworldservice",service.Client())
	stream, err :=film.GetAllMoviesSingleStream(con.Background(),&fil.RequestAllMovies{})
	if err != nil {
		fmt.Println(err)
	}
	var dataList []*fil.MovieSimpleProto
	rsp, err := stream.Recv()
	dataList=append(dataList,rsp.Movie)
	if err != nil {
		fmt.Println(err)
	}
	for i:=0;i< int(rsp.RecordCount);i++{
		rsp2, err2 := stream.Recv()
		if err2 != nil && (rsp2==nil|| rsp2.Movie==nil){
			continue
		}
		dataList=append(dataList,rsp2.Movie)
	}
	defer stream.Close()
	return dataList
}