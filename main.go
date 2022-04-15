package main

import (
	//utilities "ExpenseTracker/Utilities"
	service "ExpenseTracker/Services"
	controller "Expensetracker/Controllers"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	port := flag.String("port", "9005", "Service active Port number, default: 9999")
	fmt.Println("Hello world !")
	// utilities.ConnectDB()
	fmt.Println("Application running on port", *port)

	r := chi.NewRouter()

	//Registering all the routes for the applications
	r.Route("/Expense", func(r chi.Router) {
		controller.RegisterUserAPI(r)
	})
	r.Route("/", func(r chi.Router) {
		service.RegisterGoogleAPIs(r)
	})
	log.Println(http.ListenAndServe(":"+*port, r))
}
