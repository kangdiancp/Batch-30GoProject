package api

import (
	"day13/HTTPServer03/internal/model"
	"day13/HTTPServer03/internal/store"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Query the database
	jsonCategories := store.FindAllCategories()

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCategories)
}

func FindCategoryHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	categoryId, _ := strconv.Atoi(vars["id"])

	// Query the database
	jsonCategory := store.FindCategoryById(categoryId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonCategory)
}

func AddCategoryHandler(w http.ResponseWriter, r *http.Request) {

	// Read the body of the request
	body, err := io.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	var category model.Category

	// Convert the JSON in the request to a Go type
	if err := json.Unmarshal(body, &category); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// call repository AddCategory for insert data
	addResult := store.AddCategory(category)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(addResult)
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {

	// Get URL parameter with the category ID to delete
	vars := mux.Vars(r)
	categoryId, _ := strconv.ParseInt(vars["id"], 10, 64)

	// Query the database
	deleteResult := store.DeleteCategory(categoryId)

	// Format the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(deleteResult)
}
