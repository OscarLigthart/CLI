/**
 *  Helper functions to be used by the application
 */

package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
)

func loadJson(fileName string) Data{

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	
	// read json in bytes
	b, err := ioutil.ReadAll(jsonFile)

	// unmarshal into struct
	var data Data
	json.Unmarshal(b, &data)

	return data
}