package main

import (
	"net/http"
	"time"

	v1 "github.com/awahids/belajar-go/internal/delivery/router"
	"github.com/awahids/belajar-go/internal/infrastructure/database"
	"github.com/awahids/belajar-go/internal/infrastructure/database/seeds"
	"github.com/awahids/belajar-go/pkg/helpers"
)

func main() {
	db, _ := database.NewDB()
	r := v1.SetupRouters(db)
	seeds.Seeders(db)

	server := &http.Server{
		Addr:           ":8181",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helpers.ErrorPanic(err)
}
