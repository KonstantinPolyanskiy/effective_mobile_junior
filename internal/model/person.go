package model

type PostPersonReq struct {
	Personality
}

type Age struct {
	Age int `json:"age" db:"age"`
}

type Personality struct {
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

type Gender struct {
	// Название гендера: male, female
	Name        string  `json:"gender_name" db:"gender_name"`
	Probability float64 `json:"gender_probability" db:"gender_probability"`
}

type Country struct {
	// Код страны формата EN RU UK
	Code        string  `json:"country_code" db:"country_code"`
	Probability float64 `json:"country_probability" db:"country_probability"`
}

type PersonDTO struct {
	Personality
	Age
	Gender
	Country
}

type PersonEntity struct {
	PersonId int `db:"person_id"`

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
