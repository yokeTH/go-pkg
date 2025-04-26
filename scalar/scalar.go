package scalar

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var DefaultHandler = New()

func New(config ...Config) fiber.Handler {
	var cfg Config = defaultConfig
	if len(config) > 1 {
		cfg = loadDefaultConfig(config[0])
	}
	var (
		setupErr error
	)

	html, err := template.New("index.html").Parse(templeteHTML)
	if err != nil {
		setupErr = fmt.Errorf("Failed to parse html template:%v", err)
	}

	if cfg.DocsJsonContent == "" {
		_, err = os.Stat(cfg.DocsJsonPath)
		if os.IsNotExist(err) {
			setupErr = fmt.Errorf("%s file does not exist", cfg.DocsJsonPath)
		}

		rawSpec, err := os.ReadFile(cfg.DocsJsonPath)
		if err != nil {
			setupErr = fmt.Errorf("Failed to read provided Swagger file (%s): %v", cfg.DocsJsonPath, err.Error())
		}
		cfg.DocsJsonContent = string(rawSpec)
	}

	return func(ctx *fiber.Ctx) error {
		if setupErr != nil {
			return setupErr
		}

		if strings.HasSuffix(ctx.Path(), cfg.DocsJsonUrl) {
			return ctx.Type("json").SendString(cfg.DocsJsonContent)
		}

		ctx.Type("html")
		return html.Execute(ctx, cfg)
	}
}
