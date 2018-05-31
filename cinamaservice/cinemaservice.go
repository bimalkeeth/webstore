package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"cinamaservice/servicemanager"
cin "cinamaservice/proto/cinema"
boot"cinamaservice/servicemanager/dbbootstrap"
enu "cinamaservice/common/enums")
func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.cinamaservice"),
		micro.Version("latest"),
	)
	// Register Handler
	cin.RegisterCinamaWorldServiceHandler(service.Server(), new(servicemanager.CinemaRequest))
	// Initialise service
	service.Init()
	boost:=new(boot.BoosTapper)
	boost.InvokeDatabaseInitialize(enu.SqlConnection)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
