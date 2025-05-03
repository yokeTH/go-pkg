# Scalar

## Usage
1. Add the swagger comment following this [reference](https://github.com/swaggo/swag#declarative-comments-format)
2. Download Swag command latest
```bash
go install github.com/swaggo/swag/v2/cmd/swag@v2.0.0-rc4
```
3. Use Swag command from 2. to generate the openAPI 3.1 docs
```bash
swag init -v3.1 -o docs -g main.go --parseDependency --parseInternal
```
4. Download Scalar UI package
```bash
go get github.com/yokeTH/go-pkg/scalar
```
And use as the fiber handler to your project
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yokeTH/go-pkg/scalar"
)

//	@title Scalar Example
//	@version 1.0
//	@servers https http
func main() {
	app := fiber.New()

	app.Get("/swagger/*", scalar.HandlerDefault)

	app.Get("/swagger/*", scalar.New(scalar.Config{
		// if your generate docs output path is not docs
		DocsJsonPath: "./doc/swagger.json",
	}))

	app.Listen(":8080")
}
```

## Config
```go
type Config struct {
	// Title tag of html scalar
	// default: "Scalar API Reference"
	Title string

	// Json string of OPEN API
	// default: ""
	DocsJsonContent string

	// Url of json content
	// example: app.Get("/swagger/*", scalar.HandlerDefault) -> /swagger/doc.json will serve the json of openAPI
	// default: "doc.json"
	DocsJsonUrl string

	// Path of generated docs
	// default: "./docs/swagger.json"
	DocsJsonPath string

	// Custom Scalar Style
	// Ref: https://github.com/scalar/scalar/blob/main/packages/themes/src/variables.css
	// default: ""
	CustomStyle template.CSS

	// Proxy to avoid CORS issues
	// default: "https://proxy.scalar.com"
	ProxyUrl string
}
```
