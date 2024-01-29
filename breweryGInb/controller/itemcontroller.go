package controller

import (
	"brewery/model"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandlePostItem(c *gin.Context) {

	var item model.Item

	err := c.BindJSON(&item)

	if err != nil {

		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})

		return

	}

	validate := validator.New()

	if err = validate.Struct(item); err != nil {

		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})
		return
	}

	item.InStock = true

	resultItem := handlePostItem(item)

	if resultItem == nil {

		c.JSON(404, gin.H{
			"message": "Failed to create an item",
		})

		return

	}

	c.JSON(200, gin.H{
		"item": resultItem,
	})

}

// handlePostItem handles creating a new Item.
func handlePostItem(item model.Item) *model.Item {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil
	}

	DB.AutoMigrate(&model.Item{})

	if err := DB.Create(&item).Error; err != nil {
		fmt.Println("Error creating Item:", err)
		return nil
	}

	return &item
}

// HandleGetItem handles getting a Item by ID.
func HandleGetItem(c *gin.Context) {

	id := c.Param("id")

	item, err := handleGetItem(id)

	if err != nil {

		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})

		return

	}

	c.JSON(200, gin.H{"item": item})
}

// handleGetItem gets a Item by ID.
func handleGetItem(id string) (*model.Item, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var item model.Item
	err := DB.First(&item, id).Error

	if err != nil {
		return nil, err
	}

	return &item, nil
}

// HandleUpdateItem handles updating a Item.
func HandleUpdateItem(c *gin.Context) {

	var item model.Item

	c.BindJSON(item)

	id := c.Param("id")

	updatedItem, err := handleUpdateItem(item, id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})

		return
	}
	c.JSON(200, gin.H{
		"item": updatedItem,
	})
}

// handleUpdateItem handles updating a Item.
func handleUpdateItem(Item model.Item, id string) (*model.Item, error) {
	var existingItem model.Item
	err := DB.First(&existingItem, id).Error

	if err != nil {
		return nil, err
	}

	existingItem = Item

	err = DB.Save(&existingItem).Error

	if err != nil {
		return nil, err
	}

	return &existingItem, nil
}

func HandlePartialUpdateItem(c *gin.Context) {

	var item model.Item

	c.BindJSON(item)

	id := c.Param("id")

	updatedItem, err := handlePartialUpdateItem(item, id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})

		return
	}
	c.JSON(200, gin.H{
		"item": updatedItem,
	})
}

func handlePartialUpdateItem(item model.Item, id string) (*model.Item, error) {
	var existingItem model.Item

	err := DB.First(&existingItem, id).Error

	if err != nil {
		return nil, err
	}

	if item.Name != "" {
		existingItem.Name = item.Name
	}

	if item.Price != 0 {
		existingItem.Price = item.Price
	}

	if item.Image != "" {
		existingItem.Image = item.Image
	}

	if item.Type != "" {
		existingItem.Type = item.Type
	}

	if item.Details != "" {
		existingItem.Details = item.Details
	}

	err = DB.Save(&existingItem).Error

	if err != nil {
		return nil, err
	}

	return &existingItem, nil
}
func HandleUpdateStock(c *gin.Context) {

	id := c.Param("id")

	instock := c.DefaultQuery("instock", "false")

	stock, _ := strconv.ParseBool(instock)

	updatedItem, err := handleUpdateStock(stock, id)

	if err != nil {

		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})

		return
	}

	c.JSON(200, gin.H{
		"item": updatedItem,
	})

}

func handleUpdateStock(instock bool, id string) (*model.Item, error) {

	var existingItem model.Item

	err := DB.Find(&existingItem, id).Error

	if err != nil {
		return nil, err
	}

	existingItem.InStock = instock

	DB.Save(&existingItem)

	return &existingItem, nil

}

// HandleDeleteItem handles deleting a Item.
func HandleDeleteItem(c *gin.Context) {

	id := c.Param("id")

	err := handleDeleteItem(id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please enter valid input",
		})

		return
	}

	c.JSON(404, gin.H{
		"message": "Item got successfully deleted",
	})

}

// handleDeleteItem handles deleting a Item.
func handleDeleteItem(id string) error {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return fmt.Errorf("DB not initialized")
	}

	err := DB.Delete(&model.Item{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// HandleGetAllItems handles getting all Items.
func HandleGetAllItems(c *gin.Context) {

	items, err := handleGetAllItems()

	if err != nil {
		c.JSON(404, gin.H{
			"message": "Bad Request",
		})

		return

	}

	c.JSON(404, gin.H{
		"items": items,
	})

}

// handleGetAllItems gets all Items.
func handleGetAllItems() (*[]model.Item, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var items []model.Item
	err := DB.Where("in_stock = ?", true).Find(&items).Error

	if err != nil {
		return nil, errors.New("bad request")
	}

	return &items, nil
}

func HandleGetMostOrderedItem(c *gin.Context) {

	items := handleGetMostOrderedItem()

	c.JSON(200, gin.H{
		"items": items,
	})

}

func handleGetMostOrderedItem() []model.Item {

	var items []model.Item

	DB.Order("quantity_ordered DESC").Limit(10).Find(&items)

	return items

}

func HandleGetItemByFoodType(c *gin.Context) {

	foodType := c.Query("foodtype")

	items := handleGetItemByFoodType(foodType)

	c.JSON(200, gin.H{
		"items": items,
	})

	c.JSON(400, gin.H{
		"error": "It's an error",
	})

}

func handleGetItemByFoodType(foodType string) []model.Item {

	var items []model.Item

	DB.Where("Type = ?", foodType).Find(&items)

	return items
}
