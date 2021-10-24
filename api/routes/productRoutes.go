package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang-ecommerce-practice/package/auth"
	"golang-ecommerce-practice/package/entities"
	"golang-ecommerce-practice/package/products"
	"golang-ecommerce-practice/zapLog"
)

func ProductRouter(app fiber.Router, service products.RepoServiceProducts, serveAuth auth.RepoServiceAuth) {
	app.Get("/", getProducts(service))
	app.Get("/:id", Protect(serveAuth), getProduct(service))
	app.Post("/", Protect(serveAuth), createProduct(service))
	app.Patch("/:id", Protect(serveAuth), updateProduct(service))
	app.Delete("/:id", Protect(serveAuth), deleteProduct(service))
}

func getProducts(srv products.RepoServiceProducts) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchData, err := srv.GetProducts(&map[string]interface{}{})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message": err.GetMsg().Message + " Data Products",
			})
		}
		return c.Status(200).JSON(&fetchData)
	}
}
func getProduct(srv products.RepoServiceProducts) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetchData, err := srv.GetProduct(&map[string]interface{}{
			"id": c.Params("id"),
		})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message": err.GetMsg().Message + " Data Products",
			})
		}

		return c.Status(200).JSON(&fetchData)
	}
}

func createProduct(srv products.RepoServiceProducts) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(entities.CreateProducts)
		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser " + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		user, err := srv.InsertProducts(body)
		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(201).JSON(&user)
	}
}
func updateProduct(srv products.RepoServiceProducts) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(map[string]interface{})

		if err := c.BodyParser(&body); err != nil {
			zapLog.Error("Error body parser" + err.Error())
			return fiber.NewError(fiber.StatusBadRequest, "Bad Request")
		}

		id := c.Params("id")
		err := srv.UpdateProducts(body, &id)

		if err != nil {
			return fiber.NewError(err.Code, err.Message)
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "update success",
		})
	}
}
func deleteProduct(srv products.RepoServiceProducts) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := srv.DeleteProducts(&map[string]interface{}{
			"id": c.Params("id"),
		})
		if err != nil {
			return c.Status(err.Code).JSON(fiber.Map{
				"message": err.GetMsg().Message + " Data Products",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Delete success",
		})
	}
}
