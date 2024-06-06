package cache

import (
	"errors"
	"log/slog"
	"sync"
	"wb0-app/models"
)

type MemStorage struct {
	data map[string]models.Order
	mu   sync.Mutex
}

type persistantStorage interface {
	FindAll() []models.Order
}

func New() *MemStorage {
	return &MemStorage{data: make(map[string]models.Order)}
}

func (ms *MemStorage) Save(order models.Order) {

	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.data[order.OrderUid] = order
}

func (ms *MemStorage) FindByUid(uid string) (models.Order, error) {
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

func (ms *MemStorage) Restore(ps persistantStorage) {
	orders := ps.FindAll()
	for _, order := range orders {
		ms.Save(order)
	}
}
