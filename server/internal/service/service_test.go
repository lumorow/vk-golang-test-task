package service

import (
	"errors"
	"filmlib/server/internal/entity"
	mock "filmlib/server/internal/service/mocks"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
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
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    1,
			wantErr: false,
			args: args{
				actor: entity.Actor{
					Birthday: "1990-01-01",
					Name:     "john-doe",
					Sex:      "male",
				},
			},
			prepare: func(args args, fields fields) {
				fields.Actor.EXPECT().CreateActor(args.actor).Return(1, nil)
			},
		}, {
			name:    "test_2",
			want:    1,
			wantErr: true,
			args: args{
				actor: entity.Actor{
					Birthday: "1990-01-01",
					Name:     "john-doe",
					Sex:      "mmale",
				},
			},
			prepare: func(args args, fields fields) {
			},
		}, {
			name:    "test_3",
			want:    0,
			wantErr: true,
			args: args{
				actor: entity.Actor{
					Birthday: "1990-01-01",
					Name:     "john-doe",
					Sex:      "male",
				},
			},
			prepare: func(args args, fields fields) {
				fields.Actor.EXPECT().CreateActor(args.actor).Return(0, errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewActorService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			got, err := service.CreateActor(tt.args.actor)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func TestActorService_DeleteActorById(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    1,
			wantErr: false,
			args: args{
				id: 1,
			},
			prepare: func(args args, fields fields) {
				fields.Actor.EXPECT().DeleteActorById(args.id).Return(nil)
			},
		}, {
			name:    "test_2",
			want:    2,
			wantErr: false,
			args: args{
				id: 2,
			},
			prepare: func(args args, fields fields) {
				fields.Actor.EXPECT().DeleteActorById(args.id).Return(errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewActorService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			err := service.DeleteActorById(tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
		})
	}
}

func TestActorService_GetActorsWithFilms(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		ids []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.ActorFilms
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    []entity.ActorFilms{{Films: []entity.Film{}}, {Films: []entity.Film{}}, {Films: []entity.Film{}}},
			wantErr: false,
			args: args{
				ids: []int{1, 2, 3},
			},
			prepare: func(args args, fields fields) {
				for i := 0; i < len(args.ids); i++ {
					fields.Actor.EXPECT().GetActor(args.ids[i]).Return(entity.Actor{}, nil)
					fields.Film.EXPECT().GetFilmsByActorId(args.ids[i]).Return([]entity.Film{}, nil)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewActorService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			got, err := service.GetActorsWithFilms(tt.args.ids)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func TestActorService_UpdateActorById(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		id    int
		actor *entity.UpdateActorInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    1,
			wantErr: false,
			args: args{
				id:    1,
				actor: &entity.UpdateActorInput{},
			},
			prepare: func(args args, fields fields) {
				newName := "newName"
				args.actor.Name = &newName
				fields.Actor.EXPECT().UpdateActorById(args.id, *args.actor).Return(nil)
			},
		}, {
			name:    "test_2",
			want:    0,
			wantErr: true,
			args: args{
				id:    1,
				actor: &entity.UpdateActorInput{},
			},
			prepare: func(args args, fields fields) {
			},
		}, {
			name:    "test_3",
			want:    0,
			wantErr: true,
			args: args{
				id: 1,
				actor: &entity.UpdateActorInput{
					Name: new(string),
				},
			},
			prepare: func(args args, fields fields) {
				newName := "newName"
				args.actor.Name = &newName
				fields.Actor.EXPECT().UpdateActorById(args.id, *args.actor).Return(errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewActorService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			err := service.UpdateActorById(tt.args.id, *tt.args.actor)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
		})
	}
}

func TestFilmService_CreateFilm(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		actor entity.Film
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    1,
			wantErr: false,
			args: args{
				actor: entity.Film{
					Name:        "love",
					Description: "super love",
					ReleaseDay:  "1995-06-01",
					Rating:      8,
					ActorsId:    []int{1, 2, 3},
				},
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().CreateFilm(args.actor).Return(1, nil)
			},
		},
		{
			name:    "test_2",
			want:    0,
			wantErr: true,
			args: args{
				actor: entity.Film{
					Name:        "superloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperloversuperlover",
					Description: "ou",
					ReleaseDay:  "1995-06-01",
					Rating:      8,
					ActorsId:    []int{1, 2, 3},
				},
			},
			prepare: func(args args, fields fields) {
			},
		},
		{
			name:    "test_3",
			want:    0,
			wantErr: true,
			args: args{
				actor: entity.Film{
					Name: "love",
					Description: `Lorem ipsum dolor sit amet, 
consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud 
exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum 
dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. 
Suspendisse potenti nullam ac tortor vitae purus faucibus ornare suspendisse. Neque volutpat ac tincidunt vitae semper quis lectus nulla at. 
Tellus molestie nunc non blandit massa enim nec. Molestie nunc non blandit massa enim nec dui nunc. Turpis in eu mi bibendum neque egestas 
congue quisque egestas. Leo vel fringilla est ullamcorper eget nulla facilisi. Facilisis magna etiam tempor orci eu lobortis elementum nibh. 
Mollis nunc sed id semper risus in hendrerit. Odio ut enim blandit volutpat maecenas volutpat blandit aliquam etiam. Sit amet consectetur 
adipiscing elit pellentesque habitant morbi tristique senectus. Turpis egestas pretium aenean pharetra magna ac. Gravida arcu ac tortor 
dignissim convallis aenean et tortor. Diam vulputate ut pharetra sit amet aliquam. Ut porttitor leo a diam sollicitudin tempor id eu. Arcu 
vitae elementum curabitur vitae nunc sed velit dignissim sodales. In fermentum posuere urna nec tincidunt praesent semper feugiat nibh. 
Fermentum leo vel orci porta non pulvinar neque laoreet. Eu mi bibendum neque egestas congue quisque egestas diam in. Commodo elit at imperdiet 
dui accumsan sit amet nulla. Egestas diam in arcu cursus euismod quis viverra nibh cras. Quisque sagittis purus sit amet volutpat consequat 
mauris. Sit amet nulla facilisi morbi tempus. Sit amet justo donec enim diam vulputate ut pharetra. Semper risus in hendrerit gravida rutrum 
quisque non. Egestas maecenas pharetra convallis posuere morbi. Egestas purus viverra accumsan in nisl nisi scelerisque. At urna condimentum 
mattis pellentesque id nibh. Id aliquet lectus proin nibh nisl. Sem nulla pharetra diam sit amet nisl. Commodo viverra maecenas accumsan 
lacus vel facilisis volutpat est velit. Facilisi cras fermentum odio eu. Lectus quam id leo in vitae turpis. Scelerisque fermentum dui 
faucibus in ornare quam. Facilisi cras fermentum odio eu feugiat pretium nibh ipsum consequat. Fermentum et sollicitudin ac orci. Sem nulla 
pharetra diam sit amet. Nisi est sit amet facilisis magna etiam tempor orci. Eget nunc scelerisque viverra mauris in aliquam. Sit amet nisl 
purus in mollis nunc sed id. Dui sapien eget mi proin sed libero enim sed faucibus. Suspendisse interdum consectetur libero id faucibus nisl 
tincidunt eget nullam. Morbi tristique senectus et netus et malesuada. Diam quis enim lobortis scelerisque fermentum dui faucibus in. Est ante 
in nibh mauris cursus mattis molestie. Dolor sit amet consectetur adipiscing elit pellentesque habitant morbi. Sit amet dictum sit amet justo 
donec enim diam vulputate. Orci nulla pellentesque dignissim enim sit amet. Eget nunc lobortis mattis aliquam faucibus purus in. Venenatis 
cras sed felis eget velit aliquet sagittis id consectetur. Vitae congue eu consequat ac felis donec et odio pellentesque. Nunc non blandit 
massa enim nec dui nunc. Viverra nibh cras pulvinar mattis nunc sed blandit libero. Leo integer malesuada nunc vel risus commodo. Sed risus 
pretium quam vulputate dignissim suspendisse in est ante. Orci nulla pellentesque dignissim enim sit amet venenatis urna. Nisi quis eleifend 
quam adipiscing vitae. Volutpat consequat mauris nunc congue nisi vitae suscipit tellus mauris. Mattis molestie a iaculis at erat pellentesque 
adipiscing commodo. Libero
volutpat sed cras ornare arcu. Facilisis magna etiam tempor orci. Sagittis id consectetur purus ut faucibus pulvinar elementum integer. 
Varius quam quisque id diam vel quam. Id faucibus nisl tincidunt eget nullam non nisi est.`,
					ReleaseDay: "1995-06-01",
					Rating:     8,
					ActorsId:   []int{1, 2, 3},
				},
			},
			prepare: func(args args, fields fields) {
			},
		},
		{
			name:    "test_4",
			want:    0,
			wantErr: true,
			args: args{
				actor: entity.Film{
					Name:        "love",
					Description: "super love",
					ReleaseDay:  "01-06-1995",
					Rating:      8,
					ActorsId:    []int{1, 2, 3},
				},
			},
			prepare: func(args args, fields fields) {
			},
		},
		{
			name:    "test_4",
			want:    0,
			wantErr: true,
			args: args{
				actor: entity.Film{
					Name:        "love",
					Description: "super love",
					ReleaseDay:  "1995-06-01",
					Rating:      11,
					ActorsId:    []int{1, 2, 3},
				},
			},
			prepare: func(args args, fields fields) {
			},
		},
		{
			name:    "test_6",
			want:    0,
			wantErr: true,
			args: args{
				actor: entity.Film{
					Name:        "love",
					Description: "super love",
					ReleaseDay:  "1995-06-01",
					Rating:      8,
					ActorsId:    []int{1, 2, 3},
				},
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().CreateFilm(args.actor).Return(0, errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewFilmService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			got, err := service.CreateFilm(tt.args.actor)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, got, tt.want)
		})
	}
}

func TestFilmService_DeleteFilmById(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    1,
			wantErr: false,
			args: args{
				id: 1,
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().DeleteFilmById(args.id).Return(nil)
			},
		}, {
			name:    "test_2",
			want:    2,
			wantErr: false,
			args: args{
				id: 2,
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().DeleteFilmById(args.id).Return(errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewFilmService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			err := service.DeleteFilmById(tt.args.id)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
		})
	}
}

func TestFilmService_GetFilmWithFragment(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
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
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    []entity.Film{},
			wantErr: false,
			args: args{
				actorNameFrag: "fragment",
				filmNameFrag:  "fragment",
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().GetFilmsWithFragment(args.filmNameFrag, args.actorNameFrag).Return([]entity.Film{}, nil)
			},
		}, {
			name:    "test_2",
			want:    nil,
			wantErr: true,
			args: args{
				actorNameFrag: "fragment",
				filmNameFrag:  "fragment",
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().GetFilmsWithFragment(args.filmNameFrag, args.actorNameFrag).Return(nil, errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewFilmService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			res, err := service.GetFilmsWithFragment(tt.args.filmNameFrag, tt.args.filmNameFrag)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, res)
		})
	}
}

func TestFilmService_GetFilmsWithSort(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		sortType string
		ids      []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entity.Film
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    []entity.Film{},
			wantErr: false,
			args: args{
				sortType: "name",
				ids:      []int{1, 2, 3},
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().GetFilmsWithSort(args.sortType, args.ids).Return([]entity.Film{}, nil)
			},
		}, {
			name:    "test_2",
			want:    nil,
			wantErr: true,
			args: args{
				sortType: "description",
				ids:      []int{1, 2, 3},
			},
			prepare: func(args args, fields fields) {
			},
		}, {
			name:    "test_3",
			want:    []entity.Film{},
			wantErr: false,
			args: args{
				sortType: "",
				ids:      []int{1, 2, 3},
			},
			prepare: func(args args, fields fields) {
				args.sortType = "rating"
				fields.Film.EXPECT().GetFilmsWithSort(args.sortType, args.ids).Return([]entity.Film{}, nil)
			},
		}, {
			name:    "test_4",
			want:    nil,
			wantErr: true,
			args: args{
				sortType: "name",
				ids:      []int{1, 2, 3},
			},
			prepare: func(args args, fields fields) {
				fields.Film.EXPECT().GetFilmsWithSort(args.sortType, args.ids).Return(nil, errors.New(""))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewFilmService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			res, err := service.GetFilmsWithSort(tt.args.sortType, tt.args.ids)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, res)
		})
	}
}

func TestFilmService_UpdateFilmById(t *testing.T) {
	type fields struct {
		Actor *mock.MockActor
		Film  *mock.MockFilm
	}
	type args struct {
		id   int
		film *entity.UpdateFilmInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
		prepare func(args args, fields fields)
	}{
		{
			name:    "test_1",
			want:    0,
			wantErr: true,
			args: args{
				id:   1,
				film: &entity.UpdateFilmInput{},
			},
			prepare: func(args args, fields fields) {
			},
		}, {
			name:    "test_2",
			want:    0,
			wantErr: false,
			args: args{
				id: 1,
				film: &entity.UpdateFilmInput{
					Name:     new(string),
					ActorsId: new([]int),
				},
			},
			prepare: func(args args, fields fields) {
				newName := "newName"
				args.film.Name = &newName
				actorsId := []int{2, 4}
				args.film.ActorsId = &actorsId
				fields.Actor.EXPECT().GetActorsIdByFilmId(args.id).Return([]int{1, 2, 3}, nil)
				fields.Film.EXPECT().UpdateFilmById(args.id, []int{1, 3}, []int{4}, *args.film)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			f := fields{
				Actor: mock.NewMockActor(ctrl),
				Film:  mock.NewMockFilm(ctrl),
			}
			service := NewFilmService(f.Actor, f.Film)

			tt.prepare(tt.args, f)

			err := service.UpdateFilmById(tt.args.id, *tt.args.film)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
		})
	}
}
