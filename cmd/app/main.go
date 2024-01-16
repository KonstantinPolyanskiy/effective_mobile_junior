package main

import (
	"effective_mobile_junior/external/agify"
	"effective_mobile_junior/external/genderize"
	"effective_mobile_junior/external/nationalize"
	"fmt"
	"net/http"
)

const Name = "Konstantin"

func main() {
	client := http.Client{}

	agifyEngine := agify.NewEngine(agify.WithCustomClient(&client))
	genderizeEngine := genderize.NewEngine(genderize.WithCustomClient(&client))
	nationalizeEngine := nationalize.NewEngine(nationalize.WithCustomClient(&client))

	ageInfo, _ := agifyEngine.AgeInfoByName(Name)
	genderInfo, _ := genderizeEngine.GenderInfoByName(Name)
	countryInfo, _ := nationalizeEngine.CountryInfoByName(Name)

	fmt.Printf("Информация о возрасте - %v\n", ageInfo)
	fmt.Printf("Информация о гендере - %v\n", genderInfo)
	fmt.Printf("Информация о стране - %v\n", countryInfo)
}
