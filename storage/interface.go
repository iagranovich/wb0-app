package storage

type Storage interface {
	Save(models.Order)
	FindAll() []mosels.Order
}
