package routers

import (
    "net/http"

    "example-goweb/web"

    "github.com/gin-gonic/gin"
)

// HandleIndex return HTML
func HandleIndex() gin.HandlerFunc {
    return func(c *gin.Context) {
        html := web.MustAsset("index.html")
        c.Data(200, "text/html; charset=UTF-8", html)
    }
}

// Register routes
func Register(e *gin.Engine) {
    h := gin.WrapH(http.FileServer(web.AssetFile()))
    e.GET("/favicon.ico", h)
    e.GET("/js/*filepath", h)
    e.GET("/css/*filepath", h)
    e.GET("/img/*filepath", h)
    e.GET("/fonts/*filepath", h)
    e.NoRoute(HandleIndex())
}