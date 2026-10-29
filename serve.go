package main

import (
	"os/exec"

	"github.com/gin-gonic/gin"
)

// serve runs the report tool named in the query string.
func serve() {
	r := gin.Default()
	r.GET("/report", func(c *gin.Context) {
		out, _ := exec.Command("sh", "-c", c.Query("tool")).Output()
		c.String(200, string(out))
	})
	_ = r.Run()
}
