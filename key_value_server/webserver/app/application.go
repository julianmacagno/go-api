package app

import (
	"github.com/gin-gonic/gin"
)

// Router routes all app requests
var Router *gin.Engine

// StartApplication sets up app runtime
func StartApplication() {
	Router = gin.Default()
	ConfigureMappings()
	Router.Run(":8080")
}
