package blog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrReadFile    = errors.New("read file")
	ErrFrontMatter = errors.New("front matter")
	ErrWriteFile   = errors.New("write file")
	ErrMarkdown    = errors.New("markdown")
)

type FS interface {
	ReadFile(path string) ([]byte, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type MD interface {
	Convert(source []byte, writer io.Writer) error
}

type Renderer struct {
	fs FS
	md MD
}

func NewRenderer(fs FS, md MD) *Renderer {
	return &Renderer{fs: fs, md: md}
}

func (this *Renderer) RenderPost(sourcePath, destDir string) error {
	source, err := this.fs.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrReadFile, err)
	}

	segments := bytes.Split(source, []byte("\n+++\n"))
	frontMatter := make(map[string]string)
	err = json.Unmarshal(segments[0], &frontMatter)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFrontMatter, err)
	}

	var content bytes.Buffer
	err = this.md.Convert(segments[1], &content)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrMarkdown, err)
	}

	rendered := []byte(strings.NewReplacer(
		"{{ Title }}", frontMatter["title"],
		"{{ Slug }}", frontMatter["slug"],
		"{{ Author }}", frontMatter["author"],
		"{{ Body }}", content.String(),
	).Replace(pageTemplate))

	path := filepath.Join(destDir, frontMatter["slug"], "index.html")
	err = this.fs.WriteFile(path, rendered, 0644)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrWriteFile, err)
	}
	return nil
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
