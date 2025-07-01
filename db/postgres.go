package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ahma007/orders/config"
	"github.com/ahma007/orders/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.AppConfig.DB.Host,
		config.AppConfig.DB.Port,
		config.AppConfig.DB.User,
		config.AppConfig.DB.Password,
		config.AppConfig.DB.Name,
		config.AppConfig.DB.SSLMode,
	)

	fmt.Printf("🔧 connStr = %q\n", connStr)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("❌ Ошибка подключения к базе:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("❌ База недоступна:", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			order_uid TEXT PRIMARY KEY,
			data JSONB
		)
	`)
	if err != nil {
		log.Fatal("❌ Ошибка при создании таблицы:", err)
	}

	fmt.Println("✅ База данных подключена и готова")
}
func SaveOrder(order models.Order) error {
	jsonData, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("ошибка сериализации заказа: %w", err)
	}

	_, err = DB.Exec(`
		INSERT INTO orders (order_uid, data)
		VALUES ($1, $2)
		ON CONFLICT (order_uid) DO UPDATE SET data = EXCLUDED.data
	`, order.OrderUID, jsonData)

	return err
}
func GetOrderByID(id string) (models.Order, bool) {
	var raw []byte
	err := DB.QueryRow(`SELECT data FROM orders WHERE order_uid = $1`, id).Scan(&raw)
	if err != nil {
		return models.Order{}, false
	}

	var order models.Order
	err = json.Unmarshal(raw, &order)
	if err != nil {
		log.Println("❗Ошибка парсинга JSON из БД:", err)
		return models.Order{}, false
	}

	return order, true
}
func GetAllOrders() []models.Order {
	rows, err := DB.Query(`SELECT data FROM orders`)
	if err != nil {
		log.Println("❗Ошибка выборки всех заказов:", err)
		return nil
	}
	defer rows.Close()

	var all []models.Order

	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			continue
		}

		var order models.Order
		if err := json.Unmarshal(raw, &order); err != nil {
			continue
		}

		all = append(all, order)
	}

	return all
}
