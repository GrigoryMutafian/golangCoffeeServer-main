package HandlerModule

import (
	"encoding/json"
	"fmt"
	cm "golangCoffeeServer-main/coffeeModel"
	"net/http"
)

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed: ", http.StatusMethodNotAllowed)
		return
	}

	deleteCoffeesArray := []string{}
	err := json.NewDecoder(r.Body).Decode(&deleteCoffeesArray)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(deleteCoffeesArray) == 0 {
		http.Error(w, "DeleteCoffee list is empty", http.StatusBadRequest)
		return
	}

	for _, value := range deleteCoffeesArray {

		if _, ok := cm.CoffeeDatabase[value]; !ok {
			badDeletingResponse := fmt.Sprintf("Coffee '%s' not found", value)
			http.Error(w, badDeletingResponse, http.StatusBadRequest)
			return
		}
	}

	deletedCoffees := []string{}

	for _, value := range deleteCoffeesArray {
		delete(cm.CoffeeDatabase, value)
		deletedCoffees = append(deletedCoffees, value)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	CurrentMenuList := fmt.Sprintf("Current Menu: %+v", cm.CoffeeDatabase)
	err = json.NewEncoder(w).Encode(CurrentMenuList)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string][]string{"Deleted coffees list": deletedCoffees}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "JSON encoding error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
