package service

import (
	"context"
	"effective_mobile_junior/external/agify"
	"effective_mobile_junior/external/genderize"
	"effective_mobile_junior/external/nationalize"
	"effective_mobile_junior/internal/model"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"sync"
)

type AgeGetter interface {
	AgeInfoByName(name string) (agify.Result, error)
}

type GenderGetter interface {
	GenderInfoByName(name string) (genderize.Result, error)
}

type CountryGetter interface {
	CountryInfoByName(name string) (nationalize.Result, error)
}

func (s Service) SavePerson(ctx context.Context, person model.PostPersonReq) (model.PersonEntity, error) {
	var (
		ageRes     agify.Result
		genderRes  genderize.Result
		countryRes nationalize.Result
	)
	// Аллокация каналов с результатами стороннего API
	ageCh := make(chan agify.Result, 1)
	genderCh := make(chan genderize.Result, 1)
	countryCh := make(chan nationalize.Result, 1)

	// Канал ошибок от пакетов для работы со сторонним API
	errCh := make(chan error, 3)

	defer func() {
		close(ageCh)
		close(genderCh)
		close(countryCh)
		close(errCh)
	}()

	for {
		select {
		case <-ctx.Done():
			return model.PersonEntity{}, errors.New("timeout")
		default:
			var wg sync.WaitGroup
			wg.Add(3)

			// Записываем в канал данные от вызова сторонних API
			go s.ageInfo(person.Name, &wg, ageCh, errCh)
			go s.genderInfo(person.Name, &wg, genderCh, errCh)
			go s.countryInfo(person.Name, &wg, countryCh, errCh)

			// Ждем их исполнения
			wg.Wait()

			// Проверяем: если ли в канале ошибка и выходим, в противном случае записываем данные из каналов
			select {
			case err := <-errCh:
				if err != nil {
					return model.PersonEntity{}, errors.New("internal server error")
				}
			default:
				ageRes = <-ageCh
				genderRes = <-genderCh
				countryRes = <-countryCh
			}

			// Проверяем, не вернули ли сторонние API коды ошибок
			if ageRes.Code != http.StatusOK || genderRes.Code != http.StatusOK || countryRes.Code != http.StatusOK {
				s.log.Debug("third party service response not 200",
					zap.Int("age response", ageRes.Code),
					zap.Int("gender response", genderRes.Code),
					zap.Int("country response", countryRes.Code),
				)
				return model.PersonEntity{}, errors.New("third-party service returned error")
			}

			code, chance := mostProbableCountry(countryRes)

			// Передаем данные в слой данных и получаем записанный результат
			dto := model.PersonDTO{
				Personality: person.Personality,
				Age:         model.Age{Age: ageRes.Age},
				Gender: model.Gender{
					Name:        genderRes.Name,
					Probability: genderRes.Probability,
				},
				Country: model.Country{
					Code:        code,
					Probability: chance,
				},
			}

			recordedPerson, err := s.Repository.RecordPerson(dto)
			if err != nil {
				return model.PersonEntity{}, errors.New("person save error")
			}

			return recordedPerson, nil
		}
	}
}

func (s Service) ageInfo(name string, wg *sync.WaitGroup, ch chan<- agify.Result, errCh chan<- error) {
	defer wg.Done()

	result, err := s.AgeGetter.AgeInfoByName(name)
	if err != nil {
		errCh <- err
		return
	}
	ch <- result
}

func (s Service) genderInfo(name string, wg *sync.WaitGroup, ch chan<- genderize.Result, errCh chan<- error) {
	defer wg.Done()

	result, err := s.GenderGetter.GenderInfoByName(name)
	if err != nil {
		errCh <- err
		return
	}
	ch <- result
}

func (s Service) countryInfo(name string, wg *sync.WaitGroup, ch chan<- nationalize.Result, errCh chan<- error) {
	defer wg.Done()

	result, err := s.CountryGetter.CountryInfoByName(name)
	if err != nil {
		errCh <- err
		return
	}

	ch <- result
}

// mostProbableCountry возвращает код наиболее вероятной страны и шанс этого
func mostProbableCountry(country nationalize.Result) (string, float64) {
	code := country.Country[0].CountryId
	chance := country.Country[0].Probability

	for _, c := range country.Country {
		if c.Probability > chance {
			chance = c.Probability
			code = c.CountryId
		}
	}

	return code, chance
}
