package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/VVUx21/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.InitRoutes()
	fmt.Println("Starting the server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server started")
}