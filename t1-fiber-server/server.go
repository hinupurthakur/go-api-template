package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	fmt.Println("Go Fiber")
	app := fiber.New()
	app.Get("/",func (c *fiber.Ctx) error{
		return c.SendString("Hii from server port 3000")
	})
	
	app.Get("/:value?",valueHandler)
	err := app.Listen(":3000")
	if err!=nil{
		fmt.Printf("Error staring the server:- %s",err)
	}
}

func valueHandler(ctx *fiber.Ctx) error{
	return ctx.SendString("Query parameter: "+ctx.Params("value"))
}	