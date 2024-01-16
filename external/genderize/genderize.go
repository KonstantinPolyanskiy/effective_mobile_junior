package genderize

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseResourceURL = "https://api.genderize.io/?name="

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

// NewEngine возвращает объект для работы с ресурсом api.genderize.io
func NewEngine(options ...Option) *Engine {
	e := &Engine{
		client: &http.Client{},
	}

	for _, opt := range options {
		opt(*e)
	}

	return e
}

type Result struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

// GenderInfoByName возвращает Result (информацию о гендере) по переданному имени
func (e Engine) GenderInfoByName(name string) (Result, error) {
	var res Result

	// Создаем url вида: https://api.genderize.io/?name=Konstantin
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
