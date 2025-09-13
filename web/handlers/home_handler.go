package handlers

import (
	"fmt"
	"go-db-demo/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
	orgService domain.OrganizationService
}

func NewHomeHandler(orgService domain.OrganizationService) *HomeHandler {
	return &HomeHandler{orgService: orgService}
}

func (h *HomeHandler) Index(c *gin.Context) {
	orgs, err := h.orgService.GetAllOrganizations()
	if err != nil {
		// just render a very simple error on the home page
		c.HTML(http.StatusOK, "home.html", gin.H{
			"Title": "Home",
			"Error": fmt.Sprintf("DB error: %v", err),
		})
		return
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title":         "Home",
		"Organizations": orgs,
	})
}
