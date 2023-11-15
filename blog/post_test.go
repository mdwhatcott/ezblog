package blog

import (
	"errors"
	"os"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
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
	disk     map[string]string
	readErr  error
	renderer *Renderer
}

func (this *Fixture) Setup() {
	this.disk = make(map[string]string)
	this.renderer = NewRenderer(this)
}

func (this *Fixture) ReadFile(path string) ([]byte, error) {
	return []byte(this.disk[path]), this.readErr
}
func (this *Fixture) WriteFile(name string, data []byte, _ os.FileMode) error {
	this.disk[name] = string(data)
	return nil
}

func (this *Fixture) render() error {
	return this.renderer.RenderPost("/input.md", "/output/")
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

	this.So(err, should.Wrap, ErrReadFile)
	this.So(this.disk, should.NotContainKey, "/output/hello-world/index.html")
}
