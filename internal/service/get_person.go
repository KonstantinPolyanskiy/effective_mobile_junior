package service

import (
	"effective_mobile_junior/internal/model"
	"errors"
	"go.uber.org/zap"
	"strings"
)

func (s Service) GetPerson(params model.GetPersonReq) ([]model.PersonEntity, error) {
	// Коды стран храним в нижнем регистре
	params.CountryFilter = strings.ToLower(params.CountryFilter)

	s.log.Info("parameterized search",
		zap.String("first name", params.Name),
		zap.String("gender type", params.GenderType),
		zap.String("country code", params.CountryFilter),
		zap.Int("older then", params.Older),
		zap.Int("limit", params.Limit),
		zap.Int("offset", params.Offset),
	)

	persons, err := s.Repository.GetPersonByParams(params)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return persons, nil
}
