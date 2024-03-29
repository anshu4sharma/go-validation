package server

import (
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			// Global custom error handler
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return c.Status(fiber.StatusBadRequest).JSON(GlobalErrorHandlerResp{
					Success: false,
					Message: err.Error(),
				})
			},
		}),
	}
	return server
}
