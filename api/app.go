package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"golang-ecommerce-practice/api/routes"
	"golang-ecommerce-practice/package/users"
	"golang-ecommerce-practice/zapLog"
	"log"
	"os"
	"time"
)

func ConnectionDB() (*sqlx.DB, error) {
	cn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_NAME"))
	client, err := sqlx.Open("mysql", cn)

	if err != nil {
		zapLog.Error("Error Connect database" + err.Error())
		return nil, err
	}
	client.SetConnMaxLifetime(time.Minute * 30)
	client.SetConnMaxIdleTime(time.Minute * 1)
	client.SetMaxOpenConns(100)
	client.SetMaxIdleConns(10)

	return client, nil
}
func main() {
	//use godotenv if not using docker
	godotenv.Load(".config")

	db, err := ConnectionDB()
	if err != nil {
		log.Fatal("Database Connection Error $s", err)
	}
	fmt.Println("Database Connection Success")
	usersRepository := users.NewRepoDB(db)
	usersService := users.NewService(usersRepository)
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to Api!"))
	})
	routes.UserRouter(app.Group("/api/v1/users"), usersService)

	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))

}
