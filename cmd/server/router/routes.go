package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	appointmentHandler "github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/appointment"
	dentistHandler "github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/odontologo"
	patientHandler "github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/paciente"
	"github.com/marthadelaossa/FinalBackIIIGo/cmd/server/handler/ping"
	"github.com/marthadelaossa/FinalBackIIIGo/pkg/middleware"

	appointment "github.com/marthadelaossa/FinalBackIIIGo/internal/appointment"
	dentist "github.com/marthadelaossa/FinalBackIIIGo/internal/odontologo"
	patient "github.com/marthadelaossa/FinalBackIIIGo/internal/paciente"
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
	r.buildDentistRoutes()
	r.buildPatientsRoutes()
	r.buildAppointmentsRoutes()
}

// setGroup sets the router group.
func (r *router) setGroup() {
	r.routerGroup = r.engine.Group("/api/v1")
}

type HandlerProvider interface {
	HandlerCreate() gin.HandlerFunc
	HandlerGetAll() gin.HandlerFunc
	HandlerGetByID() gin.HandlerFunc
	HandlerUpdate() gin.HandlerFunc
	HandlerDelete() gin.HandlerFunc
	HandlerPatch() gin.HandlerFunc
}

func (r *router) buildRoutesGeneric(routeGroup *gin.RouterGroup, controller HandlerProvider) {
	routeGroup.POST("", middleware.Authenticate(), controller.HandlerCreate())
	routeGroup.GET("", controller.HandlerGetAll())
	routeGroup.GET("/:id", controller.HandlerGetByID())
	routeGroup.PUT("/:id", middleware.Authenticate(), controller.HandlerUpdate())
	routeGroup.DELETE("/:id", middleware.Authenticate(), controller.HandlerDelete())
	routeGroup.PATCH("/:id", middleware.Authenticate(), controller.HandlerPatch())
}

// buildDentistRoutes maps all routes for the product domain.
func (r *router) buildDentistRoutes() {
	// Create a new odontologo controller.
	repository := dentist.NewMySqlRepository(r.db)
	service := dentist.NewServiceOdontologo(repository)
	controller := dentistHandler.NewController(service)

	dentistRoutesGroup := r.routerGroup.Group("/odontologo")
	r.buildRoutesGeneric(dentistRoutesGroup, controller)
}

func (r *router) buildPatientsRoutes() {
	// Create a new paciente controller.
	repository := patient.NewMySqlRepository(r.db)
	service := patient.NewServicePaciente(repository)
	controller := patientHandler.NewController(service)

	patientRoutesGroup := r.routerGroup.Group("/paciente")
	r.buildRoutesGeneric(patientRoutesGroup, controller)
}

func (r *router) buildAppointmentsRoutes() {
	repository := appointment.NewAppointmentMySqlRepository(r.db)
	service := appointment.NewAppointmentService(repository)
	controller := appointmentHandler.NewController(service)

	appointmentRoutesGroup := r.routerGroup.Group("/appointments")

	r.buildRoutesGeneric(appointmentRoutesGroup, controller)
}

// buildPingRoutes maps all routes for the ping domain.
func (r *router) buildPingRoutes() {
	// Create a new ping controller.
	pingController := ping.NewControllerPing()
	r.routerGroup.GET("/ping", pingController.HandlerPing())

}
