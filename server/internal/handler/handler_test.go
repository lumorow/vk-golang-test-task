package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"filmlib/server/internal/entity"
	mock "filmlib/server/internal/handler/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestHandler_CreateActor(t *testing.T) {
	type want struct {
		msg    string
		status int
	}
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args2 args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				msg:    "id: 1",
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.Actor{
					Birthday: "1990-01-01",
					Name:     "john-doe",
					Sex:      "male",
				}
				slugInput := entity.Actor{
					Birthday: "1990-01-01",
					Name:     "john-doe",
					Sex:      "male",
				}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/api/actor", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().CreateActor(slugInput).Return(1, nil)
				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				msg:    "custom error",
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.Actor{
					Birthday: "1990-01-01",
					Name:     "John Doe",
					Sex:      "male",
				}
				slugInput := entity.Actor{
					Birthday: "1990-01-01",
					Name:     "john-doe",
					Sex:      "male",
				}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/api/actor", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().CreateActor(slugInput).Return(0, errors.New("custom error"))
				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)
			rr := httptest.NewRecorder()

			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.CreateActor(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
			require.Equal(t, tt.want.msg, rr.Body.String())
		})
	}
}

func TestHandler_CreateFilm(t *testing.T) {
	type want struct {
		msg    string
		status int
	}
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		args    args
		prepare func(args2 args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				msg:    "id: 1",
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.Film{
					ActorsId:    []int{},
					Description: "A mind-bending thriller",
					Name:        "inception",
					Rating:      8,
					ReleaseDay:  "2010-07-16",
				}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/api/film", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().CreateFilm(input).Return(1, nil)
				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				msg:    "custom error",
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.Film{
					ActorsId:    []int{},
					Description: "A mind-bending thriller",
					Name:        "inception",
					Rating:      8,
					ReleaseDay:  "2010-07-16",
				}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/api/film", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().CreateFilm(input).Return(0, errors.New("custom error"))
				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)
			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.CreateFilm(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
			require.Equal(t, tt.want.msg, rr.Body.String())
		})
	}
}

func TestHandler_DeleteActorById(t *testing.T) {
	type want struct {
		status int
	}
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		args    args
		prepare func(args2 args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("DELETE", "http://localhost:8000/api/actor/1", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().DeleteActorById(1).Return(nil)

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				status: http.StatusBadRequest,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("DELETE", "http://localhost:8000/api/actor/abc", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_3",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("DELETE", "http://localhost:8000/api/actor/1", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().DeleteActorById(1).Return(errors.New("custom error"))
				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)
			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.DeleteActorById(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_DeleteFilmById(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type want struct {
		status int
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		want    want
		args    args
		prepare func(args2 args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("DELETE", "http://localhost:8000/api/film/1", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().DeleteFilmById(1).Return(nil)

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				status: http.StatusBadRequest,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("DELETE", "http://localhost:8000/api/film/abc", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_3",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("DELETE", "http://localhost:8000/api/film/234", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().DeleteFilmById(234).Return(errors.New("custom error"))
				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.DeleteFilmById(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_GetActorsWithFilms(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args2 args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/actors?id=1,2,3", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().GetActorsWithFilms([]int{1, 2, 3}).Return([]entity.ActorFilms{}, nil)

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				status: http.StatusBadRequest,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/actors?i", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_3",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/actors?id=1,2,3", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().GetActorsWithFilms([]int{1, 2, 3}).Return(nil, errors.New("custom error"))

				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.GetActorsWithFilms(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_GetFilmsWithFragment(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/films/fragments?filmNameFr=cba&actorNameFr=abc", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().GetFilmWithFragment("abc", "cba").Return([]entity.Film{}, nil)

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/films/fragments?filmNameFr=cba&actorNameFr=abc", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().GetFilmWithFragment("abc", "cba").Return(nil, errors.New(""))

				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.GetFilmsWithFragment(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_GetFilmsWithSort(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/films/sorted/?sortType=rating&id=1,2,3", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().GetFilmsWithSort("rating", []int{1, 2, 3}).Return([]entity.Film{}, nil)

				return args.r.WithContext(ctx)
			},
		},
		{
			name: "test_2",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/films/sorted/?sortType=rating&id=1,2,3", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().GetFilmsWithSort("rating", []int{1, 2, 3}).Return(nil, errors.New(""))

				return args.r.WithContext(ctx)
			},
		},
		{
			name: "test_3",
			want: want{
				status: http.StatusBadRequest,
			},
			prepare: func(args args, fields fields) *http.Request {
				args.r, _ = http.NewRequest("GET", "http://localhost:8000/api/films/sorted/?sortType=rating&id=", nil)
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.GetFilmsWithSort(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_InitRoutes(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "test_1",
			want: []string{
				"/auth/sign-up",
				"/auth/sign-in",
				"/api/actor",
				"/api/actor/",
				"/api/actors",
				"/api/film",
				"/api/film/",
				"/api/films/fragments",
				"/api/films/sorted",
				"/api/swagger/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)
			got := handler.InitRoutes()
			for _, route := range tt.want {
				_, pattern := got.Handler(&http.Request{Method: "GET", URL: &url.URL{Path: route}})
				if pattern == "" {
					t.Errorf("Handler not found for route: %s", route)
				}
			}
		})
	}
}

func TestHandler_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authMock := mock.NewMockAuthorizationService(ctrl)
	actorMock := mock.NewMockActorService(ctrl)
	filmMock := mock.NewMockFilmService(ctrl)

	handler := NewHandler(authMock, actorMock, filmMock)

	testCases := []struct {
		name     string
		method   string
		path     string
		expected func(w *httptest.ResponseRecorder)
	}{
		{
			name:   "Sign-in",
			method: http.MethodPost,
			path:   "http://localhost:8000/auth/sign-in/",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusBadRequest)
			},
		},
		{
			name:   "Sign-up",
			method: http.MethodPost,
			path:   "http://localhost:8000/auth/sign-up/",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusBadRequest)
			},
		},
		{
			name:   "Create actor",
			method: http.MethodPost,
			path:   "http://localhost:8000/api/actor/",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusForbidden)
			},
		},
		{
			name:   "Update actor",
			method: http.MethodPatch,
			path:   "http://localhost:8000/api/actor/1",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusForbidden)
			},
		},
		{
			name:   "Delete actor",
			method: http.MethodDelete,
			path:   "http://localhost:8000/api/actor/1",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusForbidden)
			},
		},
		{
			name:   "Create film",
			method: http.MethodPost,
			path:   "http://localhost:8000/api/film/",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusForbidden)
			},
		},
		{
			name:   "Update film",
			method: http.MethodPatch,
			path:   "http://localhost:8000/api/film/1",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusForbidden)
			},
		},
		{
			name:   "Delete film",
			method: http.MethodDelete,
			path:   "http://localhost:8000/api/film/1",
			expected: func(w *httptest.ResponseRecorder) {
				require.Equal(t, w.Code, http.StatusForbidden)
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, tt.path, nil)

			ctx := context.WithValue(r.Context(), UserIdHeader, "1")
			ctx = context.WithValue(ctx, userRoleHeader, "user")
			r = r.WithContext(ctx)

			handler.ServeHTTP(w, r)
			tt.expected(w)
		})
	}
}

func TestHandler_UpdateActorById(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.UpdateActorInput{}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("PATCH", "http://localhost:8000/api/actor/1", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().UpdateActorById(1, input).Return(nil)

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				name := "abc"
				input := entity.UpdateActorInput{Name: &name}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("PATCH", "http://localhost:8000/api/actor/2", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.ActorService.EXPECT().UpdateActorById(2, input).Return(errors.New("custom error"))

				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.UpdateActorById(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_UpdateFilmById(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args args, fields fields) *http.Request
	}{
		{
			name: "test_1",
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.UpdateFilmInput{}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("PATCH", "http://localhost:8000/api/film/1", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().UpdateFilmById(1, input).Return(nil)

				return args.r.WithContext(ctx)
			},
		}, {
			name: "test_2",
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				name := "abc"
				input := entity.UpdateFilmInput{Name: &name}
				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("PATCH", "http://localhost:8000/api/film/2", bytes.NewReader(marshalled))
				ctx := context.WithValue(args.r.Context(), UserIdHeader, "1")
				ctx = context.WithValue(ctx, userRoleHeader, "admin")

				fields.FilmService.EXPECT().UpdateFilmById(2, input).Return(errors.New("custom error"))

				return args.r.WithContext(ctx)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.UpdateFilmById(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_signIn(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args args, fields fields) *http.Request
	}{
		{
			name: "Test sign in success",
			args: args{
				w: httptest.NewRecorder(),
				r: nil,
			},
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := signInInput{
					Username: "testuser",
					Password: "testpassword",
				}

				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/auth/sign-in", bytes.NewReader(marshalled))

				fields.AuthorizationService.EXPECT().GenerateToken(input.Username, input.Password).Return("mocked-token", nil)

				return args.r
			},
		}, {
			name: "Test sign in failure",
			args: args{
				w: httptest.NewRecorder(),
				r: nil,
			},
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := signInInput{
					Username: "testuser",
					Password: "testpassword2",
				}

				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/auth/sign-in", bytes.NewReader(marshalled))

				fields.AuthorizationService.EXPECT().GenerateToken(input.Username, input.Password).Return("", errors.New(""))

				return args.r
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.signIn(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_signUp(t *testing.T) {
	type fields struct {
		AuthorizationService *mock.MockAuthorizationService
		ActorService         *mock.MockActorService
		FilmService          *mock.MockFilmService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type want struct {
		status int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    want
		prepare func(args args, fields fields) *http.Request
	}{
		{
			name: "Test sign in failure",
			args: args{
				w: httptest.NewRecorder(),
				r: nil,
			},
			want: want{
				status: http.StatusOK,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.User{
					Username: "testuser",
					Password: "testpassword2",
					Role:     "admin",
				}

				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/auth/sign-in", bytes.NewReader(marshalled))

				fields.AuthorizationService.EXPECT().CreateUser(input).Return(1, nil)

				return args.r
			},
		}, {
			name: "Test sign in failure",
			args: args{
				w: httptest.NewRecorder(),
				r: nil,
			},
			want: want{
				status: http.StatusInternalServerError,
			},
			prepare: func(args args, fields fields) *http.Request {
				input := entity.User{
					Username: "testuser",
					Password: "234234",
					Role:     "adminsuper",
				}

				marshalled, _ := json.Marshal(input)
				args.r, _ = http.NewRequest("POST", "http://localhost:8000/auth/sign-in", bytes.NewReader(marshalled))

				fields.AuthorizationService.EXPECT().CreateUser(input).Return(0, errors.New(""))

				return args.r
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				AuthorizationService: mock.NewMockAuthorizationService(ctrl),
				ActorService:         mock.NewMockActorService(ctrl),
				FilmService:          mock.NewMockFilmService(ctrl),
			}
			handler := NewHandler(f.AuthorizationService, f.ActorService, f.FilmService)

			rr := httptest.NewRecorder()
			tt.args.w = rr
			tt.args.r = tt.prepare(tt.args, f)

			handler.signUp(tt.args.w, tt.args.r)

			require.Equal(t, tt.want.status, rr.Code)
		})
	}
}

func TestHandler_userIdentity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	authMock := mock.NewMockAuthorizationService(ctrl)
	actorMock := mock.NewMockActorService(ctrl)
	filmMock := mock.NewMockFilmService(ctrl)

	handler := NewHandler(authMock, actorMock, filmMock)

	type args struct {
		next http.Handler
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "Empty auth header",
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantError: true,
		},
		{
			name: "Invalid auth header",
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantError: true,
		},
		{
			name: "Successful authentication",
			args: args{
				next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()

			switch tt.name {
			case "Empty auth header":
			case "Invalid auth header":
				req.Header.Set(authorizationHeader, "Bearer")
			case "Successful authentication":
				authMock.EXPECT().ParseToken(gomock.Any()).Return(1, "admin", nil)
				token := "mocked-token"
				req.Header.Set(authorizationHeader, "Bearer "+token)
			}

			handler.userIdentity(tt.args.next).ServeHTTP(w, req)
			if tt.wantError {
				require.Equal(t, http.StatusUnauthorized, w.Code)
			} else {
				require.Equal(t, http.StatusOK, w.Code)
			}
		})
	}
}
