package repository

import (
	"context"
	"effective_mobile_junior/internal/model"
	"go.uber.org/zap"
)

func (r Repository) EditPerson(id int, person model.PersonDTO) (model.PersonEntity, error) {
	var updatedPerson model.PersonEntity

	updatePersonQuery := `
	UPDATE person 
	SET name=$1, surname=$2, patronymic=$3,
		age=$4, gender_name=$5, gender_probability=$6,
		country_code=$7, country_probability=$8
	WHERE person_id=$9 AND is_deleted=false
`

	err := r.db.QueryRow(context.Background(), updatePersonQuery,
		person.Personality.Name, person.Personality.Surname, person.Personality.Patronymic,
		person.Age.Age, person.Gender.Name, person.Gender.Probability,
		person.Country.Code, person.Country.Probability,
		id).
		Scan(&updatePersonQuery)
	if err != nil {
		r.log.Debug("error update person query",
			zap.Int("person id", id),
			zap.String("query", updatePersonQuery),
			zap.String("error", err.Error()),
		)
		return model.PersonEntity{}, err
	}

	return updatedPerson, nil
}
