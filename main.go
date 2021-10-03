package main

import (
	"fmt"

	"github.com/GoodnessEzeokafor/global-countries-go/countries"
)

func main() {
	data := countries.Countries()
	// fmt.Println(all)
	// birdJson := `[{"species":"pigeon","description":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	// var birds []Bird
	// json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("Countries : %+v", data)
}
