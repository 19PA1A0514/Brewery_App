package router

import (
	"brewery/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello,welcome to gin",
		})

	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello, " + name + " welcome to gin !!!",
		})
	})

	router.POST("/user", controller.HandlePostUser)

	router.POST("/user/login", controller.HandleLoginUser)

	router.GET("/user/:id", controller.HandleGetUser)

	router.GET("/user", controller.HandleGetAllUsers)

	router.GET("/user/email/:id", controller.HandleSendEmail)

	router.GET("/user/verifyOtp", controller.HandleVerifyOtp)

	router.DELETE("/user/:id", controller.HandleDeleteUser)

	router.POST("/item/stock/:id", controller.HandleUpdateStock)

	//Order handler functions
	router.POST("/order", controller.HandlePostOrder)
	router.GET("/order/:id", controller.HandleGetOrder)
	router.DELETE("/order/:id", controller.HandleDeleteOrder)
	router.GET("/order", controller.HandleGetAllOrders)
	router.GET("/order/getInvoice/:id", controller.HandleGetInvoice)
	router.GET("/order/user/:id", controller.HandleGetOrderByUser)
	router.GET("/orderbystatus", controller.HandleGetOrderByStatus)
	router.PATCH("/order/updateStatus/:id", controller.HandleUpdateStatus)

	//item handler functions
	router.POST("/item", controller.HandlePostItem)
	router.GET("/item/:id", controller.HandleGetItem)
	router.PUT("/item/:id", controller.HandleUpdateItem)
	router.DELETE("/item/:id", controller.HandleDeleteItem)
	router.GET("/items", controller.HandleGetAllItems)
	router.GET("/item", controller.HandleGetItemByFoodType)
	router.GET("/itemMostOrdered", controller.HandleGetMostOrderedItem)
	router.PATCH("/item/:id", controller.HandlePartialUpdateItem)
	router.PATCH("/item/stock/:id", controller.HandleUpdateStock)

	// r.HandleFunc("/rating", controller.HandlePostRating).Methods("POST")
	// r.HandleFunc("/rating/item/{id}", controller.HandleGetRatingByItem).Methods("GET")
	// r.HandleFunc("/ratings", controller.HandleGetAllratings).Methods("GET")

	return router

}
