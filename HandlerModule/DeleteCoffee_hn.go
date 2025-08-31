package HandlerModule

import (
	"encoding/json"
	"fmt"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Ключ не передан", http.StatusBadRequest)
		return
	}

	if _, ok := cm.CoffeeDatabase[key]; !ok {
		http.Error(w, "Такого кофе нет", http.StatusNotFound)
		return
	}

	delete(cm.CoffeeDatabase, key)
	fmt.Printf("Удалено кофе с ключом '%s', текущее меню: %+v\n", key, cm.CoffeeDatabase)

	w.Header().Set("Content-Type", "application-json")
	w.WriteHeader(http.StatusOK) //
	response := map[string]string{"message": "Кофе удален", "key": key}
	err := json.NewEncoder(w).Encode(response) //
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
	}
}
