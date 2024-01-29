package controller

import (
	"brewery/model"
	"fmt"
	"time"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandlePostOrder(c *gin.Context) {

	var order model.Order

	err := c.BindJSON(&order)

	if err != nil {
		c.JSON(404, gin.H{
			"message": err.Error(),
		})
		return
	}

	id := order.UserID

	var user model.User

	DB.First(&user, id)

	order.User = user

	var items []model.OrderedItem = order.FoodItems

	var amount float32 = 0.0

	for index, item := range items {

		DB.AutoMigrate(&model.OrderedItem{})

		//sDB.Create(&item)

		var it model.Item

		DB.First(&it, item.ItemID)

		it.QuantityOrdered = it.QuantityOrdered + item.Quantity

		DB.Save(&it)

		item.Item = it

		amount += (it.Price * float32(item.Quantity))

		items[index] = item

		DB.First(&it, item.ItemID)

	}

	order.FoodItems = items

	order.TotalAmount = amount

	var invoice model.Invoice = model.Invoice{Amount: amount, DeliveryFee: 60.0, Gst: amount * 0.05, TotalAmount: amount + 60.0 + amount*0.05}

	DB.Create(&invoice)

	order.Invoice = invoice

	validate := validator.New()

	if err := validate.Struct(order); err != nil {
		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return
	}

	resultOrder := handlePostOrder(order)

	c.JSON(200, gin.H{
		"user": resultOrder,
	})
}

// handlePostUser handles creating a new user.
func handlePostOrder(order model.Order) *model.Order {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil
	}

	DB.AutoMigrate(&model.Order{})

	order.OrderedTime = time.Now()

	order.Status = "Pending"

	if err := DB.Create(&order).Error; err != nil {
		fmt.Println("Error creating user:", err)
		return nil
	}

	return &order
}

// HandleGetUser handles getting a user by ID.
func HandleGetOrder(c *gin.Context) {

	id := c.Param("id")

	order, err := handleGetOrder(id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return
	}

	c.JSON(200, gin.H{
		"order": order,
	})

}

// handleGetUser gets a user by ID.
func handleGetOrder(id string) (*model.Order, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var order model.Order
	err := DB.Preload("User").Preload("Invoice").Find(&order, id).Error

	var fooditems []model.OrderedItem

	DB.Preload("Item").Where("order_id = ?", order.ID).Find(&fooditems)

	order.FoodItems = fooditems

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func HandleUpdateStatus(c *gin.Context) {

	id := c.Param("id")

	status := c.Query("status")

	var stat model.OrderStatus = model.OrderStatus(status)

	switch stat {
	case model.OrderStatusPending, model.OrderStatusConfirmed, model.OrderStatusShipped, model.OrderStatusDelivered:
		fmt.Println("ok")
	default:
		c.JSON(404, gin.H{
			"message": "please provide valid status",
		})
		return
	}

	order, err := handleUpdateStatus(stat, id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return
	}

	c.JSON(200, gin.H{
		"order": order,
	})

}

func handleUpdateStatus(status model.OrderStatus, id string) (*model.Order, error) {

	var order model.Order

	DB.Preload("User").Preload("Item").First(&order, id)

	order.Status = status

	err := DB.Save(&order).Error

	if err != nil {
		return nil, err
	}

	return &order, nil

}

func HandleDeleteOrder(c *gin.Context) {

	id := c.Param("id")

	err := handleDeleteOrder(id)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "There is no such order",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Order got Deleted",
	})
}
func handleDeleteOrder(id string) error {

	err := DB.Unscoped().Delete(&model.Order{}, id)

	if err != nil {
		return errors.New("please provide valid id")
	}

	return nil
}

// HandleGetAllUsers handles getting all users.
func HandleGetAllOrders(c *gin.Context) {

	orders, err := handleGetAllOrders()

	if err != nil {
		c.JSON(404, gin.H{
			"message": "please provide valid input",
		})
		return
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})

}

// handleGetAllUsers gets all users.
func handleGetAllOrders() (*[]model.Order, error) {
	if DB == nil {
		fmt.Println("DB is nil. Make sure it's properly initialized.")
		return nil, fmt.Errorf("DB not initialized")
	}

	var orders []model.Order

	err := DB.Preload("User").Preload("Invoice").Find(&orders).Error

	for index, order := range orders {

		var fooditems []model.OrderedItem

		DB.Preload("Item").Where("order_id = ?", order.ID).Find(&fooditems)

		orders[index].FoodItems = fooditems
	}

	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func HandleGetInvoice(c *gin.Context) {

	id := c.Param("id")

	invoice := handleGetInvoice(id)

	c.JSON(200, gin.H{
		"invoice": invoice,
	})
}

func handleGetInvoice(id string) model.Invoice {

	var invoice model.Invoice

	var order model.Order

	DB.First(&order, id)

	DB.First(&invoice, order.InvoiceID)

	return invoice
}

func HandleGetOrderByUser(c *gin.Context) {

	id := c.Param("id")

	orders := handleGetOrderByUser(id)

	c.JSON(200, gin.H{
		"orders": orders,
	})

}

func handleGetOrderByUser(id string) *[]model.Order {

	var orders []model.Order

	err := DB.Preload("User").Preload("Invoice").Where("user_id = ?", id).Find(&orders).Error

	for index, order := range orders {
		var fooditems []model.OrderedItem
		DB.Preload("Item").Where("order_id = ?", order.ID).Find(&fooditems)
		orders[index].FoodItems = fooditems

	}

	if err != nil {
		fmt.Println(err)
	}

	return &orders

}

func HandleGetOrderByStatus(c *gin.Context) {

	status := c.Query("status")

	orders := handleGetOrdersByStatus(model.OrderStatus(status))

	c.JSON(200, gin.H{
		"orders": orders,
	})
}

func handleGetOrdersByStatus(status model.OrderStatus) []struct {
	Status string
	Count  int
} {
	var groupedOrders []struct {
		Status string
		Count  int
	}

	x := DB.Model(&model.Order{}).Select("status, count(*) as count").Group("status").Scan(&groupedOrders).Error

	if x != nil {
		fmt.Println(x)
	}

	return groupedOrders
}
