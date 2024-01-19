package service

import (
	"effective_mobile_junior/internal/model"
	"errors"
	"strings"
)

func (s Service) GetPerson(params model.GetPersonReq) ([]model.PersonEntity, error) {
	// Коды стран храним в нижнем регистре
	params.CountryFilter = strings.ToLower(params.CountryFilter)

	persons, err := s.Repository.GetPersonByParams(params)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return persons, nil
}
