package presenter

import (
	"flavioltonon/hmv/infrastructure/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJSONPresenter(t *testing.T) {
	tests := []struct {
		name string
		want Presenter
	}{
		{
			name: "If I call NewJSONPresenter, a new JSONPresenter should be returned",
			want: new(JSONPresenter),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewJSONPresenter())
		})
	}
}

func TestJSONPresenter_Present(t *testing.T) {
	type args struct {
		response response.Response
	}

	tests := []struct {
		name           string
		args           args
		wantErr        bool
		wantStatusCode int
		want           []byte
	}{
		{
			name: "Given a response.Response, JSONPresenter.Present should serialize it using a JSON codec",
			args: args{
				response: response.OK("foo"),
			},
			wantErr:        false,
			wantStatusCode: http.StatusOK,
			want:           []byte(`"foo"`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &JSONPresenter{}

			recorder := httptest.NewRecorder()

			err := p.Present(recorder, tt.args.response)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantStatusCode, recorder.Result().StatusCode)
			assert.Equal(t, tt.want, recorder.Body.Bytes())
		})
	}
}
