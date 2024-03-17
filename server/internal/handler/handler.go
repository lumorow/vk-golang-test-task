package handler

import (
	_ "filmlib/server/docs"
	"filmlib/server/internal/entity"
	"net/http"
	"regexp"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

//go:generate mockgen -destination=mocks/handler.go -package=mock -source=handler.go
//go:generate touch mocks/.coverignore

var (
	AuthSignUpRe             = regexp.MustCompile(`^/auth/sign-up/*$`)
	AuthSignInRe             = regexp.MustCompile(`^/auth/sign-in/*$`)
	ActorRe                  = regexp.MustCompile(`^/api/actor/*$`)
	ActorReWithID            = regexp.MustCompile(`^/api/actor/([0-9]+)$`)
	ActorsReWithID           = regexp.MustCompile(`^/api/actors\?id=[0-9]+(,[0-9]+)*$`)
	FilmRe                   = regexp.MustCompile(`^/api/film/*$`)
	FilmReWithID             = regexp.MustCompile(`^/api/film/([0-9]+)$`)
	FilmsReWithIDAndFragment = regexp.MustCompile(`^/api/films/fragments(\?filmNameFr=[a-zA-Z0-9]+&)?(actorNameFr=[a-zA-Z0-9]+$)?`)
	FilmsReWithIDAndWithSort = regexp.MustCompile(`^/api/films/sorted\?(sortType=[a-zA-Z0-9]+&)?(id=[0-9]+(,[0-9]+)*)$`)
)

type Service interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, string, error)
	CreateActor(actor entity.Actor) (int, error)
	DeleteActorById(id int) error
	UpdateActorById(id int, actor entity.UpdateActorInput) error
	GetActorsWithFilms(actorsId []int) ([]entity.ActorFilms, error)
	CreateFilm(film entity.Film) (int, error)
	DeleteFilmById(id int) error
	UpdateFilmById(id int, film entity.UpdateFilmInput) error
	GetFilmWithFragment(actorNameFrag, filmNameFrag string) ([]entity.Film, error)
	GetFilmsWithSort(sortType string, filmsId []int) ([]entity.Film, error)
}

type Handler struct {
	Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	middleware := h.userIdentity(h)

	mux.Handle("/auth/sign-up", h)
	mux.Handle("/auth/sign-in", h)

	mux.Handle("/api/actor", middleware)
	mux.Handle("/api/actor/", middleware)
	mux.Handle("/api/actors", middleware)

	mux.Handle("/api/film", middleware)
	mux.Handle("/api/film/", middleware)
	mux.Handle("/api/films/fragments", middleware)
	mux.Handle("/api/films/sorted", middleware)

	mux.Handle("/api/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/api/swagger/doc.json"),
	))

	return mux
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// Sign-in
	case r.Method == http.MethodPost && AuthSignInRe.MatchString(r.URL.RequestURI()):
		h.signIn(w, r)
		return

	// Sign-up
	case r.Method == http.MethodPost && AuthSignUpRe.MatchString(r.URL.RequestURI()):
		h.signUp(w, r)
		return

	// Create actor
	case r.Method == http.MethodPost && ActorRe.MatchString(r.URL.RequestURI()):
		h.CreateActor(w, r)
		return

	// Update actor
	case r.Method == http.MethodPatch && ActorReWithID.MatchString(r.URL.RequestURI()):
		h.UpdateActorById(w, r)
		return

	// Delete actor
	case r.Method == http.MethodDelete && ActorReWithID.MatchString(r.URL.RequestURI()):
		h.DeleteActorById(w, r)
		return

	// Get actors with films
	case r.Method == http.MethodGet && ActorsReWithID.MatchString(r.URL.RequestURI()):
		h.GetActorsWithFilms(w, r)
		return

	// Create film
	case r.Method == http.MethodPost && FilmRe.MatchString(r.URL.RequestURI()):
		h.CreateFilm(w, r)
		return

	// Update film
	case r.Method == http.MethodPatch && FilmReWithID.MatchString(r.URL.RequestURI()):
		h.UpdateFilmById(w, r)
		return

	// Delete film
	case r.Method == http.MethodDelete && FilmReWithID.MatchString(r.URL.RequestURI()):
		h.DeleteFilmById(w, r)
		return

	// Get films with sort
	case r.Method == http.MethodGet && FilmsReWithIDAndWithSort.MatchString(r.URL.RequestURI()):
		h.GetFilmsWithSort(w, r)
		return

	// Get films with actor name fragment or film name fragment
	case r.Method == http.MethodGet && FilmsReWithIDAndFragment.MatchString(r.URL.RequestURI()):
		h.GetFilmsWithFragment(w, r)
		return

	// Another Path
	default:
		newErrorResponse(w, http.StatusNotFound, "404 Not Found")
		return
	}
}
