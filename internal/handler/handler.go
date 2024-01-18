package handler

import (
	"context"
	"effective_mobile_junior/internal/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service interface {
	SavePerson(ctx context.Context, person model.PostPersonReq) (model.PersonEntity, error)
	GetPerson(params model.GetPersonReq) ([]model.PersonEntity, error)
	DeletePerson(id int) (bool, error)
}

type Handler struct {
	service Service
	log     *zap.Logger
}

func New(log *zap.Logger, service Service) Handler {
	return Handler{
		service: service,
		log:     log,
	}
}

func (h Handler) Init() *gin.Engine {
	r := gin.New()

	r.POST("/person", h.NewPerson)
	r.GET("/person", h.GetPerson)
	r.DELETE("/person", h.DeletePerson)

	return r
}
