//go:build wireinject
// +build wireinject

package main

import (
	"akmmp241/belajar-golang-restful-api/app"
	"akmmp241/belajar-golang-restful-api/controller"
	"akmmp241/belajar-golang-restful-api/middleware"
	"akmmp241/belajar-golang-restful-api/repository"
	"akmmp241/belajar-golang-restful-api/service"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	//wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	//wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
