package app

import "github.com/bookstore_items-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
