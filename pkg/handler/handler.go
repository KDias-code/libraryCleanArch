package handler

import (
	"github.com/KDias-code/todoapp/pkg/service"
	"github.com/gin-gonic/gin"
	// "github.com/go-delve/delve/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	member := router.Group("/auth")
	{
		member.POST("/sign-up", h.SignUp)
		member.POST("/sign-in", h.SignIn)
		member.GET("/", h.getAllMembers)
		member.PUT("/", h.UpdateMember)
		member.DELETE("/", h.DeleteMember)
	}

	api := router.Group("/api", h.userIdentity)
	{
		authors := api.Group("/author")
		{
			authors.POST("/", h.createAuthor)
			authors.GET("/", h.getAllAuthors)
			authors.GET("/:id", h.getAuthorById)
			authors.PUT("/:id", h.UpdateAuthor)
			authors.DELETE("/:id", h.DeleteAuthor)

			books := authors.Group(":id/books")
			{
				books.POST("/", h.createBook)
				books.GET("/", h.getAllBooks)

			}
		}
		books := api.Group("books")
		{
			books.GET("/:id", h.getBookById)
			books.PUT("/:id", h.UpdateBook)
			books.DELETE("/:id", h.DeleteBook)
		}
	}

	return router
}
