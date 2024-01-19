package service

import "errors"

func (s Service) DeletePerson(id int) (bool, error) {
	isDeleted, err := s.Repository.DeletePerson(id)
	if err != nil {
		return isDeleted, errors.New("error delete person")
	}

	return isDeleted, nil
}
