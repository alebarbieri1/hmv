package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

// EmergencyUsecase defines all the usecases related to pacients emergencies
type EmergencyUsecase interface {
	// CreateEmergency creates a new entity.Emergency
	CreateEmergency(userID string) (*entity.Emergency, error)

	// UpdateEmergencyForm updates the EmergencyForm of a entity.Emergency with a given emergencyID. This action can only
	// be performed by users with Pacient_ProfileKind or Rescuer_ProfileKind.
	UpdateEmergencyForm(userID string, emergencyID string, form valueobject.EmergencyForm) (*entity.Emergency, error)

	// FindEmergencyByID returns an entity.Emergency with a given emergencyID. If no entities are found, entity.ErrNotFound
	// should be returned instead.
	FindEmergencyByID(emergencyID string) (*entity.Emergency, error)

	// ListEmergencies returns a list with all known entity.Emergency entities
	ListEmergencies() ([]*entity.Emergency, error)

	// ListEmergenciesByStatus returns a list with all known entity.Emergency entities that have a given valueobject.EmergencyStatus
	ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error)

	// ListEmergenciesByStatus returns a list with all known entity.Emergency entities that are related to a User
	ListEmergenciesByUser(userID string) ([]*entity.Emergency, error)

	// UpdateEmergencyStatus updates the EmergencyStatus of a entity.Emergency with a given emergencyID.
	UpdateEmergencyStatus(emergencyID string, status valueobject.EmergencyStatus) (*entity.Emergency, error)

	// SendAmbulance updates the EmergencyStatus of a entity.Emergency with a given emergencyID to AmbulanceToPacient_EmergencyStatus.
	SendAmbulance(userID string, emergencyID string) (*entity.Emergency, error)

	// RemovePacient updates the EmergencyStatus of a entity.Emergency with a given emergencyID to AmbulanceToHospital_EmergencyStatus.
	RemovePacient(userID string, emergencyID string) (*entity.Emergency, error)

	// FinishEmergencyCare updates the EmergencyStatus of a entity.Emergency with a given emergencyID to Finished_EmergencyStatus.
	FinishEmergencyCare(userID string, emergencyID string) (*entity.Emergency, error)

	// CancelEmergency updates the EmergencyStatus of a entity.Emergency with a given emergencyID to Cancelled_EmergencyStatus.
	CancelEmergency(emergencyID string) (*entity.Emergency, error)
}
