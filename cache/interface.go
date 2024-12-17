package cache

import "wb0-app/models"

type Cache interface {
	Save(models.Order)
	FindByUid(string) (models.Order, error)
	Restore(persistantStorage)
}
