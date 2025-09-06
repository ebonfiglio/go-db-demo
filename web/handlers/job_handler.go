package handlers

import (
	"go-db-demo/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobService domain.JobService
	orgService domain.OrganizationService
	orgs       []domain.Organization
}

func NewJobHandler(jobService domain.JobService, orgService domain.OrganizationService) *JobHandler {
	return &JobHandler{jobService: jobService, orgService: orgService}
}

func (h JobHandler) getOrgs() ([]domain.Organization, error) {
	if h.orgs == nil {
		orgs, err := h.orgService.GetAllOrganizations()
		if err != nil {
			return nil, err
		}
		h.orgs = orgs

	}
	return h.orgs, nil
}

func (h JobHandler) List(c *gin.Context) {
	jobs, err := h.jobService.GetAllJobs()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/list.html", gin.H{
			"Title": "Jobs",
			"Error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "jobs/list.html", gin.H{
		"Title": "Jobs",
		"Jobs":  jobs,
	})
}

func (h *JobHandler) Index(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/list.html", gin.H{
			"Title": "Jobs",
			"Error": err.Error(),
		})
		return
	}

	job, err := h.jobService.GetJob(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/list.html", gin.H{
			"Title": "Jobs",
			"Error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "jobs/index.html", gin.H{
		"Title": "Jobs - " + job.Name,
		"Job":   job,
	})
}

func (h *JobHandler) New(c *gin.Context) {
	orgs, err := h.getOrgs()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/list.html", gin.H{
			"Title": "Jobs",
			"Error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "jobs/new.html", gin.H{
		"Title":         "New Job",
		"Organizations": orgs,
	})
}

func (h *JobHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	orgIdStr := c.PostForm("organizationID")
	orgId, err := strconv.ParseInt(orgIdStr, 10, 64)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/new.html", gin.H{
			"Title":         "Jobs",
			"Error":         err.Error(),
			"Organizations": h.orgs,
		})
		return
	}

	orgs, err := h.getOrgs()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/new.html", gin.H{
			"Title":         "Jobs",
			"Error":         err.Error(),
			"Organizations": []domain.Organization{},
		})
		return
	}

	if name == "" {
		c.HTML(http.StatusBadRequest, "jobs/new.html", gin.H{
			"Title":         "New Job",
			"Error":         "Name is required",
			"Organizations": orgs,
		})
		return
	}

	if orgId == 0 {
		c.HTML(http.StatusBadRequest, "jobs/new.html", gin.H{
			"Title":         "New Job",
			"Error":         "OrganizationId is required",
			"Organizations": orgs,
		})
		return
	}

	job := &domain.Job{Name: name, OrganizationID: orgId}
	_, err = h.jobService.CreateJob(job)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "jobs/new.html", gin.H{
			"Title":         "New Job",
			"Error":         err.Error(),
			"Organizations": orgs,
		})
		return
	}

	c.Redirect(http.StatusSeeOther, "/jobs")
}
