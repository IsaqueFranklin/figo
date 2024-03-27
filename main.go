package main

import (
  //"encoding/json"
  //"fmt"
  //"io/ioutil"
  "log"
  //"net/http"
  //"time"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
)

func main() {

  app := fiber.New(fiber.Config{
    Views: html.New("./views", ".html"),
  })

  app.Static("/", "./public", fiber.Static{
    Compress: true,
  }) 

  app.Get("/", func(ctx *fiber.Ctx) error {
    return ctx.Render("index", fiber.Map{})
  })

  log.Fatal(app.Listen(":9000"))
}
