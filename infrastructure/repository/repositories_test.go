package repository

import (
	"flavioltonon/hmv/infrastructure/repository/memory"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepositories(t *testing.T) {
	tests := []struct {
		name    string
		want    *Repositories
		wantErr bool
	}{
		{
			name: "If I call NewRepositories, a new Repositories should be created",
			want: &Repositories{
				Analysts:    memory.NewAnalystsRepository(),
				Emergencies: memory.NewEmergenciesRepository(),
				Pacients:    memory.NewPacientsRepository(),
				Rescuers:    memory.NewRescuersRepository(),
				Users:       memory.NewUsersRepository(),
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
