package app

import "github.com/martinyonathann/bookstore_items-api/controllers"

func mapUrls() {
	router.GET("/items", controllers.GetAllBook)
	router.GET("/items/:items_id", controllers.GetBookById)
	router.POST("items", controllers.CreateBook)
}
