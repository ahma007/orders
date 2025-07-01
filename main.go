package main

import (
	"github.com/ahma007/orders/api"
	"github.com/ahma007/orders/cache"
	"github.com/ahma007/orders/config"
	"github.com/ahma007/orders/db"
	"github.com/ahma007/orders/kafka"
)

func main() {
	config.Load("config.json") // Загрузка конфига
	db.Init()                  // Инициализация БД
	cache.InitFromDB()         // Кеш
	go kafka.StartConsumer()   // Kafka consumer
	api.StartServer()          // HTTP-сервер
}
