package Controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	


func RegisterUserAPI(r chi.Router) {
	r.Get("/api/User/v1/GetUser", GetUserDetails)
	Handle("/api/testuser", http.HandlerFunc(TestUser))


nc GetUserDetails(w http.ResponseWriter, r *http.Request) {
/var d []User

}

fnc TestUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("the server is running fine in local host")
}
