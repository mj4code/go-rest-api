package main

import (
	"context"
	"fmt"

	"github.com/mj4code/go-rest-api/internal/comment"
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

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"f5e0c237-3fe4-4a19-8372-8d3f9a6b1c58",
	))

	fmt.Println("Successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go RESTful API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
