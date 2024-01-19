package repository

import (
	"context"
	"effective_mobile_junior/internal/model"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func (r Repository) GetPersonByParams(params model.GetPersonReq) ([]model.PersonEntity, error) {
	// Указываем, какие поля необходимо выбрать
	query := squirrel.Select("person_id", "name", "surname", "patronymic", "age", "gender_name", "gender_probability", "country_code", "country_probability").
		From("person").
		Where(conditions(params.Name, params.CountryFilter, params.GenderType, params.Older)).
		Where(squirrel.Eq{"is_deleted": false}).
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

func (r Repository) GetPersonById(id int) (model.PersonEntity, error) {
	var person model.PersonEntity

	getPersonQuery := `
	SELECT person_id, name, surname, patronymic, age, gender_name, gender_probability, country_code, country_probability
	FROM person
	WHERE person_id=$1 AND is_deleted=false
`

	err := r.db.QueryRow(context.Background(), getPersonQuery, id).Scan(
		&person.PersonId,
		&person.Personality.Name,
		&person.Personality.Surname,
		&person.Personality.Patronymic,
		&person.Age.Age,
		&person.Gender.Name,
		&person.Gender.Probability,
		&person.Country.Code,
		&person.Country.Probability)
	if err != nil {
		r.log.Debug("error select person by id",
			zap.Int("person id", id),
			zap.String("query", getPersonQuery),
			zap.String("error", err.Error()),
		)

		return model.PersonEntity{}, err
	}

	return person, nil
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
