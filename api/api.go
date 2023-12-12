package api

import "github.com/gin-gonic/gin"

var Router *gin.Engine
var V1 *gin.RouterGroup

func CreateRouter() {
	Router = gin.Default()
	V1 = Router.Group("/api/v1")
}
