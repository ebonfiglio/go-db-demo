package handlers

import (
	"net/http"
	"strconv"

	"go-db-demo/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	orgService domain.OrganizationService
}

func NewOrganizationHandler(orgService domain.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{orgService: orgService}
}

func (h *OrganizationHandler) Index(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid organization ID: " + err.Error(),
		})
		return
	}

	org, err := h.orgService.GetOrganization(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get organization: " + err.Error(),
			"debug": "Check database connection and migrations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organization": org,
	})
}

func (h *OrganizationHandler) List(c *gin.Context) {
	orgs, err := h.orgService.GetAllOrganizations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get organizations: " + err.Error(),
			"debug": "Check database connection and migrations",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"organizations": orgs,
		"count":         len(orgs),
	})
}

func (h *OrganizationHandler) New(c *gin.Context) {
	c.HTML(http.StatusOK, "organizations/new.html", gin.H{
		"Title": "New Organization",
	})
}

func (h *OrganizationHandler) Create(c *gin.Context) {
	name := c.PostForm("name")

	if name == "" {
		c.HTML(http.StatusBadRequest, "organizations/new.html", gin.H{
			"Title": "New Organization",
			"Error": "Name is required",
		})
		return
	}

	org := &domain.Organization{Name: name}
	_, err := h.orgService.CreateOrganization(org)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "organizations/new.html", gin.H{
			"Title": "New Organization",
			"Error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/organizations")
}

func (h *OrganizationHandler) Edit(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "organizations/edit.html", gin.H{
			"Title": "Edit Organization",
			"Error": "Invalid ID",
		})
		return
	}

	org, err := h.orgService.GetOrganization(id)
	if err != nil {
		c.HTML(http.StatusNotFound, "organizations/edit.html", gin.H{
			"Title": "Edit Organization",
			"Error": "Organization not found",
		})
		return
	}

	c.HTML(http.StatusOK, "organizations/edit.html", gin.H{
		"Title":        "Edit Organization",
		"Organization": org,
	})
}

func (h *OrganizationHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "organizations/edit.html", gin.H{
			"Title": "Edit Organization",
			"Error": "Invalid ID",
		})
		return
	}

	name := c.PostForm("name")
	if name == "" {
		org, _ := h.orgService.GetOrganization(id)
		c.HTML(http.StatusBadRequest, "organizations/edit.html", gin.H{
			"Title":        "Edit Organization",
			"Organization": org,
			"Error":        "Name is required",
		})
		return
	}

	org := &domain.Organization{ID: id, Name: name}
	_, err = h.orgService.UpdateOrganization(org)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "organizations/edit.html", gin.H{
			"Title":        "Edit Organization",
			"Organization": org,
			"Error":        err.Error(),
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/organizations")
}
