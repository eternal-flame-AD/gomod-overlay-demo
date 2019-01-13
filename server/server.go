package main

import (
	"plugin"

	"github.com/gin-gonic/gin"
)

func main() {
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	pluginHandler, err := p.Lookup("Handler")
	if err != nil {
		panic(err)
	}

	g := gin.Default()

	g.GET("/server", func(c *gin.Context) {
		c.Writer.WriteString("hello from server")
	})
	g.GET("/client", *pluginHandler.(*gin.HandlerFunc))

	g.Run(":10080")
}
