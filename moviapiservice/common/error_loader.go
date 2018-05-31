package common

import (
	conf "moviapiservice/common/contacts/configuration"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

var MessageData []conf.ErrorCatalog

//--------------------------------------------------
//Unmarshal error json document
//--------------------------------------------------
func  GetErrors() (error) {
	if MessageData!=nil{
		return nil
	}
	configData :=GetErrorFilePath()

	raw, err := ioutil.ReadFile(configData)
	if err != nil {
		fmt.Println(err.Error())
	}
	var errorList []conf.ErrorCatalog
	json.Unmarshal(raw, &errorList)
	MessageData=errorList
	return nil
}
//--------------------------------------------------
//Get errors into list of error code
//--------------------------------------------------
func GetErrorDescription(errorCode string)(string){
	GetErrors()
	for _,data:=range MessageData{
		if data.ErrorCode==errorCode{
			return data.ErrorDescription
		}
	}
	return ""
}