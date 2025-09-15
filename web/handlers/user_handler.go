package handlers

import (
	"go-db-demo/internal/domain"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService domain.UserService
	jobService  domain.JobService
	orgService  domain.OrganizationService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService domain.UserService, jobService domain.JobService, orgService domain.OrganizationService) *UserHandler {
	return &UserHandler{
		userService: userService,
		jobService:  jobService,
		orgService:  orgService,
	}
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		h.renderError(c, "users/list.html", "Failed to load users", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "users/list.html", gin.H{
		"Title": "Users",
		"Users": users,
	})
}

func (h *UserHandler) New(c *gin.Context) {
	// TODO: User concurrency ehre
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		log.Printf("Error fetching organizations: %v", err)
		h.renderError(c, "users/list.html", "Failed to load organizations", http.StatusInternalServerError)
		return
	}

	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		log.Printf("Error fetching jobs: %v", err)
		h.renderError(c, "users/list.html", "Failed to load jobs", http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "users/new.html", gin.H{
		"Title":         "New User",
		"Organizations": organizations,
		"Jobs":          jobs,
	})
}

func (h *UserHandler) Create(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	orgIDStr := c.PostForm("organizationID")
	jobIDStr := c.PostForm("jobID")

	if name == "" {
		h.renderCreateFormWithError(c, "Name is required")
		return
	}

	if orgIDStr == "" {
		h.renderCreateFormWithError(c, "Organization is required")
		return
	}

	if jobIDStr == "" {
		h.renderCreateFormWithError(c, "Organization is required")
		return
	}

	orgID, err := strconv.ParseInt(orgIDStr, 10, 64)
	if err != nil {
		h.renderCreateFormWithError(c, "Invalid organization ID")
		return
	}

	jobID, err := strconv.ParseInt(jobIDStr, 10, 64)
	if err != nil {
		h.renderCreateFormWithError(c, "Invalid Job ID")
		return
	}

	user := &domain.User{
		Name:           name,
		OrganizationID: orgID,
		JobID:          jobID,
	}

	_, err = h.userService.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		h.renderCreateFormWithError(c, "Failed to create user")
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (h *UserHandler) renderError(c *gin.Context, template string, message string, statusCode int) {
	c.HTML(statusCode, template, gin.H{
		"Title": "Error",
		"Error": message,
	})
}

func (h *UserHandler) renderCreateFormWithError(c *gin.Context, errorMessage string) {
	organizations, err := h.orgService.GetAllOrganizations()
	if err != nil {
		log.Printf("Error fetching organizations for error page: %v", err)
		organizations = []domain.Organization{}
	}

	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		log.Printf("Error fetching jobs for error page: %v", err)
		jobs = []domain.Job{}
	}

	c.HTML(http.StatusBadRequest, "users/new.html", gin.H{
		"Title":         "New User",
		"Error":         errorMessage,
		"Organizations": organizations,
		"Jobs":          jobs,
		// Preserve form data
		"FormData": gin.H{
			"Name":           c.PostForm("name"),
			"OrganizationID": c.PostForm("organizationID"),
			"JobID":          c.PostForm("jobID"),
		},
	})
}
