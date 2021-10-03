package countries

import (
	"errors"
	"strings"

	"github.com/GoodnessEzeokafor/global-countries-go/countries"
)

/*
 * Call this function to get all the available countries
 * @returns all the countries
 */
func Countries() []countries.Country {
	return countries.Countries()
}

/*
 * Call this function to get a country call code
 * @returns all a country call code
 */
func GetCountryCallCode(country string) (string, error) {
	if country == "" {
		return "", errors.New("invalid country")
	}
	var result string
	countries := countries.Countries()
	for i := 0; i < len(countries); i++ {
		if countries[i].Country == strings.ToLower(country) {
			result = countries[i].Code
		}
	}
	if result == "" {
		return "", errors.New("invalid country")
	}
	return result, nil
}

/*
 * Call this function to get a country flag link
 * @returns all a country flag
 */
func GetCountryFlag(country string) (string, error) {
	if country == "" {
		return "", errors.New("invalid country")
	}
	var result string
	countries := countries.Countries()
	for i := 0; i < len(countries); i++ {
		if countries[i].Country == strings.ToLower(country) {
			result = countries[i].Flag
		}
	}
	if result == "" {
		return "", errors.New("invalid country")
	}
	return result, nil
}

/*
 * Call this function to get a country Alpha-2 code/ Alpha-3 code
 * @returns Alpha-2 code/ Alpha-3 code
 */
func GetCountryIsoCodes(country string) (string, error) {
	if country == "" {
		return "", errors.New("invalid country")
	}
	var result string
	countries := countries.Countries()
	for i := 0; i < len(countries); i++ {
		if countries[i].Country == strings.ToLower(country) {
			result = countries[i].Isocodes
		}
	}
	if result == "" {
		return "", errors.New("invalid country")
	}
	return result, nil
}

/*
 * Call this function to get a country capital
 * @returns country capital
 */
func GetCountryCapital(country string) (string, error) {
	if country == "" {
		return "", errors.New("invalid country")
	}
	var result string
	countries := countries.Countries()
	for i := 0; i < len(countries); i++ {
		if countries[i].Country == strings.ToLower(country) {
			result = countries[i].Capital
		}
	}
	if result == "" {
		return "", errors.New("invalid country")
	}
	return result, nil
}

// func main() {
// 	result, err := GetCountryIsoCodes("NIGERIA")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }
