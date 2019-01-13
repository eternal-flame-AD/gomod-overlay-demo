package main

import (
	"github.com/gin-gonic/gin"
)

var Handler gin.HandlerFunc = func(c *gin.Context) {
	c.Writer.WriteString("hello from plugin")
}

func main() {
	panic("this is a go plugin")
}
