package controller

import (
	"brewery/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlePostRating(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rating model.Rating

	if err := json.NewDecoder(r.Body).Decode(&rating); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user model.User

	var item model.Item

	DB.Find(&user, rating.UserID)

	DB.Find(&item, rating.ItemID)

	item.Rating = (item.Rating + rating.Rating) / 2.0

	DB.Save(&item)

	rating.Item = item

	rating.User = user

	resultItem := handlePostRating(rating)

	json.NewEncoder(w).Encode(resultItem)
}

// handlePostItem handles creating a new Item.
func handlePostRating(rating model.Rating) *model.Rating {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil
	}

	DB.AutoMigrate(&model.Rating{})

	if err := DB.Create(&rating).Error; err != nil {
		fmt.Println("Error creating Item:", err)
		return nil
	}

	return &rating
}

// HandleGetItem handles getting a Item by ID.
func HandleGetRatingByItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	ratings, err := handleGetRatingByItem(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ApiError{Message: "please provide valid input", StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(ratings)
}

// handleGetItem gets a Item by ID.
func handleGetRatingByItem(id string) ([]model.Rating, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var ratings []model.Rating
	result := DB.Preload("User").Preload("Item").Where("item_id = ?", id).Find(&ratings)
	if result.Error != nil {
		return nil, result.Error
	}
	return ratings, nil
}

func HandleGetAllratings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := handleGetAllRatings()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ApiError{Message: err.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	json.NewEncoder(w).Encode(items)
}

// handleGetAllItems gets all Items.
func handleGetAllRatings() ([]model.Rating, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var ratings []model.Rating
	err := DB.Preload("User").Preload("Item").Find(&ratings).Error

	if err != nil {
		return nil, errors.New("bad request")
	}

	return ratings, nil
}
