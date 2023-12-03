package odontologo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
	"github.com/marthadelaossa/FinalBackIIIGo/internal/odontologo"
	"github.com/marthadelaossa/FinalBackIIIGo/pkg/web"
)

const (
	badRequestErrorMessage     = "Bad request"
	invalidIdErrorMessage      = "Invalid ID"
	internalServerErrorMessage = "Internal server error"
)

type Controller struct {
	service odontologo.Service
}

func NewController(service odontologo.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.Odontologo

		if err := ctx.Bind(&request); err != nil {
			web.Error(ctx, http.StatusBadRequest, badRequestErrorMessage)
			return
		}

		odontologo, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, internalServerErrorMessage)
			return
		}

		web.Success(ctx, http.StatusOK, odontologo)
	}
}

func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		odontologos, err := c.service.GetAll(ctx)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, internalServerErrorMessage)
			return
		}

		web.Success(ctx, http.StatusOK, odontologos)
	}
}

func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, invalidIdErrorMessage)
			return
		}

		odontologo, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, internalServerErrorMessage)
			return
		}

		web.Success(ctx, http.StatusOK, odontologo)
	}
}

func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request domain.Odontologo

		if err := ctx.Bind(&request); err != nil {
			web.Error(ctx, http.StatusBadRequest, badRequestErrorMessage)
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, badRequestErrorMessage)
			return
		}

		odontologo, err := c.service.Update(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, internalServerErrorMessage)
			return
		}

		web.Success(ctx, http.StatusOK, odontologo)
	}
}

func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, invalidIdErrorMessage)
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, internalServerErrorMessage)
			return
		}

		web.Success(ctx, http.StatusOK, "Odontólogo eliminado")
	}
}

func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, invalidIdErrorMessage)
			return
		}

		var request domain.Odontologo

		if err := ctx.Bind(&request); err != nil {
			web.Error(ctx, http.StatusBadRequest, badRequestErrorMessage)
			return
		}

		odontologo, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, internalServerErrorMessage)
			return
		}

		web.Success(ctx, http.StatusOK, odontologo)
	}
}
