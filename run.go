package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-restful"
)

type HelloResource struct {
	*gin_restful.Resource
}

func (r HelloResource) Get() (string, int) {
	return "Hello, World!", 404
}

func main() {
	e:=gin.Default()
	api:=gin_restful.NewApi(e, "/")
	api.AddResource(HelloResource{gin_restful.InitResource()}, "/hello")
	e.Run(":8080")
}
