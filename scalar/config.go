package scalar

import "html/template"

type Config struct {
	// Title tag of html scalar
	// default: "Scalar API Reference"
	Title string

	// Json string of OPEN API
	// default: ""
	DocsJsonContent string

	// Url of json content
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

var defaultConfig = Config{
	Title:        "Scalar API Reference",
	DocsJsonUrl:  "doc.json",
	DocsJsonPath: "./docs/swagger.json",
	ProxyUrl:     "https://proxy.scalar.com",
}

func loadDefaultConfig(config Config) Config {
	var cfg Config = defaultConfig

	if config.Title != "" {
		cfg.Title = config.Title
	}

	if config.DocsJsonUrl != "" {
		cfg.DocsJsonUrl = config.DocsJsonPath
	}

	if config.DocsJsonPath != "" {
		cfg.DocsJsonPath = config.DocsJsonPath
	}

	if config.ProxyUrl != "" {
		cfg.ProxyUrl = config.ProxyUrl
	}

	return cfg
}
