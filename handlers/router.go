package handlers

import (
	"fautest/handlers/auth"
	"fautest/handlers/public"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.New()

	router.Use(auth.CheckAuth)

	router.POST("/signup", public.Signup)
	router.POST("/signin", public.Signin)
	router.POST("/signout", public.Signout)

	return router
}
