package app

import "github.com/martinyonathann/bookstore_items-api/controllers"

func mapUrls() {
	router.GET("/items", controllers.GetAllBook)
}
