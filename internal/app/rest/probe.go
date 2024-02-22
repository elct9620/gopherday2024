package rest

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var _ Router = &ProbeRouter{}

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
