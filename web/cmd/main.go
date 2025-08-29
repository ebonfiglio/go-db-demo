package main

import (
	"fmt"
	"go-db-demo/internal/config"
	"go-db-demo/internal/db"
	"go-db-demo/internal/service"
	"go-db-demo/web/handlers"
	"go-db-demo/web/routes"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.Connect()
	defer dbConn.Close()

	orgRepo := db.NewOrganizationRepository(dbConn)
	orgService := service.NewOrganizationService(orgRepo)

	router := gin.Default()

	entityTemplates, _ := filepath.Glob("web/templates/**/*")
	homeTemplates, _ := filepath.Glob("web/templates/*.html")

	allTemplates := append(entityTemplates, homeTemplates...)
	router.LoadHTMLFiles(allTemplates...)

	homeHandler := handlers.NewHomeHandler()
	organizationHandler := handlers.NewOrganizationHandler(orgService)

	routes.SetupHomeRoutes(router, homeHandler)
	routes.SetupOrganizationRoutes(router, organizationHandler)

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", serverAddr)
	router.Run(serverAddr)
}
