# global-countries-gosdk

A Golang implementation of [global-countries](https://www.npmjs.com/package/global-countries)

### Installation

```go
go get github.com/GoodnessEzeokafor/global-countries-go
```

```go
  // get all countries
	data:= countries.Countries()	
	fmt.Println(data)
```

```go
  // get call codes
	country := "iraq"
	capital, err := countries.GetCountryCallCode(country)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(capital)
```

```go
  // get flag
	country := "iraq"
	capital, err := countries.GetCountryFlag(country)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(capital)
```


```go
  // get iso codes - Alpha-2 code/ Alpha-3 code
	country := "zambia"
	capital, err := countries.GetCountryIsoCodes(country)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(capital)
```


```go
  // get capital
	country := "nigeria"
	capital, err := countries.GetCountryCapital(country)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(capital)
```