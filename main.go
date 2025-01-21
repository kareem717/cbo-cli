package main

import (
	"context"

	"github.com/kareem717/k7-cbo/cmd"
	"github.com/kareem717/k7-cbo/internal/service"
	"github.com/kareem717/k7-cbo/internal/storage/sqllite"
)

func main() {
	db, err := sqllite.NewRepository()
	if err != nil {
		panic(err)
	}

	if err := db.Migrate(context.Background()); err != nil {
		panic(err)
	}

	service := service.NewService(db)

	cmd.Execute(context.Background(), service)
}
