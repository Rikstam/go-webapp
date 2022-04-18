package main

import (
	"fmt"
	"net/http"

	"lenslocked.com/controllers"
	"lenslocked.com/models"

	"github.com/gorilla/mux"
)

const (
	host   = "localhost"
	port   = 5433
	user   = "riksa"
	dbname = "lenslocked_dev"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 not found</h1>")
}

func main() {
	// Create a DB connection string and then use it to
	// create our model services.

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"dbname=%s sslmode=disable",
		host, port, user, dbname)

	us, err := models.NewUserService(psqlInfo)

	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)
	galleriesC := controllers.NewGalleries()

	var h http.Handler = http.HandlerFunc(notFound)
	r := mux.NewRouter()
	r.NotFoundHandler = h
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/galleries", galleriesC.New).Methods("GET")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	http.ListenAndServe(":3000", r)
}
