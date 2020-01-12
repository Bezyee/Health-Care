package main
import(
  "net/http"
	"log"
	"github.com/gorilla/mux"

	"github.com/abdissabayou/Online-ticket-store"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handler.AllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users", handler.CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", handler.UpdateUserEndPoint).Methods("PUT")
	r.HandleFunc("/users", handler.DeleteUserEndPoint).Methods("DELETE")
	r.HandleFunc("/users/{id}", handler.FindUserEndpoint).Methods("GET")
	
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
