package apiroutemanager

import "github.com/gorilla/mux"
type RouterRegistry interface {
	RegisterRoutes()
	MovieApiRegister()
}
type RouteInstance struct {
	MovieRouter *mux.Router
}
func (r *RouteInstance)RegisterRoutes(){
	r.MovieRouter=mux.NewRouter().StrictSlash(true)
}




