package routes

import (
	"github.com/gorilla/mux"
	"bookstore/pkg/controller"
)

var RegisterBookStoreRouter = func(router *mux.Router){

	router.HandleFunc("/book/", controller.Getbook).Method("GET")
	router.HandleFunc("/blook/", controller.CreateBook).Method("POST")
	router.HandleFunc("/book/{bookid}", controller.GetBookById).Method("GET")
	router.HandleFunc("/book/{bookid}", controller.UpdateBook).Method("PUT")
	router.HandleFunc("/book/{bookid}", controller.DeleteBook).Method("DELETE")
}