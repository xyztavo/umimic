package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/ustav/umimic/config"
	"github.com/ustav/umimic/mimic"
	"github.com/ustav/umimic/models"
	"github.com/ustav/umimic/utils"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{config.GetAllowedOrigin()},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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
