package utils

import (
	"InmoGo/src/api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRequest(c *gin.Context, model interface{}) (request models.Request) {
	request.Headers = GetHeaders(c)
	request.QueryParams = GetQueryParams(c)
	request.Body = GetBody(c, model)

	return request
}

func GetHeaders(c *gin.Context) http.Header {
	return c.Request.Header
}

func GetQueryParams(c *gin.Context) (queryParams models.QueryParams) {
	if err := c.BindQuery(&queryParams); err != nil {
		panic(err)
	}
	return queryParams
}

func GetBody(c *gin.Context, model interface{}) interface{} {
	if c.Request.Body == http.NoBody {
		return nil
	}
	if err := c.BindJSON(&model); err != nil {
		panic(err)
	}
	return model
}
