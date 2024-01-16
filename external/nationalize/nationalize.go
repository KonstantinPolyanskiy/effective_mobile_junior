package nationalize

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseResourceURL = "https://api.nationalize.io/?name="

type Option func(engine Engine)

// WithCustomClient позволяет установить настроенный http.Client
func WithCustomClient(client *http.Client) Option {
	return func(engine Engine) {
		engine.client = client
	}
}

type Engine struct {
	client *http.Client
}

// NewEngine возвращает объект для работы с ресурсом api.nationalize.io
func NewEngine(options ...Option) *Engine {
	e := &Engine{
		client: &http.Client{},
	}

	for _, opt := range options {
		opt(*e)
	}

	return e
}

type Country struct {
	CountryId   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type Result struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []Country `json:"country"`
}

// CountryInfoByName возвращает Result (информацию о стране) по переданному имени
func (e Engine) CountryInfoByName(name string) (Result, error) {
	var res Result

	// Создаем url вида: https://api.nationalize.io/?name=Konstantin
	url := BaseResourceURL + name

	resp, err := e.client.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
