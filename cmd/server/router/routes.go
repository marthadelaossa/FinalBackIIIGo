package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	handlerOdontologo "github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/odontologo"
	handlerPaciente "github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/paciente"
	"github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/ping"
	"github.com/marthadelaossa/FinalBackIIIGo/pkg/middleware"

	odontologo "github.com/marthadelaossa/FinalBackIIIGo/internal/odontologo"
	paciente "github.com/marthadelaossa/FinalBackIIIGo/internal/paciente"
)

// Router interface defines the methods that any router must implement.
type Router interface {
	MapRoutes()
}

// router is the Gin router.
type router struct {
	engine      *gin.Engine
	routerGroup *gin.RouterGroup
	db          *sql.DB
}

// NewRouter creates a new Gin router.
func NewRouter(engine *gin.Engine, db *sql.DB) Router {
	return &router{
		engine: engine,
		db:     db,
	}
}

// MapRoutes maps all routes.
func (r *router) MapRoutes() {
	r.setGroup()
	r.buildPingRoutes()
	r.buildProductRoutes()
	r.buildPacienteRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

// buildProductRoutes maps all routes for the product domain.
func (r *router) buildProductRoutes() {
	// Create a new odontologo controller.
	repository := odontologo.NewMySqlRepository(r.db)
	service := odontologo.NewServiceOdontologo(repository)
	controlador := handlerOdontologo.NewControladorProducto(service)

	grupoOdontologo := r.routerGroup.Group("/odontologo")
	{
		grupoOdontologo.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoOdontologo.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoOdontologo.GET("/:id", controlador.HandlerGetByID())
		grupoOdontologo.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoOdontologo.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoOdontologo.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())

	}

}

func (r *router) buildPacienteRoutes() {
	// Create a new paciente controller.
	repository := paciente.NewMySqlRepository(r.db)
	service := paciente.NewServicePaciente(repository)
	controlador := handlerPaciente.NewControladorPaciente(service)

	grupoPaciente := r.routerGroup.Group("/paciente")
	{
		grupoPaciente.POST("", middleware.Authenticate(), controlador.HandlerCreate())
		grupoPaciente.GET("", middleware.Authenticate(), controlador.HandlerGetAll())
		grupoPaciente.GET("/:id", controlador.HandlerGetByID())
		grupoPaciente.PUT("/:id", middleware.Authenticate(), controlador.HandlerUpdate())
		grupoPaciente.DELETE("/:id", middleware.Authenticate(), controlador.HandlerDelete())
		grupoPaciente.PATCH("/:id", middleware.Authenticate(), controlador.HandlerPatch())

	}

}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}
