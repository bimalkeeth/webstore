using System;

namespace WebJet.Models
{
    public class MovieVm
    {
      public  int   Id { get; set; }
      public string    Title { get; set; }
      public  string   ReleaseDate { get; set; }
      public  string   MovieType { get; set; }
      public string    Language { get; set; }
      public  decimal   Price { get; set; }
      public string    Genre { get; set; }
      public  string   Director { get; set; }
      public string Provider { get; set; }
       
    }

    public class MoviePVm
    {
        public  int   Id { get; set; }
        public  string    Title { get; set; }
        public  string   ReleaseDate { get; set; }
        public  int   TypeId { get; set; }
        public int    LanguageId { get; set; }
        public  decimal   Price { get; set; }
        public int    GenreId { get; set; }
        public  int   DirectorId { get; set; }
        public MovieActorVM[] Actors { get; set; }
        public  MovieDirectorVM   MovieDirector { get; set; }
        public MovieTypeVM    MovieType { get; set; }
        public  MovieLanguagesVM   Language { get; set; }
        public  MovieGenresVM   Genre { get; set; }
        
    }
    
    public class MovieActorVM {
        
        public  int   Id { get; set; }
        public  string    Name { get; set; }
        public  string   FirstName { get; set; }
        public  string   LastName { get; set; }
        public int    GenderId { get; set; }
        public  string   Gender { get; set; }
 
    }
    
    public class MovieDirectorVM {
        
        public  int   Id { get; set; }
        public  string    Name { get; set; }
        public  string   FirstName { get; set; }
        public  string   LastName { get; set; }
        public int    GenderId { get; set; }
        public  string   Gender { get; set; }
 
    }
    public class MovieGenresVM {
        
        public  int   Id { get; set; }
        public  string    Genre { get; set; }
        
 
    }
    public class MovieLanguagesVM {
        
        public  int   Id { get; set; }
        public  string    Language { get; set; }
        
 
    }
    public class MovieTypeVM {
        
        public  int   Id { get; set; }
        public  string    Name { get; set; }
        
 
    }
    


}

