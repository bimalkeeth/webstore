package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"filmworldservice/servicemanager"
cin "filmworldservice/proto/filmworld"
boot"filmworldservice/servicemanager/dbbootstrap"
enu "filmworldservice/common/enums"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.filmworldservice"),
		micro.Version("latest"),
	)
	// Register Handler
	cin.RegisterFilmWorldServiceHandler(service.Server(), new(servicemanager.FilmWorldRequest))
	// Initialise service
	service.Init()
	boost:=new(boot.BoosTapper)
	boost.InvokeDatabaseInitialize(enu.SqlConnection)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
