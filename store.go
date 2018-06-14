package main

import (
	"database/sql"
	"fmt"
)

// Our store will have two methods, to add a new bird,
// and to get all existing birds
// Each method returns an error, in case something goes wrong
/*
type Store interface {
	CreateBird(bird *Bird) error
	GetBirds() ([]*Bird, error)
}*/

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
type dbStore struct {
	db *sql.DB
}
var store dbStore

func InitStore(s *sql.DB) {
	//store = s
	store=dbStore{s}
}
func (store *dbStore) CreateDog(dog *Dog) error {
	var lastInsertId int
	err := store.db.QueryRow("insert into dogs(name) values($1) returning id", dog.Name).Scan(&lastInsertId)
	for i,v:= range dog.Vacinnes{
			_,err=store.db.Query("UPDATE dogs set vacinnes[$1]=$2 where id=$3",i+1,v,lastInsertId)
			checkErr(err)
	}
	return err
}

func (store *dbStore) GetDogs() ([]*Dog, error) {
	fmt.Printf("\nGetting Birds'\n")

	rows, err := store.db.Query("SELECT name, vacinnes from dogs")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {	return nil, err	}
	defer rows.Close()


	dogs := []*Dog{}
	for rows.Next() {
		dog := &Dog{}
		if err := rows.Scan(&dog.Name, &dog.Vacinnes); err != nil {return nil, err}
		dogs = append(dogs, dog)
	}

	return dogs, nil
}


func (store *dbStore) CreateBird(bird *Bird) error {
	// 'Bird' is a simple struct which has "species" and "description" attributes
	// THe first underscore means that we don't care about what's returned from
	// this insert query. We just want to know if it was inserted correctly,
	// and the error will be populated if it wasn't
	_, err := store.db.Query("INSERT INTO birds(bird, description) VALUES ($1,$2)", bird.Species, bird.Description)
	return err
}

func (store *dbStore) GetBirds() ([]*Bird, error) {
	fmt.Printf("\nGetting Birds'\n")
	// Query the database for all birds, and return the result to the
	// `rows` object
	rows, err := store.db.Query("SELECT bird, description from birds")
	// We return incase of an error, and defer the closing of the row structure
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Create the data structure that is returned from the function.
	// By default, this will be an empty array of birds
	birds := []*Bird{}
	for rows.Next() {
		// For each row returned by the table, create a pointer to a bird,
		bird := &Bird{}
		// Populate the `Species` and `Description` attributes of the bird,
		// and return incase of an error
		if err := rows.Scan(&bird.Species, &bird.Description); err != nil {
			return nil, err
		}
		// Finally, append the result to the returned array, and repeat for
		// the next row
		birds = append(birds, bird)
	}
	return birds, nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
//var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
