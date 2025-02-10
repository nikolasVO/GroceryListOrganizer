package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"GroceryListOrganizer/internal/service"
)

// OrganizeRequest описывает входной JSON-запрос.
type OrganizeRequest struct {
	Products []string `json:"products"`
}

// OrganizeResponse описывает ответ сервиса.
type OrganizeResponse struct {
	OrganizedProducts []string `json:"organized_products"`
}

// OrganizeHandler принимает POST-запрос с JSON-списком продуктов,
// вызывает сервис для сортировки и возвращает отсортированный список.
func OrganizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req OrganizeRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	organized := service.OrganizeProducts(req.Products)
	resp := OrganizeResponse{OrganizedProducts: organized}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
