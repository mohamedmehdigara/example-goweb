package main

import (
    "example-goweb/routers"

    "github.com/gin-gonic/gin"
)

func main() {
    e := gin.Default()
    routers.Register(e)
    e.Run(":8080")
}