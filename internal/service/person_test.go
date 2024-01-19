package service

import (
	"context"
	"effective_mobile_junior/external/agify"
	"effective_mobile_junior/external/genderize"
	"effective_mobile_junior/external/nationalize"
	"effective_mobile_junior/internal/model"
	"effective_mobile_junior/internal/service/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestSavePerson(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		ageGetter := mocks.NewAgeGetter(t)
		countryGetter := mocks.NewCountryGetter(t)
		genderGetter := mocks.NewGenderGetter(t)

		service := New(zap.NewNop(), ageGetter, countryGetter, genderGetter, repo)

		ageGetter.On("AgeInfoByName", "Konstantin").Return(agify.Result{Age: 25, Code: http.StatusOK}, nil).Once()
		countryGetter.On("CountryInfoByName", "Konstantin").Return(nationalize.Result{Country: []nationalize.Country{{"RU", 1}}, Code: http.StatusOK}, nil).Once()
		genderGetter.On("GenderInfoByName", "Konstantin").Return(genderize.Result{Name: "male", Probability: 1, Code: http.StatusOK}, nil).Once()
		repo.On("RecordPerson", mock.AnythingOfType("model.PersonDTO")).Return(model.PersonEntity{}, nil).Once()

		ctx := context.Background()
		p := model.PostPersonReq{
			Name:       "Konstantin",
			Surname:    "TestSurname",
			Patronymic: "TestPatronymic",
		}

		recordedP, err := service.SavePerson(ctx, p)

		assert.Equal(t, nil, err)
		assert.NotNil(t, recordedP)

		ageGetter.AssertExpectations(t)
		countryGetter.AssertExpectations(t)
		genderGetter.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
	t.Run("third party error", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		ageGetter := mocks.NewAgeGetter(t)
		countryGetter := mocks.NewCountryGetter(t)
		genderGetter := mocks.NewGenderGetter(t)

		service := New(zap.NewNop(), ageGetter, countryGetter, genderGetter, repo)

		ageGetter.On("AgeInfoByName", "Konstantin").Return(agify.Result{Age: 25, Code: http.StatusUnauthorized}, nil)
		countryGetter.On("CountryInfoByName", "Konstantin").Return(nationalize.Result{Country: []nationalize.Country{{"RU", 1}}, Code: http.StatusUnauthorized}, nil)
		genderGetter.On("GenderInfoByName", "Konstantin").Return(genderize.Result{Name: "male", Probability: 1, Code: http.StatusUnauthorized}, nil)

		// Если сторонний сервис вернул не 200, мы не должны вызывать метод слоя данных
		repo.AssertNotCalled(t, "RecordPerson", mock.AnythingOfType("model.PersonDTO"))

		ctx := context.Background()
		p := model.PostPersonReq{
			Name:       "Konstantin",
			Surname:    "TestSurname",
			Patronymic: "TestPatronymic",
		}

		recordedP, err := service.SavePerson(ctx, p)

		assert.EqualError(t, err, "third-party service returned error")
		assert.NotNil(t, recordedP)

		repo.AssertExpectations(t)
	})
	t.Run("package error", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		ageGetter := mocks.NewAgeGetter(t)
		countryGetter := mocks.NewCountryGetter(t)
		genderGetter := mocks.NewGenderGetter(t)

		service := New(zap.NewNop(), ageGetter, countryGetter, genderGetter, repo)

		e := errors.New("some package error")

		ageGetter.On("AgeInfoByName", "Konstantin").Return(agify.Result{}, e)
		countryGetter.On("CountryInfoByName", "Konstantin").Return(nationalize.Result{}, e)
		genderGetter.On("GenderInfoByName", "Konstantin").Return(genderize.Result{}, e)

		// Если возникла ошибка в использовании пакета, мы не должны вызывать метод слоя данных
		repo.AssertNotCalled(t, "RecordPerson")

		ctx := context.Background()
		p := model.PostPersonReq{
			Name:       "Konstantin",
			Surname:    "TestSurname",
			Patronymic: "TestPatronymic",
		}

		recordedP, err := service.SavePerson(ctx, p)

		assert.EqualError(t, err, "internal server error")
		assert.NotNil(t, recordedP)

		repo.AssertExpectations(t)
	})
	t.Run("timeout", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		ageGetter := mocks.NewAgeGetter(t)
		countryGetter := mocks.NewCountryGetter(t)
		genderGetter := mocks.NewGenderGetter(t)

		service := New(zap.NewNop(), ageGetter, countryGetter, genderGetter, repo)

		ageGetter.On("AgeInfoByName", "Konstantin").Return(agify.Result{Age: 25, Code: http.StatusUnauthorized}, nil).Maybe()
		countryGetter.On("CountryInfoByName", "Konstantin").Return(nationalize.Result{Country: []nationalize.Country{{"RU", 1}}, Code: http.StatusUnauthorized}, nil).Maybe()
		genderGetter.On("GenderInfoByName", "Konstantin").Return(genderize.Result{Name: "male", Probability: 1, Code: http.StatusUnauthorized}, nil).Maybe()
		repo.On("RecordPerson", mock.AnythingOfType("model.PersonDTO")).Return(model.PersonEntity{}, nil).Maybe()

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		p := model.PostPersonReq{
			Name:       "Konstantin",
			Surname:    "TestSurname",
			Patronymic: "TestPatronymic",
		}

		recordedP, err := service.SavePerson(ctx, p)

		assert.EqualError(t, err, "timeout")
		assert.NotNil(t, recordedP)

		ageGetter.AssertExpectations(t)
		countryGetter.AssertExpectations(t)
		genderGetter.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
}
