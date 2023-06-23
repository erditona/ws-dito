package main

import (
	"log"

	"github.com/erditona/ws-dito/config"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/whatsauth/whatsauth"

	"github.com/erditona/ws-dito/url"

	"github.com/gofiber/fiber/v2"

	_ "github.com/erditona/ws-dito/docs"
)

// @title Swagger Documentation API
// @version 1.0
// @description This is a sample server

// @contact.name API Support
// @contact.url https://github.com/erditona
// @contact.email erditonaushaadam@gmail.com

// @host ws-dito.herokuapp.com
// @BasePath /
// @schemes http https

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
