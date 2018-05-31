package apiroutemanager

func (r *RouteInstance )MovieApiRegister(){
	//r.MovieRouter.HandleFunc("/api/cinemaworld/movies",tkn.ValidateAccess(GetAllCinemaMovies)).Methods("GET")
	//r.MovieRouter.HandleFunc("/api/cinemaworld/movie/{id}", tkn.ValidateAccess(GetCinenmaMovie)).Methods("GET")
	//r.MovieRouter.HandleFunc("/api/filmworld/movies",tkn.ValidateAccess(GetAllFilmWorldMovies)).Methods("GET")
	//r.MovieRouter.HandleFunc("/api/filmworld/movie/{id}", tkn.ValidateAccess(GetFilmWorldMovie)).Methods("GET")
	//r.MovieRouter.HandleFunc("/api/movies/", tkn.ValidateAccess(GetAllMovies)).Methods("GET")

	r.MovieRouter.HandleFunc("/api/cinemaworld/movies",GetAllCinemaMovies).Methods("GET")
	r.MovieRouter.HandleFunc("/api/cinemaworld/movie/{id}", GetCinenmaMovie).Methods("GET")
	r.MovieRouter.HandleFunc("/api/filmworld/movies",GetAllFilmWorldMovies).Methods("GET")
	r.MovieRouter.HandleFunc("/api/filmworld/movie/{id}", GetFilmWorldMovie).Methods("GET")
	r.MovieRouter.HandleFunc("/api/movies/", GetAllMovies).Methods("GET")
}

