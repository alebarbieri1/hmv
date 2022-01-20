package emergencies

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) listEmergencies(w http.ResponseWriter, r *http.Request) {
	emergencies, err := c.usecases.Emergencies.ListEmergencies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(emergencies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
