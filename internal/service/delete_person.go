package service

import (
	"errors"
	"go.uber.org/zap"
)

func (s Service) DeletePerson(id int) (bool, error) {
	s.log.Info("deleting person",
		zap.Int("id", id),
	)

	isDeleted, err := s.Repository.SoftDeletePerson(id)
	if err != nil {
		return isDeleted, errors.New("error delete person")
	}

	return isDeleted, nil
}
