package emergencies

import (
	"encoding/json"
	"net/http"
)

func (c *Controller) createEmergency(w http.ResponseWriter, r *http.Request) {
	userID, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	emergency, err := c.usecases.Emergencies.CreateEmergency(userID.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(emergency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
