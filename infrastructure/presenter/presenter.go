package presenter

import (
	"encoding/json"
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
)

// Presenter is a HTTP response presenter
type Presenter interface {
	Present(w http.ResponseWriter, response response.Response) error
}

// JSONPresenter is a JSON Presenter
type JSONPresenter struct{}

// NewJSONPresenter creates a new JSONPresenter
func NewJSONPresenter() Presenter {
	return new(JSONPresenter)
}

// Present serializes a presentable response.Response, writing it to the HTTTP response using a JSON codec
func (p *JSONPresenter) Present(w http.ResponseWriter, response response.Response) error {
	b, err := json.Marshal(response.Data())
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(response.StatusCode())

	if _, err := w.Write(b); err != nil {
		return err
	}

	return nil
}
