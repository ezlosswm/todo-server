package main

import "log"

func main() {

	// opens connection to db
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	// initializes the database table
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	s := NewAPIServer(3000, store)

	s.Run()

}
