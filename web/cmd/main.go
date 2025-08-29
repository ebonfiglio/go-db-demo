package main

import (
	"go-db-demo/internal/db"
	"go-db-demo/internal/service"
	"go-db-demo/web/handlers"
	"go-db-demo/web/routes"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	dbConn := db.Connect()
	defer dbConn.Close()

	orgRepo := db.NewOrganizationRepository(dbConn)
	orgService := service.NewOrganizationService(orgRepo)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/**/*.tmpl")
	entityTemplates, _ := filepath.Glob("web/templates/**/*.tmpl")
	homeTemplates, _ := filepath.Glob("web/templates/*.tmpl")

	allTemplates := append(entityTemplates, homeTemplates...)
	router.LoadHTMLFiles(allTemplates...)

	homeHandler := handlers.NewHomeHandler()
	organizationHandler := handlers.NewOrganizationHandler(orgService)

	routes.SetupHomeRoutes(router, homeHandler)
	routes.SetupOrganizationRoutes(router, organizationHandler)

	router.Run(":8080")
}
