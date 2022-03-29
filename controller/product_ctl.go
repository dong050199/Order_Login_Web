package controller

import (
	"SQLite_JWT/config"
	"SQLite_JWT/driver"
	models "SQLite_JWT/model/product"
	"SQLite_JWT/repository/repoimpl"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	ProductRepo := repoimpl.NewProductRepo(driver.Mongo.Client.Database(config.DB_NAME))
	data, err := ProductRepo.SelectId(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	ProductRepo := repoimpl.NewProductRepo(driver.Mongo.Client.Database(config.DB_NAME))
	data, err := ProductRepo.Select()
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	ProductRepo := repoimpl.NewProductRepo(driver.Mongo.Client.Database(config.DB_NAME))
	id, _ := strconv.Atoi(params["id"])
	err := ProductRepo.Delete(id)
	json.NewEncoder(w).Encode(err)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	ProductRepo := repoimpl.NewProductRepo(driver.Mongo.Client.Database(config.DB_NAME))
	err := ProductRepo.Insert(product)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applycation/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	ProductRepo := repoimpl.NewProductRepo(driver.Mongo.Client.Database(config.DB_NAME))
	err := ProductRepo.Update(product, id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(product)

}
