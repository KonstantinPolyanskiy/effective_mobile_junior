package service

import (
	"dario.cat/mergo"
	"effective_mobile_junior/internal/model"
	"errors"
	"go.uber.org/zap"
)

func (s Service) ChangePerson(id int, updatePerson model.PersonDTO) (model.PersonEntity, error) {
	// Получаем текущую сущность по ID
	currentPerson, err := s.GetPersonById(id)
	if err != nil {
		return model.PersonEntity{}, errors.New("error getting person by id")
	}

	// Создаем сущность DTO с идентичными данными из БД
	currentDTO := model.PersonDTO{
		Personality: currentPerson.Personality,
		Age:         currentPerson.Age,
		Gender:      currentPerson.Gender,
		Country:     currentPerson.Country,
	}

	err = mergo.Merge(&currentDTO, updatePerson, mergo.WithOverride)
	if err != nil {
		s.log.Info("error merging current persona and updated persona",
			zap.String("error", err.Error()),
			zap.Any("current person", currentPerson),
			zap.Any("updated person", updatePerson),
		)
		return model.PersonEntity{}, errors.New("internal server error")
	}

	return model.PersonEntity{}, nil
}
