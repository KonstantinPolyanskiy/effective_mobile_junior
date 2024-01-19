package handler

import (
	"effective_mobile_junior/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) PatchPerson(c *gin.Context) {
	var input model.PersonDTO

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   convertType,
			Action: stringToIntAction,
		})
		return
	}

	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   marshallingType,
			Action: err.Error(),
		})
		return
	}

	updatedPerson, err := h.service.ChangePerson(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AppError{
			Type:   updatePersonType,
			Action: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedPerson)
}
