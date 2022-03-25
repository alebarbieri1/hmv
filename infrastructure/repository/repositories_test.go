package repository

import (
	"flavioltonon/hmv/infrastructure/repository/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepositories(t *testing.T) {
	var (
		analysts, _    = memory.NewAnalystsRepository()
		emergencies, _ = memory.NewEmergenciesRepository()
		pacients, _    = memory.NewPacientsRepository()
		rescuers, _    = memory.NewRescuersRepository()
		users, _       = memory.NewUsersRepository()
	)

	tests := []struct {
		name    string
		want    *Repositories
		wantErr bool
	}{
		{
			name: "If I call NewRepositories, a new Repositories should be created",
			want: &Repositories{
				Analysts:    analysts,
				Emergencies: emergencies,
				Pacients:    pacients,
				Rescuers:    rescuers,
				Users:       users,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRepositories()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
