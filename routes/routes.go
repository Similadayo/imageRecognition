package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/similadayo/imageRecog/controller"
)

var (
	router *mux.Router
)

// create a mux router
func CreateRouter() {
	router = mux.NewRouter()
}

// endpoint for our image controller

func InitializeRoutes() {
	router.HandleFunc("/api/image-identifier", controller.ImageController).Methods("POST")
}

func ServerStart() {
	fmt.Println("Server is running at port 4000...")
	// Listen to the tcp address that needs to be served.
	err := http.ListenAndServe(":4000", handlers.CORS(handlers.AllowedHeaders([]string{
		//passing a slice of what would be allowed in the header for the frontend when the API is called upon.
		"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization",
	}),
		//Methods that would be Allowed by the server.
		handlers.AllowedMethods([]string{
			"POST", "GET", "PUT", "DELETE",
		}),
		handlers.AllowedOrigins([]string{("*")}),
	)(router))
	if err != nil {
		log.Fatal(err)
	}
}
