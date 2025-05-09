package controller

import (
	// "context"
	// "time"

	"github.com/gofiber/fiber/v2"
	// "github.com/yebology/travique-go/constant"
	// "github.com/yebology/travique-go/model"
	// "github.com/yebology/travique-go/output"
)

func Register(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
        "message": "Login endpoint reached",
    })

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// var user model.User
	// err := c.BodyParser(&user)
	// if err != nil {
	// 	return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	// }


	
	// return output.GetSuccess(c, )
}

func Login(c *fiber.Ctx) error {
	 
	return c.JSON(fiber.Map{
        "message": "Login endpoint reached",
    })

}

func EditProfile(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
        "message": "",
    })

}