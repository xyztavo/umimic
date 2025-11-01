package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ustav/umimic/config"
	"github.com/ustav/umimic/mimic"
	"github.com/ustav/umimic/models"
	"github.com/ustav/umimic/utils"
)

func main() {
	r := chi.NewRouter()
	r.Get("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		m, _ := json.Marshal(map[string]string{"status": "ok"})
		w.Write([]byte(m))
	})
	r.Post("/api/message", func(w http.ResponseWriter, r *http.Request) {
		body := new(models.Body)
		if err := utils.BindAndValidate(r, body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		message, err := mimic.SendMessage(r.Context(), body.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]string{"reply": message}
		m, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(m)
	})
	port := config.GetPort()
	fmt.Printf("\n listening on %s", port)
	http.ListenAndServe(port, r)
}
