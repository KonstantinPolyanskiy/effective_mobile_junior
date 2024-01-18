package handler

import (
	"context"
	"effective_mobile_junior/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

const (
	savePersonType    = "save person"
	getPersonType     = "get person"
	deletePersonType  = "delete person"
	marshallingType   = "marshalling"
	convertType       = "convert"
	stringToIntAction = "input is not a string"
)

func (h Handler) NewPerson(c *gin.Context) {
	var input model.PostPersonReq

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
	defer cancel()

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   marshallingType,
			Action: err.Error(),
		})
		return
	}

	personInfo, err := h.service.SavePerson(ctx, input)
	if err != nil {
		// Если ошибка слоя бизнес логики вызвана отменой по времени выполнения, устанавливаем соотвествующий http код
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, AppError{
				Type:   savePersonType,
				Action: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, AppError{
			Type:   savePersonType,
			Action: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, personInfo)
}
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
