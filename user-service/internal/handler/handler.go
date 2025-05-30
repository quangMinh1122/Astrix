package handler

import(
	"encoding/json"
	"net/http"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
//Avoid copy requests, use a pointer *http.Request
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req SignUpRequest
	// Decode the JSON request body into the SignUpRequest struct
	// If the request body is not valid JSON, return a 400 Bad Request error
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}	

	// Validate the request data
	if req.Email == "" || req.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// Respond with a success message
	// Use json.NewEncoder to encode the response as JSON
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Respond with a simple health check message
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

