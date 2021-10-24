package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-ecommerce-practice/package/entities"
	"golang-ecommerce-practice/package/users"
	"golang-ecommerce-practice/zapLog"
)

func UserRouter(app fiber.Router, service users.RepoServiceUsers) {
	app.Get("/", getUsers(service))
	app.Get("/:id", getUser(service))
	app.Post("/", createUser(service))
	app.Patch("/:id", updateUser(service))
	app.Delete("/:id", deleteUser(service))
}

func getUsers(srv users.RepoServiceUsers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchData, err := srv.GetUsers(&map[string]interface{}{
			"status": "on",
		})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message": err.GetMsg().Message + " Data Users",
			})
		}
		return c.Status(200).JSON(&fetchData)
	}
}
func getUser(srv users.RepoServiceUsers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchData, err := srv.GetUser(&map[string]interface{}{
			"id":     c.Params("id"),
			"status": "on",
		})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message": err.GetMsg().Message + " Data User",
			})
		}

		return c.Status(200).JSON(&fetchData)
	}
}

func createUser(srv users.RepoServiceUsers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.RegisUser)
		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		user, err := srv.InsertUser(body)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(201).JSON(&user)
	}
}
func updateUser(srv users.RepoServiceUsers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(map[string]interface{})

		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser" + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		id := c.Params("id")
		err := srv.UpdateUser(body, &id)

		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "update success",
		})
	}
}

func deleteUser(srv users.RepoServiceUsers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := srv.DeleteUser(&id); err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "delete success",
		})
	}
}
