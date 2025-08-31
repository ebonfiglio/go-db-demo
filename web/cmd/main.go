package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"go-db-demo/internal/config"
	"go-db-demo/internal/db"
	"go-db-demo/internal/service"
	"go-db-demo/web"
	"go-db-demo/web/handlers"
	"go-db-demo/web/routes"
)

func main() {
	cfg := config.LoadConfig()

	dbConn := db.Connect()
	defer dbConn.Close()

	orgRepo := db.NewOrganizationRepository(dbConn)
	orgService := service.NewOrganizationService(orgRepo)

	router := gin.Default()

	// Load HTML templates with error handling
	htmlTemplate := web.Parse()
	if htmlTemplate != nil {
		router.SetHTMLTemplate(htmlTemplate)
		log.Println("Templates loaded successfully")
	} else {
		log.Println("Warning: No templates loaded")
	}

	// Health check endpoint - this was missing the HTTP 200 response
	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "ok")
	})

	homeHandler := handlers.NewHomeHandler()
	organizationHandler := handlers.NewOrganizationHandler(orgService)

	routes.SetupHomeRoutes(router, homeHandler)
	routes.SetupOrganizationRoutes(router, organizationHandler)

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", serverAddr)
	log.Printf("Server starting on %s", serverAddr)
	_ = router.Run(serverAddr)
}
