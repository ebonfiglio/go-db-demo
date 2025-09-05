package handlers

import (
	"go-db-demo/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	jobService domain.JobService
}

func NewJobHandler(jobService domain.JobService) *JobHandler {
	return &JobHandler{jobService: jobService}
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
