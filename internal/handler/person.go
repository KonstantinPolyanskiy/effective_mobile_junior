package handler

import (
	"context"
	"effective_mobile_junior/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	savePersonAction  = "save person"
	marshallingAction = "marshalling"
)

func (h Handler) NewPerson(c *gin.Context) {
	var person model.PostPersonReq

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Microsecond)
	defer cancel()

	if err := c.BindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, AppError{
			Type:   marshallingAction,
			Action: err.Error(),
		})
		return
	}

	personInfo, err := h.service.SavePerson(ctx, person)
	if err != nil {
		// Если ошибка слоя бизнес логики вызвана отменой по времени выполнения, устанавливаем соотвествующий http код
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			c.JSON(http.StatusRequestTimeout, AppError{
				Type:   savePersonAction,
				Action: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, AppError{
			Type:   savePersonAction,
			Action: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, personInfo)
}
