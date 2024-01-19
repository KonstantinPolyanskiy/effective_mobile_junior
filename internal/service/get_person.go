package service

import (
	"effective_mobile_junior/internal/model"
	"errors"
)

func (s Service) GetPerson(params model.GetPersonReq) ([]model.PersonEntity, error) {
	persons, err := s.Repository.GetPersonByParams(params)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return persons, nil
}
