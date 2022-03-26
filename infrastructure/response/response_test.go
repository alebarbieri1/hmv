package response

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessResponse_StatusCode(t *testing.T) {
	type fields struct {
		statusCode int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Given a SuccessResponse, StatusCode should return its statusCode field value",
			fields: fields{
				statusCode: http.StatusOK,
			},
			want: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := SuccessResponse{
				statusCode: tt.fields.statusCode,
			}

			assert.Equal(t, tt.want, e.StatusCode())
		})
	}
}

func TestSuccessResponse_Data(t *testing.T) {
	type fields struct {
		data interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Given a SuccessResponse, Data should return its data field value",
			fields: fields{
				data: "foo",
			},
			want: "foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := SuccessResponse{
				data: tt.fields.data,
			}

			assert.Equal(t, tt.want, e.Data())
		})
	}
}

func TestSuccessResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		statusCode int
		data       interface{}
	}

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Given a SuccessResponse, MarshalJSON should return its serialized data",
			fields: fields{
				statusCode: http.StatusOK,
				data:       "foo",
			},
			want:    []byte(`"foo"`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := SuccessResponse{
				statusCode: tt.fields.statusCode,
				data:       tt.fields.data,
			}

			got, err := e.MarshalJSON()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestSuccess(t *testing.T) {
	type args struct {
		statusCode int
		data       interface{}
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call Success with valid parameters, a new SuccessResponse should be created",
			args: args{
				statusCode: http.StatusOK,
				data:       "foo",
			},
			want: SuccessResponse{
				statusCode: http.StatusOK,
				data:       "foo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Success(tt.args.statusCode, tt.args.data))
		})
	}
}

func TestOK(t *testing.T) {
	type args struct {
		data interface{}
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call OK with valid parameters, a new SuccessResponse with an http.StatusOK status code should be created",
			args: args{
				data: "foo",
			},
			want: SuccessResponse{
				statusCode: http.StatusOK,
				data:       "foo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OK(tt.args.data))
		})
	}
}

func TestCreated(t *testing.T) {
	type args struct {
		data interface{}
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call Created with valid parameters, a new SuccessResponse with an http.StatusCreated status code should be created",
			args: args{
				data: "foo",
			},
			want: SuccessResponse{
				statusCode: http.StatusCreated,
				data:       "foo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Created(tt.args.data))
		})
	}
}

func TestErrorResponse_StatusCode(t *testing.T) {
	type fields struct {
		statusCode int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Given a ErrorResponse, StatusCode should return its statusCode field value",
			fields: fields{
				statusCode: http.StatusBadRequest,
			},
			want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorResponse{
				statusCode: tt.fields.statusCode,
			}

			assert.Equal(t, tt.want, e.StatusCode())
		})
	}
}

func TestErrorResponse_Data(t *testing.T) {
	type fields struct {
		message string
	}

	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Given a ErrorResponse, Data should return its data field value",
			fields: fields{
				message: "foo",
			},
			want: "foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorResponse{
				message: tt.fields.message,
			}

			assert.Equal(t, tt.want, e.Data())
		})
	}
}

func TestErrorResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		statusCode int
		message    string
	}

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Given a ErrorResponse, MarshalJSON should return its serialized message",
			fields: fields{
				statusCode: http.StatusOK,
				message:    "foo",
			},
			want:    []byte(`{"message":"foo"}`),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ErrorResponse{
				statusCode: tt.fields.statusCode,
				message:    tt.fields.message,
			}

			got, err := e.MarshalJSON()
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		statusCode int
		message    string
		err        error
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call Error with valid parameters, a new ErrorResponse should be created",
			args: args{
				statusCode: http.StatusBadRequest,
				message:    "foo",
				err:        errors.New("bar"),
			},
			want: ErrorResponse{
				statusCode: http.StatusBadRequest,
				message:    "foo: bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Error(tt.args.statusCode, tt.args.message, tt.args.err))
		})
	}
}

func TestBadRequest(t *testing.T) {
	type args struct {
		message string
		err     error
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call BadRequest with valid parameters, a new ErrorResponse with a http.StatusBadRequest status code should be created",
			args: args{
				message: "foo",
				err:     errors.New("bar"),
			},
			want: ErrorResponse{
				statusCode: http.StatusBadRequest,
				message:    "foo: bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, BadRequest(tt.args.message, tt.args.err))
		})
	}
}

func TestNotFound(t *testing.T) {
	type args struct {
		message string
		err     error
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call NotFound with valid parameters, a new ErrorResponse with a http.StatusNotFound status code should be created",
			args: args{
				message: "foo",
				err:     errors.New("bar"),
			},
			want: ErrorResponse{
				statusCode: http.StatusNotFound,
				message:    "foo: bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NotFound(tt.args.message, tt.args.err))
		})
	}
}

func TestInternalServerError(t *testing.T) {
	type args struct {
		message string
		err     error
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call InternalServerError with valid parameters, a new ErrorResponse with a http.StatusInternalServerError status code should be created",
			args: args{
				message: "foo",
				err:     errors.New("bar"),
			},
			want: ErrorResponse{
				statusCode: http.StatusInternalServerError,
				message:    "foo: bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, InternalServerError(tt.args.message, tt.args.err))
		})
	}
}

func TestForbidden(t *testing.T) {
	type args struct {
		message string
		err     error
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call Forbidden with valid parameters, a new ErrorResponse with a http.StatusForbidden status code should be created",
			args: args{
				message: "foo",
				err:     errors.New("bar"),
			},
			want: ErrorResponse{
				statusCode: http.StatusForbidden,
				message:    "foo: bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Forbidden(tt.args.message, tt.args.err))
		})
	}
}

func TestUnauthorized(t *testing.T) {
	type args struct {
		message string
		err     error
	}

	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "If I call Unauthorized with valid parameters, a new ErrorResponse with a http.StatusUnauthorized status code should be created",
			args: args{
				message: "foo",
				err:     errors.New("bar"),
			},
			want: ErrorResponse{
				statusCode: http.StatusUnauthorized,
				message:    "foo: bar",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Unauthorized(tt.args.message, tt.args.err))
		})
	}
}
