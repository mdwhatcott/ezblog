package blog

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

func RenderPost(sourcePath, destDir string) {
	source, _ := os.ReadFile(sourcePath)
	segments := bytes.Split(source, []byte("\n+++\n"))
	frontMatter := make(map[string]string)
	_ = json.Unmarshal(segments[0], &frontMatter)
	var content bytes.Buffer
	_ = goldmark.New().Convert(segments[1], &content)
	rendered := []byte(strings.NewReplacer(
		"{{ Title }}", frontMatter["title"],
		"{{ Slug }}", frontMatter["slug"],
		"{{ Author }}", frontMatter["author"],
		"{{ Body }}", content.String(),
	).Replace(pageTemplate))
	path := filepath.Join(destDir, frontMatter["slug"], "index.html")
	_ = os.WriteFile(path, rendered, 0644)
}

const pageTemplate = `<html>
<head>
	<title>{{ Title }}</title>
	<link rel="canonical" href="https://your-domain-here.com{{ Slug }}">
</head>
<body>
	<h1>{{ Title }}</h1>
	<p>By {{ Author }}</p>
	<div>
	{{ Body }}
	</div>
</body>
</html>`
