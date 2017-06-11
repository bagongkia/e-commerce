package main

import (
	"./daos"
	"./models"
	"./services"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	db, err := models.OpenDB()
	checkErr(err)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/item-categories", ItemCategories(db))

	panic(http.ListenAndServe(":9090", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to e-commerce api!\n")
}

func ItemCategories(db *models.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		itemCategoryDAO := daos.NewItemCategoryDAO()
		itemCategories, err := services.NewItemCategoryService(itemCategoryDAO).Query(*db)
		checkErr(err)

		itemCategoriesJSON, _ := json.Marshal(itemCategories)
		w.Header().Set("Content-Type", "application/json")
		w.Write(itemCategoriesJSON)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}