package models

import "net/http"

type Request struct {
	Headers http.Header
	QueryParams QueryParams
	Body interface{}
}

type QueryParams struct {
	Id string `form:"id"`
}