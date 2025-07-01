package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Покажем, откуда читаем index.html
	wd, _ := os.Getwd()
	log.Println("📁 Рабочая директория:", wd)

	// Проверим, что index.html существует
	if _, err := os.Stat("web/index.html"); err != nil {
		log.Fatal("❌ index.html не найден:", err)
	} else {
		log.Println("✅ index.html найден")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("📥 Получен запрос:", r.Method, r.URL.Path)
		http.ServeFile(w, r, "web/index.html")
	})

	log.Println("🌐 Запускаем сервер на :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
