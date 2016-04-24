package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	ResponseCode int
	ResponseData map[string]interface{}
}


func (h *Handler) ReturnResponse(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Language", "*")
	c.JSON(h.ResponseCode, h.ResponseData)
}

func (h *Handler) Ping(c *gin.Context) {
	h.ResponseCode = http.StatusOK
	h.ResponseData = make(map[string]interface{})
	h.ResponseData["response"] = "pong"
	h.ResponseData["error"] = nil
	h.ReturnResponse(c)
}