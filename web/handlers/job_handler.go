package handlers

import (
	"go-db-demo/internal/domain"
	"net/http"

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
