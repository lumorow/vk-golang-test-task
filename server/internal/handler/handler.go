package handler

import (
	_ "filmlib/server/docs"
	"filmlib/server/internal/service"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"net/http"
	"regexp"
)

var (
	AuthSignUpRe             = regexp.MustCompile(`^/auth/sign-up/*$`)
	AuthSignInRe             = regexp.MustCompile(`^/auth/sign-in/*$`)
	ActorRe                  = regexp.MustCompile(`^/api/actor/*$`)
	ActorReWithID            = regexp.MustCompile(`^/api/actor/([0-9]+)$`)
	ActorsReWithID           = regexp.MustCompile(`^/api/actors\?(id=[0-9]+&)*id=[0-9]+$`)
	FilmRe                   = regexp.MustCompile(`^/api/film/*$`)
	FilmReWithID             = regexp.MustCompile(`^/api/film/([0-9]+)$`)
	FilmsReWithIDAndFragment = regexp.MustCompile(`^/api/films\?filmNameFr=[a-zA-Z0-9]+&actorNameFr=[a-zA-Z0-9]+$`)
	FilmsReWithIDAndWithSort = regexp.MustCompile(`^/api/films/sortType/?(id=[0-9]+&)*id=[0-9]+$`)
	SwaggerRe                = regexp.MustCompile(`^/api/swagger/*$`)
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	middleware := h.userIdentity(h)

	mux.Handle("/auth/sign-in", h)
	mux.Handle("/auth/sign-up", h)

	mux.Handle("/api/actor", middleware)
	mux.Handle("/api/actor/", middleware)
	mux.Handle("/api/actors/", middleware)

	mux.Handle("/api/film", middleware)
	mux.Handle("/api/film/", middleware)
	mux.Handle("/api/films", middleware)
	mux.Handle("/api/films/", middleware)

	// TODO: need implemented swagger API
	mux.Handle("/api/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/api/swagger/doc.json"),
	))

	return mux
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// Sign-in
	case r.Method == http.MethodPost && AuthSignInRe.MatchString(r.URL.Path):
		h.signIn(w, r)
		return

	// Sign-up
	case r.Method == http.MethodPost && AuthSignUpRe.MatchString(r.URL.Path):
		h.signUp(w, r)
		return

	// Create actor
	case r.Method == http.MethodPost && ActorRe.MatchString(r.URL.Path):
		h.CreateActor(w, r)
		return

	// Update actor
	case r.Method == http.MethodPatch && ActorReWithID.MatchString(r.URL.Path):
		h.UpdateActorById(w, r)
		return

	// Delete actor
	case r.Method == http.MethodDelete && ActorReWithID.MatchString(r.URL.Path):
		h.DeleteActorById(w, r)
		return

	// Get actors with films
	case r.Method == http.MethodGet && ActorsReWithID.MatchString(r.URL.Path):
		h.GetActorsWithFilms(w, r)
		return

	// Create film
	case r.Method == http.MethodPost && FilmRe.MatchString(r.URL.Path):
		h.CreateFilm(w, r)
		return

	// Update film
	case r.Method == http.MethodPatch && FilmReWithID.MatchString(r.URL.Path):
		h.UpdateFilmById(w, r)
		return

	// Delete film
	case r.Method == http.MethodDelete && FilmReWithID.MatchString(r.URL.Path):
		h.DeleteFilmById(w, r)
		return

	// Get films with sort
	case r.Method == http.MethodGet && FilmsReWithIDAndWithSort.MatchString(r.URL.Path):
		h.GetFilmsWithSort(w, r)
		return

	// Get films with actor name fragment or film name fragment
	case r.Method == http.MethodGet && FilmsReWithIDAndFragment.MatchString(r.URL.Path):
		h.GetFilmsWithFragment(w, r)
		return

	// Get swagger API
	case r.Method == http.MethodGet && SwaggerRe.MatchString(r.URL.Path):
		h.GetSwaggerAPI(w, r)
		return

	// Another Path
	default:
		newErrorResponse(w, http.StatusNotFound, "404 Not Found")
		return
	}
}
