package store

import (
	"day13/HTTPServer03/internal/model"
	"encoding/json"
	"fmt"
	"log"
)

// Find all category and return as JSON
func FindAllCategories() []byte {
	//1. call model
	var categories model.Categories
	var category model.Category

	// melakukan koneksi ke db
	CategoryResults, err := database.Query("select category_id, category_name from categories")
	if err != nil {
		log.Fatal(err)
	}
	defer CategoryResults.Close()

	for CategoryResults.Next() {
		CategoryResults.Scan(&category.CategoryId, &category.CategoryName)
		categories = append(categories, category)
	}

	jsonCategories, err := json.Marshal(categories)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonCategories
}

// Find a single category based on ID and return as JSON
func FindCategoryById(id int) []byte {
	var category model.Category

	//fetch one record
	err := database.QueryRow(`
	select category_id, category_name
	from categories where category_id =$1`, id).Scan(&category.CategoryId, &category.CategoryName)
	if err != nil {
		log.Fatal(err)
	}

	jsonCategory, err := json.Marshal(category)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return jsonCategory
}

func AddCategory(category model.Category) []byte {

	var addResult model.ResponseMsg

	// Create prepared statement
	stmt, err := database.Prepare(`
	INSERT INTO categories(category_name) VALUES($1)`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(category.CategoryName)
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected

	addResult.Affected = rowCnt

	// Convert to JSON and return
	newCategory, err := json.Marshal(addResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return newCategory
}

func DeleteCategory(id int64) []byte {
	var deleteResult model.ResponseMsg

	// Create prepared statement
	stmt, err := database.Prepare(`DELETE FROM categories WHERE category_id=$1`)
	if err != nil {
		log.Fatal(err)
	}

	// Execute the prepared statement and retrieve the results
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Populate DBUpdate struct with last Id and num rows affected
	deleteResult.Id = id
	deleteResult.Affected = rowCnt

	// Convert to JSON and return
	deletedCategory, err := json.Marshal(deleteResult)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return deletedCategory
}
