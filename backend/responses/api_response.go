package responses

import "github.com/gofiber/fiber"

func Response(value fiber.Map, c *fiber.Ctx) {
	if err := c.JSON(value); err != nil {
		c.Status(500).Send(err)
		return
	}
}
