package main

import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	DB_HOST= "localhost"
	DB_PORT= 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "1"
	DB_NAME     = "bird_encyclopedia"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	
	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	
	r.HandleFunc("/assets", getBirdIndex).Methods("GET")
	r.HandleFunc("/assets", createBirdPage).Methods("POST")


	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Printf("connection to databased failed")
		fmt.Scanln()
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping to database failed")
		fmt.Scanln()
		panic(err)
	}		
	checkErr(err)
	defer db.Close()
	
	//InitStore(&dbStore{db: db})
	InitStore(db)



	
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}