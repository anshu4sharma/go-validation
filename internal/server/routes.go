package server

import (
	"github.com/gofiber/fiber/v2"
	"go-validation/internal/validation"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)

}

type User struct {
	Name string `validate:"required,min=5,max=20"` // Required field, min 5 char long max 20
	Age  int    `validate:"required"`              // Required field, and client needs to implement our 'teener' tag format which we'll see later
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	user := &User{
		Name: c.Query("name"),
		Age:  c.QueryInt("age"),
	}

	if errs := validation.Custom_Validator.Validate(user); len(errs) > 0 && errs[0].Error {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Please check the input fields !",
		}
	}
	return c.JSON(resp)
}
