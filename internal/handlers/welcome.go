package handlers

func welcome(c *fiber.ctx) error {
	return c.Render("welcome", nil, "layouts/main")
}
