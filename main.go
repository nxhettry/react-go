package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("Hello World")

	app := fiber.New()

	log.Fatal(app.Listen(":8080"))
}
