package model

type PostPersonReq struct {
	Personality
}

type Age struct {
	Age int `json:"age"`
}

type Personality struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Gender struct {
	// Название гендера: male, female
	Name        string  `json:"gender_name"`
	Probability float64 `json:"gender_probability"`
}

type Country struct {
	// Код страны формата EN RU UK
	Code        string
	Probability float64
}

type PersonDTO struct {
	Personality
	Age
	Gender
	Country
}

type PersonEntity struct {
	PersonId int

	Personality
	Age
	Gender
	Country
}

// GetPersonReq описывает, по каким параметрам и фильтрам нужно отдать пользователей
type GetPersonReq struct {
	Personality
	Age
	Gender
	Country

	Limit      int
	Offset     int
	GenderType string
	Older      int
	// CountryFilter указывает на то, пользователей из какой страны необходимо отдать в ответе
	CountryFilter string
}
