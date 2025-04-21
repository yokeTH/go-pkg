package scalar

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	docsPath = "./docs/swagger.json"
)

var DefaultHandler = new(docsPath)

func Handler(path string) fiber.Handler {
	return new(path)
}

func new(docsPath string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var rawSpec []byte
		_, err := os.Stat(docsPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("%s file does not exist", docsPath)
		}
		rawSpec, err = os.ReadFile(docsPath)
		if err != nil {
			return fmt.Errorf("Failed to read provided Swagger file (%s): %v", docsPath, err.Error())
		}
		return ctx.Type("html").SendString(getHtmlByContent(string(rawSpec)))
	}
}
