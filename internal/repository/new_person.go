package repository

import (
	"context"
	"effective_mobile_junior/internal/model"
	"go.uber.org/zap"
)

func (r Repository) RecordPerson(person model.PersonDTO) (model.PersonEntity, error) {
	var recordedPersonId int
	var recordedPerson model.PersonEntity

	insertPersonQuery := `
	INSERT INTO person (name, surname, patronymic, age, gender_name, gender_probability, country_code, country_probability) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING person_id
`
	// Вставляем в БД полученного пользователя
	err := r.db.QueryRow(context.Background(), insertPersonQuery,
		person.Personality.Name, person.Personality.Surname, person.Personality.Patronymic,
		person.Age.Age,
		person.Gender.Name, person.Gender.Probability,
		person.Country.Code, person.Country.Probability).
		Scan(&recordedPersonId)
	if err != nil {
		r.log.Debug("error exec insert person query",
			zap.String("query", insertPersonQuery),
			zap.Any("recorded person", person),
			zap.String("error", err.Error()),
		)

		return model.PersonEntity{}, err
	}

	getInsertedPersonQuery := `
	SELECT person_id, name, surname, patronymic, age, gender_name, gender_probability, country_code, country_probability 
	FROM person
	WHERE person_id=$1
`

	// Получаем только что вставленного пользователя
	err = r.db.QueryRow(context.Background(), getInsertedPersonQuery, recordedPersonId).Scan(
		&recordedPerson.PersonId,
		&recordedPerson.Personality.Name,
		&recordedPerson.Personality.Surname,
		&recordedPerson.Patronymic,
		&recordedPerson.Age.Age,
		&recordedPerson.Gender.Name,
		&recordedPerson.Gender.Probability,
		&recordedPerson.Country.Code,
		&recordedPerson.Country.Probability)
	if err != nil {
		r.log.Debug("error exec get inserted person query",
			zap.Int("person id", recordedPersonId),
			zap.String("query", getInsertedPersonQuery),
			zap.String("error", err.Error()),
		)

		return model.PersonEntity{}, err
	}

	return recordedPerson, nil
}
