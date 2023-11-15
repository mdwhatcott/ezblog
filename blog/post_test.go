package blog

import (
	"errors"
	"io"
	"log"
	"os"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
	"github.com/yuin/goldmark"
)

const inputFile = `
{
	"slug":"/hello-world",
	"title":"Greetings",
	"author":"Michael Whatcott"
}
+++
### Hello, world!
`

const expectedContent = `<html>
<head>
	<title>Greetings</title>
	<link rel="canonical" href="https://your-domain-here.com/hello-world">
</head>
<body>
	<h1>Greetings</h1>
	<p>By Michael Whatcott</p>
	<div>
	<h3>Hello, world!</h3>

	</div>
</body>
</html>`

func TestFixture(t *testing.T) {
	gunit.Run(new(Fixture), t)
}

type Fixture struct {
	*gunit.Fixture
	disk        map[string]string
	readErr     error
	markdownErr error
	writeErr    error
	renderer    *Renderer
}

func (this *Fixture) Setup() {
	this.disk = make(map[string]string)
	this.renderer = NewRenderer(this, this, log.New(this, "["+this.Name()+"] ", 0))
}

func (this *Fixture) ReadFile(path string) ([]byte, error) {
	return []byte(this.disk[path]), this.readErr
}
func (this *Fixture) WriteFile(name string, data []byte, _ os.FileMode) error {
	if this.writeErr != nil {
		return this.writeErr
	}
	this.disk[name] = string(data)
	return nil
}
func (this *Fixture) Convert(source []byte, writer io.Writer) error {
	_ = goldmark.Convert(source, writer)
	return this.markdownErr
}

func (this *Fixture) render() error {
	return this.renderer.RenderPost("/input.md", "/output/")
}
func (this *Fixture) assertErrorResult(err error, expected error) {
	this.So(err, should.Wrap, expected)
	this.So(this.disk, should.NotContainKey, "/output/hello-world/index.html")
}

func (this *Fixture) TestRenderPost() {
	this.disk["/input.md"] = inputFile

	err := this.render()

	this.So(err, should.BeNil)
	this.So(this.disk["/output/hello-world/index.html"], should.Equal, expectedContent)
}
func (this *Fixture) TestRenderPost_ReadFileErr() {
	this.readErr = errors.New("boink")

	err := this.render()

	this.assertErrorResult(err, ErrReadFile)
}
func (this *Fixture) TestRenderPost_MarkdownErr() {
	this.disk["/input.md"] = inputFile
	this.markdownErr = errors.New("boink")

	err := this.render()

	this.assertErrorResult(err, ErrMarkdown)
}
func (this *Fixture) TestRenderPost_WriteFileErr() {
	this.disk["/input.md"] = inputFile
	this.writeErr = errors.New("boink")

	err := this.render()

	this.assertErrorResult(err, ErrWriteFile)
}
