package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/ahma007/orders/cache"
	"github.com/ahma007/orders/config"
	"github.com/ahma007/orders/db"
	"github.com/ahma007/orders/models"
	"github.com/segmentio/kafka-go"
)

func StartConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: config.AppConfig.Kafka.Brokers,
		Topic:   config.AppConfig.Kafka.Topic,
		GroupID: config.AppConfig.Kafka.GroupID,
	})
	defer r.Close()

	log.Println("🟢 Kafka consumer запущен...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("❗Ошибка чтения из Kafka:", err)
			continue
		}

		var order models.Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Println("❗Ошибка JSON в сообщении Kafka:", err)
			continue
		}

		log.Println("📥 Получен заказ из Kafka:", order.OrderUID)

		// Сохраняем в БД и кеш
		if err := db.SaveOrder(order); err != nil {
			log.Println("❌ Ошибка сохранения заказа в БД:", err)
		} else {
			cache.Set(order)
		}
	}
}
