package main
import (
	api "moviapiservice/manageservice/apiroutemanager"
	"net/http"
	"log"
	"fmt"
	"time"
	"github.com/gorilla/handlers"
)
func main() {

	fmt.Println("Movie Api server is now on ",time.Now())

	movieRouter:=new(api.RouteInstance)
	movieRouter.RegisterRoutes()
	movieRouter.MovieApiRegister()

	log.Fatal(http.ListenAndServe(":8081",handlers.CORS()( movieRouter.MovieRouter)))
}
