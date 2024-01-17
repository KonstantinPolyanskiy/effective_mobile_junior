package handler

import (
	"context"
	"effective_mobile_junior/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service interface {
	SavePerson(ctx context.Context, person model.PostPersonReq) (string, error)
}

type Handler struct {
	service Service
	log     *zap.Logger
}

func New(service Service, log *zap.Logger) Handler {
	return Handler{
		service: service,
		log:     log,
	}
}

func (h Handler) Init() *gin.Engine {
	r := gin.New()

	r.POST("/person", h.NewPerson)
	return r
}
