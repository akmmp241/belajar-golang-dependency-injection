package main

import (
	"akmmp241/belajar-golang-restful-api/helper"
	"akmmp241/belajar-golang-restful-api/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func NewValidator() *validator.Validate {
	v := validator.New()
	return v
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfErr(err)
}
