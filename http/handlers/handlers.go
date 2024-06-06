package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"wb0-app/models"
)

type InCacheOrderFinder interface {
	FindByUid(string) (models.Order, error)
}

func MakeOrderHandler(cache InCacheOrderFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("userInput")

		order, err := cache.FindByUid(id)

		var data []byte
		if err != nil {
			data = []byte("order not found")
		} else {
			data, _ = json.Marshal(order)
		}

		t, _ := template.ParseFiles("http/templates/order.html")
		t.Execute(w, string(data))
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("http/templates/index.html")
	t.Execute(w, nil)
}
