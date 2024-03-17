package handler

import (
	"net/http"
	"reflect"
	"testing"
)

func TestHandler_CreateActor(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.CreateActor(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_CreateFilm(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.CreateFilm(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_DeleteActorById(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.DeleteActorById(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_DeleteFilmById(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.DeleteFilmById(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_GetActorsWithFilms(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.GetActorsWithFilms(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_GetFilmsWithFragment(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.GetFilmsWithFragment(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_GetFilmsWithSort(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.GetFilmsWithSort(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_InitRoutes(t *testing.T) {
	type fields struct {
		Service Service
	}
	tests := []struct {
		name   string
		fields fields
		want   *http.ServeMux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			if got := h.InitRoutes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitRoutes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.ServeHTTP(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_UpdateActorById(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.UpdateActorById(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_UpdateFilmById(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.UpdateFilmById(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.signIn(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_signUp(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			h.signUp(tt.args.w, tt.args.r)
		})
	}
}

func TestHandler_userIdentity(t *testing.T) {
	type fields struct {
		Service Service
	}
	type args struct {
		next http.Handler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				Service: tt.fields.Service,
			}
			if got := h.userIdentity(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHandler(t *testing.T) {
	type args struct {
		service Service
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkAdminRule(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkAdminRule(tt.args.w, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("checkAdminRule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getUserId(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUserId(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newErrorResponse(t *testing.T) {
	type args struct {
		w          http.ResponseWriter
		statusCode int
		message    string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newErrorResponse(tt.args.w, tt.args.statusCode, tt.args.message)
		})
	}
}
