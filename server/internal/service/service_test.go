package service

import (
	"filmlib/server/internal/entity"
	mock "filmlib/server/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestActorService_CreateActor(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		actor entity.Actor
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			s := &ActorService{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			got, err := s.CreateActor(tt.args.actor)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateActor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateActor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActorService_DeleteActorById(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ActorService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			if err := s.DeleteActorById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteActorById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActorService_GetActorsWithFilms(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		actorsId []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.ActorFilms
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ActorService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			got, err := s.GetActorsWithFilms(tt.args.actorsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActorsWithFilms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActorsWithFilms() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestActorService_UpdateActorById(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		id    int
		actor entity.UpdateActorInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ActorService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			if err := s.UpdateActorById(tt.args.id, tt.args.actor); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActorById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorizationService_CreateUser(t *testing.T) {
	type fields struct {
		Authorization Authorization
		roles         map[string]struct{}
	}
	type args struct {
		user entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthorizationService{
				Authorization: tt.fields.Authorization,
				roles:         tt.fields.roles,
			}
			got, err := s.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizationService_GenerateToken(t *testing.T) {
	type fields struct {
		Authorization Authorization
		roles         map[string]struct{}
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthorizationService{
				Authorization: tt.fields.Authorization,
				roles:         tt.fields.roles,
			}
			got, err := s.GenerateToken(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizationService_ParseToken(t *testing.T) {
	type fields struct {
		Authorization Authorization
		roles         map[string]struct{}
	}
	type args struct {
		accessToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthorizationService{
				Authorization: tt.fields.Authorization,
				roles:         tt.fields.roles,
			}
			got, got1, err := s.ParseToken(tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFilmService_CreateFilm(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		film entity.Film
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FilmService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			got, err := s.CreateFilm(tt.args.film)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFilm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateFilm() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilmService_DeleteFilmById(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FilmService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			if err := s.DeleteFilmById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFilmById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFilmService_GetFilmWithFragment(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		actorNameFrag string
		filmNameFrag  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Film
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FilmService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			got, err := s.GetFilmWithFragment(tt.args.actorNameFrag, tt.args.filmNameFrag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilmWithFragment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilmWithFragment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilmService_GetFilmsWithSort(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		sortType string
		filmsId  []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Film
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FilmService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			got, err := s.GetFilmsWithSort(tt.args.sortType, tt.args.filmsId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilmsWithSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilmsWithSort() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilmService_UpdateFilmById(t *testing.T) {
	type fields struct {
		Actor Actor
		Film  Film
	}
	type args struct {
		filmId int
		film   entity.UpdateFilmInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &FilmService{
				Actor: tt.fields.Actor,
				Film:  tt.fields.Film,
			}
			if err := s.UpdateFilmById(tt.args.filmId, tt.args.film); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFilmById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewActorService(t *testing.T) {
	type args struct {
		actorRepository Actor
		filmRepository  Film
	}
	tests := []struct {
		name string
		args args
		want *ActorService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewActorService(tt.args.actorRepository, tt.args.filmRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewActorService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAuthorizationService(t *testing.T) {
	type args struct {
		authRepository Authorization
	}
	tests := []struct {
		name string
		args args
		want *AuthorizationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthorizationService(tt.args.authRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthorizationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFilmService(t *testing.T) {
	type args struct {
		actorRepository Actor
		filmRepository  Film
	}
	tests := []struct {
		name string
		args args
		want *FilmService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFilmService(tt.args.actorRepository, tt.args.filmRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFilmService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generatePasswordHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePasswordHash(tt.args.password); got != tt.want {
				t.Errorf("generatePasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
