-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS Person (
    person_id SERIAL PRIMARY KEY NOT NULL,

    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    patronymic VARCHAR NOT NULL DEFAULT '',

    age INTEGER NOT NULL,

    gender_name VARCHAR NOT NULL,
    gender_probability REAL NOT NULL,

    country_code VARCHAR(2) NOT NULL,
    country_probability REAL NOT NULL
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS Person;
-- +goose StatementEnd
