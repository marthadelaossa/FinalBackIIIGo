package paciente

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marthadelaossa/FinalBackIIIGo/internal/domain"
	"github.com/marthadelaossa/FinalBackIIIGo/internal/paciente"
	"github.com/marthadelaossa/FinalBackIIIGo/pkg/web"
)

type Controlador struct {
	service paciente.Service
}

func NewControladorPaciente(service paciente.Service) *Controlador {
	return &Controlador{
		service: service,
	}
}

// Paciente godoc
// @Summary paciente example
// @Description Create a new paciente
// @Tags paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /paciente [post]
func (c *Controlador) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Paciente

		err := ctx.Bind(&request)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		paciente, err := c.service.Create(ctx, request)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, paciente)

	}
}

// Paciente godoc
// @Summary paciente example
// @Description Get all pacientes
// @Tags paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /pacientes [get]
func (c *Controlador) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		paciente, err := c.service.GetAll(ctx)

		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, paciente)
	}
}

// Paciente godoc
// @Summary paciente example
// @Description Get paciente by id
// @Tags paciente
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [get]
func (c *Controlador) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		paciente, err := c.service.GetByID(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, paciente)
	}
}

// Paciente godoc
// @Summary paciente example
// @Description Update paciente by id
// @Tags paciente
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [put]
func (c *Controlador) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request domain.Paciente

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		id := ctx.Param("id")

		idInt, err := strconv.Atoi(id)

		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request param")
			return
		}

		paciente, err := c.service.Update(ctx, request, idInt)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, paciente)

	}
}

// Paciente godoc
// @Summary paciente example
// @Description Delete paciente by id
// @Tags paciente
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [delete]
func (c *Controlador) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, "paciente eliminado")
	}
}

// Paciente godoc
// @Summary paciente example
// @Description Patch paciente
// @Tags paciente
// @Param id path int true "id del paciente"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /pacientes/:id [patch]
func (c *Controlador) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var request domain.Paciente

		errBind := ctx.Bind(&request)

		if errBind != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		paciente, err := c.service.Patch(ctx, request, id)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, paciente)
	}
}
