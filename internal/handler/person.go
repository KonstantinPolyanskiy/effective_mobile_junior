package handler

import (
	"effective_mobile_junior/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h Handler) NewPerson(c *gin.Context) {
	var person model.PostPersonReq

	c.Request.Context()

	if err := c.BindJSON(&person); err != nil {
		c.String(http.StatusBadRequest, "error unmarshall person")
	}

	personInfo, err := h.service.SavePerson(person)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}

	c.String(http.StatusOK, personInfo)
}
