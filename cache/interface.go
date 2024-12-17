package cache

import "wb0-app/models"

type Cache interface {
	Save(order models.Order)
	FindByUid(uid string) (models.Order, error)
	Restore(ps persistantStorage)
}
