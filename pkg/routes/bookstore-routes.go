package routes

import (
	"github.com/cruzD21/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookSoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("GET")
	router.HandleFunc("/book/", controllers.getBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}