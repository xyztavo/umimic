package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
		AllowedOrigins: config.GetAllowedOrigins(),
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
		message, err := mimic.SendMessage(r.Context(), body.Message, body.History)
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
	r.Post("/short/{userlink}", func(w http.ResponseWriter, r *http.Request) {
		userlink := chi.URLParam(r, "userlink")
		var body struct {
			Redirect string `json:"redirect"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid json", 400)
			return
		}
		if body.Redirect == "" {
			http.Error(w, "redirect required", 400)
			return
		}
		url := fmt.Sprintf("%s/set/%s?EX=86400", config.GetRedisURL(), userlink)
		req, _ := http.NewRequest("POST", url, strings.NewReader(body.Redirect))
		req.Header.Set("Authorization", "Bearer "+config.GetRedisToken())
		resp, err := http.DefaultClient.Do(req)
		if err != nil || resp.StatusCode != 200 {
			http.Error(w, "failed to store", 500)
			return
		}
		resp.Body.Close()
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]string{"status": "created"})
	})
	r.Get("/{userlink}", func(w http.ResponseWriter, r *http.Request) {
		userlink := chi.URLParam(r, "userlink")
		url := fmt.Sprintf("%s/get/%s", config.GetRedisURL(), userlink)
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("Authorization", "Bearer "+config.GetRedisToken())
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, "error", 500)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == 404 {
			http.NotFound(w, r)
			return
		}
		var result struct {
			Result string `json:"result"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			http.Error(w, "error", 500)
			return
		}
		if result.Result == "" {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, result.Result, 302)
	})
	port := ":" + config.GetPort()
	fmt.Printf("\n listening on %s", port)
	http.ListenAndServe(port, r)
}
