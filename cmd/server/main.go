package main

import (
	"fmt"

	"github.com/mj4code/go-rest-api/internal/db"
)

// Going to be responsible for instantiation
// and startup of go application
func Run() error {
	fmt.Println("Starting the application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to Migrate the database")
		return err
	}

	fmt.Println("Successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go RESTful API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
