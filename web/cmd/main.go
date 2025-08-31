package main

import (
	"fmt"

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

	router.SetHTMLTemplate(web.Parse())

	router.GET("/healthz", func(c *gin.Context) { c.String(200, "ok") })

	homeHandler := handlers.NewHomeHandler()
	organizationHandler := handlers.NewOrganizationHandler(orgService)

	routes.SetupHomeRoutes(router, homeHandler)
	routes.SetupOrganizationRoutes(router, organizationHandler)

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", serverAddr)
	_ = router.Run(serverAddr)
}
