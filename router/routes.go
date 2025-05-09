package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/travique-go/controller"
)

func SetUp(app *fiber.App) {

	app.Post("/api/register", controller.Register)

	app.Post("/api/login", controller.Login)

	app.Patch("/api/edit_profile", controller.EditProfile)

}