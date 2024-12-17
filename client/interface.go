package main

import "wb0-app/models"

type Client interface {
	Subscribe(models.Order, ...func(models.Order))
}
