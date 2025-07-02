package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ahma007/orders/cache"
	"github.com/ahma007/orders/config"
)

func StartServer() error {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	http.HandleFunc("/order", handleGetOrder)

	addr := "0.0.0.0:" + config.AppConfig.ServerPort
	log.Println("🌐 HTTP-сервер запущен на", addr)
	return http.ListenAndServe(addr, nil)
}

func handleGetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	order, found := cache.Get(id)
	if !found {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "failed to encode order", http.StatusInternalServerError)
	}
}
