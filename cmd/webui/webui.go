package main

import (
	"net/http"

	nexk8s "github.com/akishitara/cronjob-monitor/pkg/client"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl")
	router.Static("/static", "static")
	router.GET("/", func(c *gin.Context) {
		cronJobs := nexk8s.ActiveCronjobsAllGet()
		c.HTML(http.StatusOK, "top.tmpl", gin.H{
			"title":    "CronJobChart",
			"cronJobs": cronJobs,
		})
	})
	router.GET("/tables.js", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tables.js.tmpl", gin.H{
			"data": nexk8s.ActiveCronjobsAllGet(),
		})
	})
	router.GET("/timeline.js", func(c *gin.Context) {
		c.HTML(http.StatusOK, "timeline.js.tmpl", gin.H{
			"data": nexk8s.ActiveCronjobsAllGet(),
		})
	})
	router.GET("/api/cronjobs/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, nexk8s.ActiveCronjobsAllGet())
	})
	router.GET("/api/cronjobs/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, nexk8s.ActiveCronjobsAllGet())
	})
	router.GET("/api/logs/:podName/json", func(c *gin.Context) {
		name := string(c.Param("podName"))
		c.JSON(http.StatusOK, nexk8s.PodLogs(name))
	})
	router.GET("/logs/:podName", func(c *gin.Context) {
		cronJobs := nexk8s.ActiveCronjobsAllGet()
		name := string(c.Param("podName"))
		c.HTML(http.StatusOK, "logs.tmpl", gin.H{
			"title":    "CronJobChart(Log)",
			"cronJobs": cronJobs,
			"log":      nexk8s.PodLogs(name),
			"podName":  name,
		})
	})
	router.Run(":80")
}
