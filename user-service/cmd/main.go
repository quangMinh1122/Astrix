package main
import (
	"log"
	"net/http"
	"github.com/quangMinh1122/user-service/internal/handler"
)
func main() {
	mux := http.NewServeMux()

	// Ensure the mux is not nil before starting it
	if mux == nil {
		log.Fatal("Mux is nil")
	}

	mux.HandleFunc("/health", handler.HealthCheck)
	mux.HandleFunc("/register", handler.RegisterUser)

	startServer(mux)
	
}

func startServer(mux *http.ServeMux) {
	// Ensure the mux is not nil before starting it
	if mux == nil {
		log.Fatal("Mux is nil")
	}

	log.Println("Starting user service on port 8080")
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}	