package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed assets/*

var Dist embed.FS

func ServerAssets(app *fiber.App) {

	app.Get("app.wasm", func(c *fiber.Ctx) error {
		wasm, _ := Dist.ReadFile("assets/app.wasm")
		c.Set(fiber.HeaderContentType, "application/wasm")
		return c.Send(wasm)

	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(Dist),
		PathPrefix: "assets",
	}))
}

func SetData(c *fiber.Ctx) error {
	if username := c.FormValue("username"); username != "" {

		fmt.Printf("Receive POST with %s\n", username)
	}
	return c.Send([]byte("OK"))

}

func Routing(app *fiber.App) {
	approute := app.Group("/app")

	approute.Post("/data", SetData)
}

func main() {

	app := fiber.New(fiber.Config{
		ServerHeader: "hogosuruserver",
	})

	Routing(app)
	ServerAssets(app)
	app.Listen(":8080")
}
