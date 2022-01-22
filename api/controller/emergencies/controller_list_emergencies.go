package emergencies

import (
	"encoding/json"
	"flavioltonon/hmv/api/presenter"
	"net/http"
)

func (c *Controller) listEmergencies(w http.ResponseWriter, r *http.Request) {
	emergencies, err := c.usecases.Emergencies.ListEmergencies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(presenter.NewEmergencies(emergencies))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
