package countries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Country struct {
	Country  string
	Capital  string
	Code     string
	Isocodes string
	Flag     string
}

func AllCountries() []Country {
	jsonFile, err := os.Open("countries.json")
	if err != nil {
		fmt.Println(err)
	} // defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data []Country
	json.Unmarshal([]byte(byteValue), &data)
	return data

}
