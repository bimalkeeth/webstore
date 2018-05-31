package common

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
	conf "filmworldservice/common/contacts/configuration"
)
//-------------------------------------------------
//Database configuration
//-------------------------------------------------
func  GetDbConfig() (conf.Configuration) {
	configDataVal :=GetConfigFilePath()
	raw, err := ioutil.ReadFile(configDataVal)
	if err != nil {
		fmt.Println(err.Error())
	}
	var dbConfig conf.Configuration
	json.Unmarshal(raw, &dbConfig)
	return dbConfig
}
