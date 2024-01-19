package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) DeletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   convertType,
			Action: stringToIntAction,
		})
		return
	}

	isDeleted, err := h.service.DeletePerson(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AppError{
			Type:   deletePersonType,
			Action: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"is_deleted": isDeleted,
	})
}
