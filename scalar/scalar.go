package scalar

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

const (
	docsPathJson = "./docs/swagger.json"
)

var DefaultHandler = new(docsPathJson)

func Handler(path string) fiber.Handler {
	return new(path)
}

func new(docsPath string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		_, err := os.Stat(docsPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("%s file does not exist", docsPath)
		}

		rawSpec, err := os.ReadFile(docsPath)
		if err != nil {
			return fmt.Errorf("Failed to read provided Swagger file (%s): %v", docsPath, err.Error())
		}

		if strings.HasSuffix(ctx.Path(), "docs.json") {
			return ctx.Type("json").Send(rawSpec)
		}

		return ctx.Type("html").SendString(getHtmlByContent(string(rawSpec)))
	}
}
