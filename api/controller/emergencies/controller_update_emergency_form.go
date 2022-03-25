package emergencies

import (
	"encoding/json"
	"net/http"

	"flavioltonon/hmv/api/presenter"
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/response"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"
)

func (c *Controller) updateEmergencyForm(w http.ResponseWriter, r *http.Request) {
	user, err := entity.NewUserFromRequest(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
		c.drivers.Presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
		return
	}

	payload, err := c.newUpdateEmergencyFormPayload(r)
	if err != nil {
		c.drivers.Logger.Info(application.FailedToValidateRequest, logging.Error(err))
		c.drivers.Presenter.Present(w, response.BadRequest(application.FailedToValidateRequest, err))
		return
	}

	vars := mux.Vars(r)

	emergency, err := c.usecases.Emergencies.UpdateEmergencyForm(user.ID, vars["emergency_id"], payload.toValueObject())
	if err != nil {
		c.drivers.Logger.Error(application.FailedToUpdateEmergency, err)
		c.drivers.Presenter.Present(w, response.InternalServerError(application.FailedToUpdateEmergency, err))
		return
	}

	c.drivers.Presenter.Present(w, response.Created(presenter.NewEmergency(emergency)))
}

type updateEmergencyFormPayload struct {
	Headache struct {
		Has       *bool  `json:"has"`
		Intensity string `json:"intensity"`
	} `json:"headache"`

	BreathingDifficulties struct {
		Has *bool `json:"has"`
	} `json:"breathing_difficulties"`

	ChestPain struct {
		Has             *bool  `json:"has"`
		Characteristics string `json:"characteristics"`
	} `json:"chest_pain"`

	AbdominalPain struct {
		Has       *bool  `json:"has"`
		Intensity string `json:"intensity"`
	} `json:"abdominal_pain"`

	Backache struct {
		Has *bool `json:"has"`
	} `json:"backache"`

	BodyTemperature struct {
		CelsiusDegrees *float64 `json:"celsius_degrees"`
	} `json:"body_temperature"`

	BloodPressure struct {
		Systolic  *float64 `json:"systolic"`
		Diastolic *float64 `json:"diastolic"`
	} `json:"blood_pressure"`

	OxygenSaturation struct {
		Value *float64 `json:"value"`
	} `json:"oxygen_saturation"`
}

func (p *updateEmergencyFormPayload) Validate() error {
	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.Headache),
		ozzo.Field(&p.BreathingDifficulties),
		ozzo.Field(&p.ChestPain),
		ozzo.Field(&p.AbdominalPain),
		ozzo.Field(&p.Backache),
		ozzo.Field(&p.BodyTemperature),
		ozzo.Field(&p.BloodPressure),
		ozzo.Field(&p.OxygenSaturation),
	)
}

func (p *updateEmergencyFormPayload) toValueObject() valueobject.EmergencyForm {
	return valueobject.EmergencyForm{
		Headache: valueobject.HeadacheEmergencyFormSession{
			Has:       p.Headache.Has,
			Intensity: valueobject.NewHeadacheIntensityFromString(p.Headache.Intensity),
		},
		BreathingDifficulties: valueobject.BreathingDifficultiesEmergencyFormSession{
			Has: p.BreathingDifficulties.Has,
		},
		ChestPain: valueobject.ChestPainEmergencyFormSession{
			Has:             p.ChestPain.Has,
			Characteristics: valueobject.NewChestPainCharacteristicsFromString(p.ChestPain.Characteristics),
		},
		AbdominalPain: valueobject.AbdominalPainEmergencyFormSession{
			Has:       p.AbdominalPain.Has,
			Intensity: valueobject.NewAbdominalPainIntensityFromString(p.AbdominalPain.Intensity),
		},
		Backache: valueobject.BackacheEmergencyFormSession{
			Has: p.Backache.Has,
		},
		BodyTemperature: valueobject.BodyTemperatureEmergencyFormSession{
			CelsiusDegrees: p.BodyTemperature.CelsiusDegrees,
		},
		BloodPressure: valueobject.BloodPressureEmergencyFormSession{
			Systolic:  p.BloodPressure.Systolic,
			Diastolic: p.BloodPressure.Diastolic,
		},
		OxygenSaturation: valueobject.OxygenSaturationEmergencyFormSession{
			Value: p.OxygenSaturation.Value,
		},
	}
}

func (c *Controller) newUpdateEmergencyFormPayload(r *http.Request) (*updateEmergencyFormPayload, error) {
	var payload updateEmergencyFormPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return nil, err
	}

	if err := payload.Validate(); err != nil {
		return nil, err
	}

	return &payload, nil
}
