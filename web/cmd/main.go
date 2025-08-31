package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-db-demo/internal/config"
	"go-db-demo/internal/db"
	"go-db-demo/internal/service"
	"go-db-demo/web/handlers"
	"go-db-demo/web/routes"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.Connect()
	if dbConn != nil {
		defer dbConn.Close()
	}

	var orgService *service.OrganizationService
	if dbConn != nil {
		orgRepo := db.NewOrganizationRepository(dbConn)
		orgService = service.NewOrganizationService(orgRepo)
	}

	router := gin.Default()

	// Health check endpoint - ensure this is working
	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// Simple debug endpoint
	router.GET("/debug", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "running",
			"server": cfg.Server,
		})
	})

	// Comment out template loading for now to avoid panics
	// htmlTemplate := web.Parse()
	// if htmlTemplate != nil {
	// 	router.SetHTMLTemplate(htmlTemplate)
	// 	log.Println("Templates loaded successfully")
	// } else {
	// 	log.Println("Warning: No templates loaded")
	// }

	homeHandler := handlers.NewHomeHandler()

	routes.SetupHomeRoutes(router, homeHandler)

	// Only setup organization routes if we have a database connection
	if orgService != nil {
		organizationHandler := handlers.NewOrganizationHandler(orgService)
		routes.SetupOrganizationRoutes(router, organizationHandler)
	} else {
		// Add a simple handler for organization routes when DB is not available
		router.GET("/organizations", func(c *gin.Context) {
			c.String(http.StatusServiceUnavailable, "Database not available")
		})
	}

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", serverAddr)
	log.Printf("Server starting on %s", serverAddr)
	_ = router.Run(serverAddr)
}
