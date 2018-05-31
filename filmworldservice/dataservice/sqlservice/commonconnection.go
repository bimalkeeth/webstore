package sqlservice

import (
 conf "filmworldservice/common/contacts/configuration"
      "filmworldservice/common"
     _"github.com/jinzhu/gorm/dialects/postgres"
	 _"github.com/jinzhu/gorm/dialects/mysql"
	 _"github.com/jinzhu/gorm/dialects/sqlite"
      "github.com/jinzhu/gorm"
	  "fmt"
	  "log"
	  dbt "filmworldservice/common/enums"

 com  "filmworldservice/common"
)
//--------------------------------------------
// Get DbConfig Data
//--------------------------------------------
func  getDatabaseConfig() conf.Configuration {
	configData := common.GetDbConfig()
	return configData
}
//----------------------------------------------
// Setup database connection
//----------------------------------------------
var db *gorm.DB
var selectedDb dbt.DbTypes
func DbConn() (*gorm.DB,error,conf.Configuration) {

	config:=getDatabaseConfig()
	connectionStr,providerType:=getConnectionString(config)
	if db!=nil && selectedDb==providerType{
		return db,nil,config
	}
	db, err := gorm.Open(config.Provider, connectionStr)
	if err!=nil{
		log.Fatal( com.GetErrorDescription("ERR001"),err)
		return nil,err,config
	}
	return db,nil,config
}
//----------------------------------------------
// setup Connection string
//----------------------------------------------
func getConnectionString(config conf.Configuration) (string,dbt.DbTypes) {
	var connectionString =""
	var dbProvider dbt.DbTypes
	switch config.Provider {
	    case "postgres":
			connectionString= fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", config.HostName,config.Port, config.User,config.Database,config.Password)
			dbProvider=dbt.Postgress
			createSchemaIfNotExists(config)
	     break
	    case "mysql":
			connectionString=  fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",config.User,config.Password,config.HostName,config.Port,config.Database)
			dbProvider=dbt.MySql
			createSchemaIfNotExists(config)
		 break
	    case "sqllite" :
			connectionString= com.GetExecutionPath()+"movie.db"
			dbProvider=dbt.SqlLite
	default:
		connectionString= ""
	}
	return connectionString,dbProvider
}
//-----------------------------------------------
//Create Schema if not exists
//-----------------------------------------------
func createSchemaIfNotExists(config conf.Configuration){
	var connectionString =""
	switch config.Provider {
	case "postgres":
		connectionString= fmt.Sprintf("host=%s port=%s user=%s password=%s", config.HostName,config.Port, config.User,config.Password)
		dbPostGre, err := gorm.Open(config.Provider, connectionString)
		if err!=nil{
			log.Fatal( com.GetErrorDescription("ERR001"),err)
		}
		defer dbPostGre.Close()
		dbPostGre.Exec("CREATE DATABASE IF NOT EXISTS "+config.Database)
		break
	case "mysql":
		connectionString=  fmt.Sprintf("%s:%s@tcp(%s:%s)/",config.User,config.Password,config.HostName,config.Port)
		dbMySql, err := gorm.Open(config.Provider, connectionString)
		if err!=nil{
			log.Fatal( com.GetErrorDescription("ERR001"),err)
		}
		defer dbMySql.Close()
		dbMySql.Exec("CREATE DATABASE IF NOT EXISTS "+config.Database)
		break

	default:
		connectionString= ""
	}

}