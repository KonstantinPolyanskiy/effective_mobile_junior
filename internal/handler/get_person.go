package handler

import (
	"effective_mobile_junior/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h Handler) GetPerson(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   convertType,
			Action: stringToIntAction,
		})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   convertType,
			Action: stringToIntAction,
		})
		return
	}

	older, err := strconv.Atoi(c.DefaultQuery("older", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   convertType,
			Action: stringToIntAction,
		})
		return
	}

	countryFilter := c.Query("country")
	name := c.Query("name")
	genderType := c.Query("gender")

	input := model.GetPersonReq{
		Limit:         limit,
		Offset:        offset,
		GenderType:    genderType,
		Older:         older,
		CountryFilter: countryFilter,
		Name:          name,
	}

	persons, err := h.service.GetPerson(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, AppError{
			Type:   getPersonType,
			Action: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, persons)
}
