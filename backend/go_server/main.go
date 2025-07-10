package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/paula-dot/echoGraph/backend/go_server/handlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zmb3/spotify"
)

// Global variables for Spotify authentication
var (
	auth         spotify.Authenticator
	ch           = make(chan *spotify.Client)
	state        = "echograph-state-token"
	clientID     string
	clientSecret string
	db           *sql.DB
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get Spotify API credentials from environment variables
	clientID = os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret = os.Getenv("SPOTIFY_CLIENT_SECRET")
	redirectURI := os.Getenv("SPOTIFY_REDIRECT_URI")

	// Check if all required environment variables are set
	if clientID == "" || clientSecret == "" || redirectURI == "" {
		log.Fatal("Error: SPOTIFY_CLIENT_ID, SPOTIFY_CLIENT_SECRET, and SPOTIFY_REDIRECT_URI must be set in .env or environment variables")
	}

	// Configure the Spotify authenticator
	auth = spotify.NewAuthenticator(
		redirectURI,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserReadRecentlyPlayed,
		spotify.ScopeUserTopRead,
	)
	auth.SetAuthInfo(clientID, clientSecret)

	// Connect to the database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Error: DATABASE_URL must be set in .env or environment variables")
	}
	var errDb error
	db, errDb = sql.Open("postgres", dbURL)
	if errDb != nil {
		log.Fatal(errDb)
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error: Could not connect to the database: ", err)
	}
	log.Println("Successfully connected to the database!")

	r := mux.NewRouter()

	// Example endpoint
	r.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"pong"}`))
	}).Methods("GET")

	// Example: GET /api/users
	r.HandleFunc("/api/users", GetUsersHandler).Methods("GET")

	// Define HTTP handlers
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/login", loginHandler)

	// Start the HTTP server
	fmt.Println(">> EchoGraph is running on http://localhost:8080 ...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// loginHandler redirects the user to Spotify for authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
	url := auth.AuthURL(state)
	http.Redirect(w, r, url, http.StatusFound)
}

// completeAuth handles the callback from Spotify after authentication
func completeAuth(w http.ResponseWriter, r *http.Request) {
	// Check for state mismatch
	tokState := r.FormValue("state")
	if tokState != state {
		http.Error(w, "State mismatch", http.StatusForbidden)
		return
	}

	// Exchange the authorization code for an access token
	token, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Println(err)
		return
	}

	// Create a new Spotify client
	client := auth.NewClient(token)

	// Start polling for recently played tracks in the background
	go handlers.StartPolling(&client, db)

	// Send the client to the main goroutine
	ch <- &client

	// Redirect to the frontend
	http.Redirect(w, r, "/", http.StatusFound)
}

// Example handler function
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Fetch users from DB and return as JSON
	w.Write([]byte(`[]`))
}
