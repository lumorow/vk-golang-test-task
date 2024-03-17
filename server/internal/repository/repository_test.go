package repository

import (
	"filmlib/server/internal/entity"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestMigrate(t *testing.T) {
	type args struct {
		db *sqlx.DB
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
			if err := Migrate(tt.args.db); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewPostgresDB(t *testing.T) {
	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		args    args
		want    *sqlx.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPostgresDB(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPostgresDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPostgresDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRepository(t *testing.T) {
	type args struct {
		db DBTX
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreateActor(t *testing.T) {
	type fields struct {
		db DBTX
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.CreateActor(tt.args.actor)
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

func TestRepository_CreateFilm(t *testing.T) {
	type fields struct {
		db DBTX
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
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.CreateFilm(tt.args.film)
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

func TestRepository_CreateUser(t *testing.T) {
	type fields struct {
		db DBTX
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
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.CreateUser(tt.args.user)
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

func TestRepository_DeleteActorById(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		actorId int
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
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.DeleteActorById(tt.args.actorId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteActorById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_DeleteFilmById(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		filmId int
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
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.DeleteFilmById(tt.args.filmId); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFilmById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetActor(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		actorId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetActor(tt.args.actorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetActorsIdByFilmId(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		filmId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetActorsIdByFilmId(tt.args.filmId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActorsIdByFilmId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActorsIdByFilmId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetFilmsByActorId(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		actorId int
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
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetFilmsByActorId(tt.args.actorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilmsByActorId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilmsByActorId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetFilmsWithFragment(t *testing.T) {
	type fields struct {
		db DBTX
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
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetFilmsWithFragment(tt.args.actorNameFrag, tt.args.filmNameFrag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilmsWithFragment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilmsWithFragment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetFilmsWithSort(t *testing.T) {
	type fields struct {
		db DBTX
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
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetFilmsWithSort(tt.args.sortType, tt.args.filmsId)
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

func TestRepository_GetUser(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				db: tt.fields.db,
			}
			got, err := r.GetUser(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_UpdateActorById(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		actorId int
		actor   entity.UpdateActorInput
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
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.UpdateActorById(tt.args.actorId, tt.args.actor); (err != nil) != tt.wantErr {
				t.Errorf("UpdateActorById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdateFilmById(t *testing.T) {
	type fields struct {
		db DBTX
	}
	type args struct {
		filmId    int
		deleteIds []int
		addIds    []int
		film      entity.UpdateFilmInput
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
			r := &Repository{
				db: tt.fields.db,
			}
			if err := r.UpdateFilmById(tt.args.filmId, tt.args.deleteIds, tt.args.addIds, tt.args.film); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFilmById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getActorFragmentFilmFragmentQuery(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getActorFragmentFilmFragmentQuery(); got != tt.want {
				t.Errorf("getActorFragmentFilmFragmentQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getActorFragmentQuery(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getActorFragmentQuery(); got != tt.want {
				t.Errorf("getActorFragmentQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFilmFragmentQuery(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFilmFragmentQuery(); got != tt.want {
				t.Errorf("getFilmFragmentQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
