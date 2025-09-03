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

	deleteCoffeesArray := []string{}
	err := json.NewDecoder(r.Body).Decode(&deleteCoffeesArray)
	if err != nil {
		http.Error(w, "Некоректный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(deleteCoffeesArray) == 0 {
		http.Error(w, "Список кофе для удаления пуст", http.StatusBadRequest)
		return
	}

	for _, value := range deleteCoffeesArray {

		if _, ok := cm.CoffeeDatabase[value]; !ok {
			badDeletingResponse := fmt.Sprintf("Кофе '%s' не найден", value)
			http.Error(w, badDeletingResponse, http.StatusBadRequest)
			return
		}
	}

	deletedCoffees := []string{}

	for _, value := range deleteCoffeesArray {
		delete(cm.CoffeeDatabase, value)
		deletedCoffees = append(deletedCoffees, value)
	}

	fmt.Printf("Текущее меню: %+v\n", cm.CoffeeDatabase)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string][]string{"Список удалённых кофе": deletedCoffees}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Ошибка кодирования JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
