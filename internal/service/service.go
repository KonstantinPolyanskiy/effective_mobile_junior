package service

import "go.uber.org/zap"

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
