package model

type PostPersonReq struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
}

type PatchPersonReq struct {
	Name       string `json:"name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        int    `json:"age,omitempty"`
	Gender     `json:",omitempty"`
	Country    `json:"country,omitempty"`
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

type PersonEntity struct {
	PersonId int `json:"person_id" db:"person_id"`

	Personality
	Age
	Gender
	Country
}

type PersonDTO struct {
	Personality `json:",omitempty"`
	Age         `json:",omitempty"`
	Gender      `json:",omitempty"`
	Country     `json:",omitempty"`
}

// GetPersonReq описывает, по каким параметрам и фильтрам нужно отдать пользователей
type GetPersonReq struct {
	Name       string
	Limit      int
	Offset     int
	GenderType string
	Older      int
	// CountryFilter указывает на то, пользователей из какой страны необходимо отдать в ответе
	CountryFilter string
}
