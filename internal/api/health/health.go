package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(app gin.IRouter) {
	app.GET("/readiness", noContentHandler)
	app.GET("/liveness", noContentHandler)
}

func noContentHandler(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
