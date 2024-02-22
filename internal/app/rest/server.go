package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Namespace() string
	chi.Router
}

func New(routers ...Router) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	for _, router := range routers {
		r.Mount(router.Namespace(), router)
	}

	return r
}

type ProbeRouter struct {
	*chi.Mux
}

func NewProbeRouter() *ProbeRouter {
	r := chi.NewRouter()

	r.Get("/livez", LivenessHandler)

	return &ProbeRouter{r}
}

func (p *ProbeRouter) Namespace() string {
	return "/"
}

func LivenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "{\"status\": \"ok\"}")
}
