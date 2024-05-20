package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/nusa-exchange/finex/controllers"
	"github.com/nusa-exchange/finex/controllers/admin_controllers"
	"github.com/nusa-exchange/finex/routes/middlewares"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	api_v2_public := app.Group("/api/v2/public")
	{
		api_v2_public.Get("/timestamp", controllers.GetTimestamp)
		api_v2_public.Get("/global_price", controllers.GetGlobalPrice)
		api_v2_public.Get("/ieo/list", controllers.GetIEOList)
		api_v2_public.Get("/ieo/:id", controllers.GetIEO)
		api_v2_public.Get("/markets/:market/depth", controllers.GetDepth)
	}

	api_v2_admin := app.Group("/api/v2/admin", middlewares.Authenticate, middlewares.AdminVaildator)
	{
		api_v2_admin.Get("/trades", admin_controllers.GetTrades)
		api_v2_admin.Get("/ieo/list", admin_controllers.GetIEOList)
		api_v2_admin.Get("/ieo/:id", admin_controllers.GetIEO)
		api_v2_admin.Post("/ieo", admin_controllers.CreateIEO)
		api_v2_admin.Put("/ieo", admin_controllers.UpdateIEO)
		api_v2_admin.Delete("/ieo", admin_controllers.DeleteIEO)
		api_v2_admin.Post("/ieo/currencies", admin_controllers.AddIEOCurrencies)
		api_v2_admin.Delete("/ieo/currencies", admin_controllers.RemoveIEOCurrencies)

		api_v2_admin.Post("/orders/:uuid/cancel", admin_controllers.CancelOrder)
		api_v2_admin.Post("/orders/cancel", admin_controllers.CancelAllOrders)
	}

	return app
}
