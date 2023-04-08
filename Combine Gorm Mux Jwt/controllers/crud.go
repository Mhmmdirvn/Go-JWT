package controllers

import (
	"Combine-Gorm-Mux-Jwt/product"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ReadAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBMarket()
	if err != nil {
		response, _  := json.Marshal(map[string]string{"Message": "not connected to database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	var product []product.Product
	if err = db.Find(&product).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	w.Write(response)
}


func ReadProductById(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBMarket()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Not connected to database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	vars := mux.Vars(r)
	fmt.Println(vars)
	id := vars["id"]

	var product product.Product

	if err = db.First(&product, id).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Data not found"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	response, err := json.Marshal(product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Data cannot be converted to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}else {
		w.Write(response)
	}
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	db, err := product.ConnectDBMarket()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"message": "not connected to database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	var product product.Product

	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "Failed decode to json"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	}

	if err = db.Create(&product).Error; err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "failed add new data"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} else {
		response := map[string]string{"Message": "Success add new data"}
		json.NewEncoder(w).Encode(response)
	}
}


func EditProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := product.ConnectDBMarket()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "cannot connect to database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var product product.Product
	err = json.NewDecoder(r.Body).Decode(&product)


	if db.Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		response, _ := json.Marshal(map[string]string{"Message": "Data cannot be changed"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} else {
		response := map[string]string{"Message": "Update Data Success"}
		json.NewEncoder(w).Encode(response)
	}
}


func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	db, err := product.ConnectDBMarket()
	if err != nil {
		response, _ := json.Marshal(map[string]string{"Message": "cannot connect to database"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	var product product.Product

	if db.Delete(&product, "id = ?", id).RowsAffected == 0 {
		response, _ := json.Marshal(map[string]string{"Message": "Cannot delete data"})
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
		return
	} else {
		response := map[string]string{"Message": "Delete Data Success"}
		json.NewEncoder(w).Encode(response)
	}
}