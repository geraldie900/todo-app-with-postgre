package apiv1

import (
	"github.com/gofiber/fiber/v2"
)

func ApiV1(app fiber.Router) {
	v1 := app.Group("/v1")

	TodoHandler(v1)
}
