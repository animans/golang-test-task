package http

import "net/http"

func Router(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/numbers", h.AddNumber)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	return mux
}
