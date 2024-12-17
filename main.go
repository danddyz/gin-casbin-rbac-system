package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	e, _ := casbin.NewEnforcer("rbac_model.conf", "policy.csv")
	r.Use(CasbinMiddleware(e))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":9000")
}
