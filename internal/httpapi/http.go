package httpapi

import (
	"coldchain/internal/domain"
	"coldchain/internal/parser"
	"coldchain/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Handler struct{ S *service.Service }

func New(s *service.Service) *Handler { return &Handler{S: s} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 4 || parts[0] != "v1" || parts[1] != "parks" {
		http.NotFound(w, r)
		return
	}
	park, err := parser.ParsePark(parts[2])
	if err != nil {
		json.NewEncoder(w).Encode([]domain.DispatchOrder{})
		return
	}
	if r.Method == http.MethodGet {
		ds, e := h.S.List(park, r.URL.Query().Get("zone"))
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(ds)
		return
	}
	if r.Method == http.MethodPost {
		var q parser.DispatchRequest
		if e := json.NewDecoder(r.Body).Decode(&q); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		q.Park = park
		d, e := h.S.CreateDispatch(q)
		if errors.Is(e, domain.ErrInvalidPark) {
			http.Error(w, e.Error(), 400)
			return
		}
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(d)
		return
	}
	http.Error(w, "method not allowed", 405)
}
