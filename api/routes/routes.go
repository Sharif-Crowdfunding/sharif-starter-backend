package routes

import (
	"net/http"
	"sharif-starter-backend/api"
	"sharif-starter-backend/internal/controllers"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	corsMw := mux.CORSMethodMiddleware(r)
	r.Use(CommonMiddleware)

	r.HandleFunc("/", controllers.TestAPI).Methods("GET", "OPTIONS")
	r.HandleFunc("/api", controllers.TestAPI).Methods("GET", "OPTIONS")
	r.HandleFunc("/register", controllers.CreateUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(api.JwtVerify)
	s.HandleFunc("/user", controllers.FetchUsers).Methods("GET", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT", "OPTIONS")
	s.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE", "OPTIONS")

	// Projects
	p := r.PathPrefix("/project").Subrouter()
	p.Use(api.JwtVerify)
	p.HandleFunc("/create", controllers.CreateProject).Methods("POST", "OPTIONS")
	p.HandleFunc("/addTokenInfo", controllers.AddProjectTokenDistribution).Methods("POST", "OPTIONS")
	p.HandleFunc("/list", controllers.GetUserProjects).Methods("GET", "OPTIONS")

	r.Use(corsMw)
	return r
}

// CommonMiddleware --Set content-type
func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
