package common

import (
	"path/filepath"
	"fmt"
	"strings"
)
func GetConfigFilePath()string{
	return GetExecutionPath()+ "/accessconfig.json"
}
func GetErrorFilePath()string{
	return GetExecutionPath()+ "/errorcodes.json"
}
func GetExecutionPath() string {
	path,err:=filepath.Abs("./")
	if err!=nil{
		fmt.Println("Path Error")
	}
	if strings.Contains(path,"filmworldservice"){
		return path
	}
	return path+"/filmworldservice"
}