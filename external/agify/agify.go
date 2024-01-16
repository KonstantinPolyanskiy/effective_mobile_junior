package agify

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseResourceURL = "https://api.agify.io/?name="

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

// NewEngine возвращает объект для работы с ресурсом api.agify.io
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
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`

	Code int
}

// AgeInfoByName возвращает Result (информацию о возрасте) по переданному имени
func (e Engine) AgeInfoByName(name string) (Result, error) {
	var res struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Age   int    `json:"age"`
	}

	// Создаем url вида: https://api.agify.io/?name=Konstantin
	url := BaseResourceURL + name

	resp, err := e.client.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Count: res.Count,
		Name:  res.Name,
		Age:   res.Age,
		Code:  resp.StatusCode,
	}, nil
}
