package repository

import (
	"context"
	"effective_mobile_junior/internal/model"
	"github.com/Masterminds/squirrel"
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

func (r Repository) EditPerson() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) DeletePerson(id int) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) SelectPerson(params model.GetPersonReq) ([]model.PersonEntity, error) {
	// Указываем, какие поля необходимо выбрать
	query := squirrel.Select("person_id", "name", "surname", "patronymic", "age", "gender_name", "gender_probability", "country_code", "country_probability").
		From("person").
		Where(conditions(params.Name, params.CountryFilter, params.GenderType, params.Older)).
		Limit(uint64(params.Limit)).
		Offset(uint64(params.Offset)).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		r.log.Debug("error prepare query",
			zap.String("error", err.Error()))
		return nil, err
	}

	// Выводим получившийся запрос
	r.log.Debug("resulting query",
		zap.String("query", sqlQuery),
	)

	rows, err := r.db.Query(context.Background(), sqlQuery, args...)
	if err != nil {
		r.log.Debug("error select persons query",
			zap.Any("args", args))
		return nil, err
	}
	defer rows.Close()

	var persons []model.PersonEntity

	for rows.Next() {
		var person model.PersonEntity

		err := rows.Scan(
			&person.PersonId,
			&person.Personality.Name,
			&person.Personality.Surname,
			&person.Patronymic,
			&person.Age.Age,
			&person.Gender.Name,
			&person.Gender.Probability,
			&person.Country.Code,
			&person.Country.Probability)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func conditions(name, country, gender string, older int) squirrel.Sqlizer {
	cond := make([]squirrel.Sqlizer, 0)

	if name != "" {
		cond = append(cond, squirrel.Eq{"name": name})
	}

	if country != "" {
		cond = append(cond, squirrel.Eq{"country_code": country})
	}

	if older > 0 {
		cond = append(cond, squirrel.Gt{"age": older})
	}

	if gender != "" {
		cond = append(cond, squirrel.Eq{"gender_name": gender})
	}

	return squirrel.And(cond)
}
