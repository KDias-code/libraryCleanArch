package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autharizationHeader = "Autharization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context){
	header := c.GetHeader(autharizationHeader)
	if header == ""{
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return 
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2{
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Autharization.ParseToken(headerParts[1])
	if err != nil{
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error){
	id, ok := c.Get(userCtx)
	if !ok{
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	IdInt , ok := id.(int)
	if !ok{
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	return IdInt, nil
}