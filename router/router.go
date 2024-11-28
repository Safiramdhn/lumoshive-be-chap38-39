package router

import (
	"lumoshive-be-chap38-39/infra"

	"github.com/gin-gonic/gin"
)

// APIRouter configures the API routes for the shipping service
func NewRouter(ctx infra.ServiceContext) *gin.Engine {
	// Initialize the Gin router
	router := gin.Default()

	// Define route groups
	registerShippingCourierRoutes(router, ctx)
	registerShippingRoutes(router, ctx)

	return router
}

// registerShippingRoutes sets up the shipping-related routes.
func registerShippingCourierRoutes(router *gin.Engine, ctx infra.ServiceContext) {
	shippingCourierGroup := router.Group("/shipping-couriers")

	// Define shipping routes
	shippingCourierGroup.GET("/list", ctx.Ctl.ShippingCourierController.GetAllShippingCourierController)
	shippingCourierGroup.GET("/:id", ctx.Ctl.ShippingCourierController.GetShippingCourierByIdController)
	shippingCourierGroup.GET(
		"/cost/:id/:quantity/:origin_longlat/:destination_longlat",
		ctx.Ctl.ShippingCourierController.GetShippingCostController,
	)
}

func registerShippingRoutes(router *gin.Engine, ctx infra.ServiceContext) {
	shippingGroup := router.Group("/shipping")

	shippingGroup.POST("/", ctx.Ctl.ShippingController.CreateShippingController)
	shippingGroup.GET("/:id", ctx.Ctl.ShippingController.GetShippingByID)

	historyGroup := shippingGroup.Group("/history")
	historyGroup.POST("/:shipping_id", ctx.Ctl.ShippingHistoryController.UpdateShippingHistoryController)
	historyGroup.GET("/:shipping_id", ctx.Ctl.ShippingHistoryController.GetShippingHistoryController)
}
