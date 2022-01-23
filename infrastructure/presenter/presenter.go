package presenter

import (
	"encoding/json"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

type Presenter interface {
	Present(w http.ResponseWriter, response response.Response) error
}

type JSONPresenter struct{}

func NewJSONPresenter() Presenter {
	return new(JSONPresenter)
}

func (p *JSONPresenter) Present(w http.ResponseWriter, response response.Response) error {
	b, err := json.Marshal(response.Data())
	if err != nil {
		return err
	}

	w.WriteHeader(response.StatusCode())
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(b); err != nil {
		return err
	}

	return nil
}
