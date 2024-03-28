package main

import (
  //"encoding/json"
  "fmt"
  //"io/ioutil"
  "log"
  //"net/http"
  "time"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
  "strconv"
)

type Count struct {
  Number int
}

func main() {

  app := fiber.New(fiber.Config{
    Views: html.New("./views", ".html"),
  })

  app.Static("/", "./public", fiber.Static{
    Compress: true,
  }) 

  app.Get("/", func(ctx *fiber.Ctx) error {
    return ctx.Render("index", fiber.Map{
      "Count": 0,
    })
  })
  
  app.Post("/count", func(ctx *fiber.Ctx) error {

    time.Sleep(1 *time.Second)
    count := ctx.FormValue("count")
   
    fmt.Println(count)
    
    num, err := strconv.Atoi(count)
	  if err != nil {
		  fmt.Println("Conversion failed:", err)
		  return nil
	  }
	  fmt.Println("Converted integer:", num)
    
    numer := num + 1

    return ctx.Render("index", fiber.Map{
      "Count": numer,
    })
  })

  log.Fatal(app.Listen(":9000"))
}
