package cache

import (
	"errors"
	"log/slog"
	"sync"
	"wb0-app/models"
)

type memStorage struct {
	data map[string]models.Order
	mu   sync.Mutex
}

func New() *memStorage {
	return &memStorage{data: make(map[string]models.Order)}
}

func (ms *memStorage) Save(order models.Order) {

	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.data[order.OrderUid] = order
}

func (ms *memStorage) FindByUid(uid string) (models.Order, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	order, ok := ms.data[uid]
	if !ok {
		msg := "cache: cannot find order with uid=" + uid
		slog.Error(msg)
		return models.Order{}, errors.New(msg)
	}
	return order, nil
}