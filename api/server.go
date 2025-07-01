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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос: %s %s", r.Method, r.URL.Path)
		fs.ServeHTTP(w, r)
	})

	http.HandleFunc("/order", handleGetOrder)

	log.Println("🌐 HTTP-сервер запущен на :" + config.AppConfig.ServerPort)
	return http.ListenAndServe(":"+config.AppConfig.ServerPort, nil)
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
