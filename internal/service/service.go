package service

import (
	"effective_mobile_junior/internal/model"
	"go.uber.org/zap"
)

type Repository interface {
	RecordPerson(person model.PersonDTO) (model.PersonEntity, error)
	EditPerson()
	DeletePerson(id int) (bool, error)
	SelectPerson(params model.GetPersonReq) ([]model.PersonEntity, error)
}

type Service struct {
	log *zap.Logger

	// Внешние зависимости
	AgeGetter
	CountryGetter
	GenderGetter

	// Слой данных
	Repository
}

func New(log *zap.Logger, ag AgeGetter, cg CountryGetter, gg GenderGetter, repository Repository) Service {
	return Service{
		AgeGetter:     ag,
		CountryGetter: cg,
		GenderGetter:  gg,
		Repository:    repository,
		log:           log,
	}
}
