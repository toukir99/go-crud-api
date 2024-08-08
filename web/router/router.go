package router

import (
	"go-crud-api/web/middlewares"
	"go-crud-api/web/handlers"
	"net/http"
)

func SetupRouter(router *http.ServeMux)  {
    // Routes
    router.Handle("GET /users", middlewares.JSONContentTypeMiddleware(http.HandlerFunc(handlers.GetAllUsers)))
	router.Handle("POST /users", middlewares.JSONContentTypeMiddleware(http.HandlerFunc(handlers.CreateUser)))
	router.Handle("GET /users/", middlewares.JSONContentTypeMiddleware(http.HandlerFunc(handlers.GetUserByID)))
    router.Handle("PUT /users", middlewares.JSONContentTypeMiddleware(http.HandlerFunc(handlers.UpdateUser)))
    router.Handle("DELETE /users", middlewares.JSONContentTypeMiddleware(http.HandlerFunc(handlers.DeleteUser)))
}