package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yebology/travique-go/constant"
	"github.com/yebology/travique-go/controller/helper"
	"github.com/yebology/travique-go/database"
	"github.com/yebology/travique-go/jwt"
	"github.com/yebology/travique-go/model"
	"github.com/yebology/travique-go/output"
	"go.mongodb.org/mongo-driver/bson"
)

func Register(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToParseData))
	}

	hashedPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToHashPassword))
	}

	user.Password = hashedPass

	collection := database.GetDatabase().Collection("user")
	filter := bson.M{
		"email": user.Email,
	}

	_, err = helper.GetSpecificUser(ctx, filter, collection)
	if err == nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.DuplicateDataError))
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToInsertData))
	}

	specificUser, err := helper.GetSpecificUser(ctx, filter, collection)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToLoadUserData))
	}

	jwt, err := jwt.GenerateJwt(user)
	if err != nil {
		return output.GetError(c, fiber.StatusBadRequest, string(constant.FailedToGenerateTokenAccess))
	}

	return output.GetSuccess(c, fiber.Map{
		"message": "Successfully registered new user!",
		"data": fiber.Map{
			"user": specificUser,
		},
		"jwt": jwt,
	})
	
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