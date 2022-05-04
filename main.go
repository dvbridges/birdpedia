// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)

	"database/sql"
	"fmt"

	// The "net/http" library has methods to implement HTTP clients and servers
	"net/http"

	"github.com/gorilla/mux"

	// lib/pg is required to load the database drivers used in database/sql
	_ "github.com/lib/pq"
)

func main() {
	// connString := "dbname=bird_encyclopedia sslmode=disable host=0.0.0.0 port=5432 user=user password=password"
	connString := "postgres://user:password@postgresql/bird_encyclopedia?sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()
	http.ListenAndServe(":8080", r)
}

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	// Declare a new router
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
