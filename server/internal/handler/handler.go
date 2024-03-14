package handler

import (
	"filmlib/server/internal/service"
	"net/http"
	"regexp"
)

var (
	AuthSignUpRe             = regexp.MustCompile(`^/auth/sign-up/*$`)
	AuthSignInRe             = regexp.MustCompile(`^/auth/sign-in/*$`)
	ActorRe                  = regexp.MustCompile(`^/actor/*$`)
	ActorReWithID            = regexp.MustCompile(`^/actor/([0-9]+)$`)
	ActorsReWithID           = regexp.MustCompile(`^/actors\?(id=[0-9]+&)*id=[0-9]+$`)
	FilmRe                   = regexp.MustCompile(`^/film/*$`)
	FilmReWithID             = regexp.MustCompile(`^/film/([0-9]+)$`)
	FilmsReWithIDAndFragment = regexp.MustCompile(`^/films\?filmNameFr=[a-zA-Z0-9]+&actorNameFr=[a-zA-Z0-9]+$`)
	FilmsReWithIDAndWithSort = regexp.MustCompile(`^/films\?sortType=[a-zA-Z0-9]+$`)
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/actor", h)

	mux.Handle("/actors/", h)

	mux.Handle("/actors", h)

	mux.Handle("/film", h)

	// Update && Delete film
	mux.Handle("/film/", h)

	mux.Handle("/film", h)

	mux.Handle("/film/search", h)
	return mux
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// Create actor
	case r.Method == http.MethodPost && ActorRe.MatchString(r.URL.Path):
		h.CreateActor(w, r)
	// Update actor
	case r.Method == http.MethodPatch && ActorReWithID.MatchString(r.URL.Path):
		h.UpdateActorById(w, r)
	// Delete actor
	case r.Method == http.MethodDelete && ActorReWithID.MatchString(r.URL.Path):
		h.DeleteActorById(w, r)
	// Get actors with films
	case r.Method == http.MethodGet && ActorsReWithID.MatchString(r.URL.Path):
		h.GetActors(w, r)
	// Create film
	case r.Method == http.MethodPost && FilmRe.MatchString(r.URL.Path):
		h.CreateFilm(w, r)
	// Update film
	case r.Method == http.MethodPatch && FilmReWithID.MatchString(r.URL.Path):
		h.UpdateFilmById(w, r)
	// Delete film
	case r.Method == http.MethodDelete && FilmReWithID.MatchString(r.URL.Path):
		h.DeleteFilmById(w, r)
	// Get films with sort
	case r.Method == http.MethodGet && FilmsReWithIDAndWithSort.MatchString(r.URL.Path):
		h.GetFilmsWithSort(w, r)
	// Get films with actor name fragment or film name fragment
	case r.Method == http.MethodGet && FilmsReWithIDAndFragment.MatchString(r.URL.Path):
		h.GetFilmWithFragment(w, r)
	}

}

//
//func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/actors":
//		switch r.Method {
//		case http.MethodPost:
//			h.CreateActor(w, r)
//		default:
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		}
//	case "/actors/{actor_id}":
//		switch r.Method {
//		case http.MethodPut:
//			h.UpdateActorById(w, r)
//		case http.MethodDelete:
//			h.DeleteActorById(w, r)
//		default:
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		}
//	case "/actors/{actors_id}/films":
//		switch r.Method {
//		case http.MethodGet:
//			h.GetActorFilms(w, r)
//		default:
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		}
//	case "/film":
//		switch r.Method {
//		case http.MethodPost:
//			h.CreateFilm(w, r)
//		default:
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		}
//	case "/films/{film_id}":
//		switch r.Method {
//		case http.MethodPut:
//			h.UpdateFilmById(w, r)
//		case http.MethodDelete:
//			h.DeleteFilmById(w, r)
//		default:
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		}
//	case "/films/search":
//		switch r.Method {
//		case http.MethodGet:
//			h.SearchFilms(w, r)
//		default:
//			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//		}
//	default:
//		http.Error(w, "Not Found", http.StatusNotFound)
//	}
//}
