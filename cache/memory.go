package cache

import (
	"fmt"
	"sync"

	"github.com/ahma007/orders/db"
	"github.com/ahma007/orders/models"
)

var (
	mu     sync.RWMutex
	orders = make(map[string]models.Order)
)

// InitFromDB загружает все заказы из БД в кеш
func InitFromDB() error {
	all := db.GetAllOrders()
	mu.Lock()
	defer mu.Unlock()

	for _, o := range all {
		orders[o.OrderUID] = o
	}

	fmt.Println("🧠 Кэш инициализирован")
	return nil
}

// Get возвращает заказ из кеша по UID
func Get(id string) (models.Order, bool) {
	mu.RLock()
	defer mu.RUnlock()

	order, ok := orders[id]
	return order, ok
}

// Set добавляет или обновляет заказ в кеше
func Set(order models.Order) {
	mu.Lock()
	defer mu.Unlock()

	orders[order.OrderUID] = order
}
