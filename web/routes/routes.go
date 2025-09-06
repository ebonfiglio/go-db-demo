package routes

import (
	"go-db-demo/web/handlers"

	"github.com/gin-gonic/gin"
)

func SetupHomeRoutes(router *gin.Engine, homeHandler *handlers.HomeHandler) {
	router.GET("/", homeHandler.Index)
}

func SetupOrganizationRoutes(router *gin.Engine, organizationHandler *handlers.OrganizationHandler) {
	orgGroup := router.Group("/organizations")
	{
		orgGroup.GET("", organizationHandler.List)
		orgGroup.GET(":id", organizationHandler.Index)
		orgGroup.GET("/new", organizationHandler.New)
		orgGroup.POST("", organizationHandler.Create)
		orgGroup.GET("/:id/edit", organizationHandler.Edit)
		orgGroup.POST("/:id", organizationHandler.Update)
	}
}

func SetupJobRoutes(router *gin.Engine, jobHandler *handlers.JobHandler) {
	jobGroup := router.Group("/jobs")
	{
		jobGroup.GET("", jobHandler.List)
		jobGroup.GET(":id", jobHandler.Index)
		jobGroup.GET("/new", jobHandler.New)
		jobGroup.POST("", jobHandler.Create)
	}
}
