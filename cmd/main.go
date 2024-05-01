package main

import (
	"net/http"
	"time"

	v1 "github.com/awahids/belajar-gin/api/v1"
	"github.com/awahids/belajar-gin/pkg/helpers"
)

func main() {
	r := v1.SetupRouters()

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
