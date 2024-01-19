package service

import (
	"effective_mobile_junior/internal/model"
	"go.uber.org/zap"
)

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=Repository
type Repository interface {
	RecordPerson(person model.PersonDTO) (model.PersonEntity, error)
	EditPerson(id int, person model.PersonDTO) (model.PersonEntity, error)
	SoftDeletePerson(id int) (bool, error)
	GetPersonByParams(params model.GetPersonReq) ([]model.PersonEntity, error)
	GetPersonById(id int) (model.PersonEntity, error)
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
