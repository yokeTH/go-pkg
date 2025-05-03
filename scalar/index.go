package scalar

const templeteHTML = `
	<!doctype html>
<html>
  <head>
    <title>{{.Title}}</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
      	{{- if .CustomStyle}}
       	<style>
       	:root {
       	{{ .CustomStyle }}
       	}
		</style>
        {{end}}
  </head>

  <body>
    <div id="app"></div>

    <!-- Load the Script -->
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>

    <!-- Initialize the Scalar API Reference -->
    <script>
      Scalar.createApiReference('#app', {
        // The content of the OpenAPI/Swagger document
        content: {{.DocsJsonContent}},

        // Avoid CORS issues
        proxyUrl: {{.ProxyUrl}},
      })
    </script>
  </body>
</html>`
