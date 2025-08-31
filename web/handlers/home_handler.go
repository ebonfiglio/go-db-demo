package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) Index(c *gin.Context) {
	// Simple HTML response to avoid template issues
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, `
<!DOCTYPE html>
<html>
<head>
	<title>Go DB Demo</title>
</head>
<body>
	<h1>Welcome to Go DB Demo</h1>
	<p>Management System</p>
	<ul>
		<li><a href="/organizations">Manage Organizations</a></li>
		<li><a href="/jobs">Manage Jobs</a></li>
		<li><a href="/users">Manage Users</a></li>
	</ul>
</body>
</html>`)
}
