package blog

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

type FS interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type Renderer struct {
	fs FS
}

func NewRenderer(fs FS) *Renderer {
	return &Renderer{fs: fs}
}

func (this *Renderer) RenderPost(sourcePath, destDir string) {
	source, _ := this.fs.ReadFile(sourcePath)
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
	_ = this.fs.WriteFile(path, rendered, 0644)
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
