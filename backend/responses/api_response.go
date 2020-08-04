package responses

import "github.com/gofiber/fiber"

func response(code int64, content fiber.Map, c *fiber.Ctx) {
	mapToSend := fiber.Map{
		"code":    code,
		"content": content,
	}

	if err := c.JSON(mapToSend); err != nil {
		c.Status(500).Send(err)
		return
	}
}

func SuccessResponse(content fiber.Map, c *fiber.Ctx) {
	// 0 means there went nothing wrong
	response(0, content, c)
}
func ErrorResponse(code int64, content fiber.Map, c *fiber.Ctx) {
	response(code, content, c)
}
