package config

import (
	"encoding/json"
	"log"
	"os"
)

// Структура, описывающая конфигурацию
type Config struct {
	ServerPort string `json:"server_port"`

	DB struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		SSLMode  string `json:"sslmode"`
	} `json:"db"`

	Kafka struct {
		Brokers []string `json:"brokers"`
		Topic   string   `json:"topic"`
		GroupID string   `json:"group_id"`
	} `json:"kafka"`
}

// Глобальная переменная для доступа к конфигу
var AppConfig Config

// Загрузка конфигурации из JSON-файла
func Load(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("❌ Не удалось открыть конфиг-файл: %v", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("❌ Ошибка парсинга config.json: %v", err)
	}

	log.Println("✅ Конфигурация загружена")
}
