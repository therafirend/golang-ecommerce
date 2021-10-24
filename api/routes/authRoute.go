package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-ecommerce-practice/package/auth"
	"golang-ecommerce-practice/package/entities"
	"golang-ecommerce-practice/zapLog"
)

func AuthRouter(app fiber.Router, service auth.RepoServiceAuth) {
	app.Post("/", login(service))
	app.Post("/change-password", Protect(service), changePassword(service))
}

func login(srv auth.RepoServiceAuth) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.BodyLogin)

		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser" + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "internal server error body parser")
		}

		res, err := srv.Login(body)

		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(&res)
	}
}

func changePassword(srv auth.RepoServiceAuth) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.ChangePassword)

		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "bad request")
		}

		body.ID = string(c.Request().Header.Peek("ID"))

		if err := srv.ChangePassword(body); err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(&fiber.Map{
			"message": "password successfuly change",
		})
	}
}

func Protect(srv auth.RepoServiceAuth) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Request().Header.Peek("Authorization")

		if authorization == nil {
			zapLog.Error("not authorization")
			return fiber.NewError(fiber.StatusUnauthorized, "please login")
		}

		authString := string(authorization)

		res, err := srv.Auth(&authString)

		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		c.Request().Header.Add("ID", *res)

		return c.Next()
	}
}
