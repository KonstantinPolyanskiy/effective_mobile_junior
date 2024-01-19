package handler

import (
	"context"
	"effective_mobile_junior/internal/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
