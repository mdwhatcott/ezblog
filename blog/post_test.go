package blog

import (
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
	renderer *Renderer
}

func (this *Fixture) Setup() {
	this.disk = make(map[string]string)
	this.renderer = NewRenderer(this)
}

func (this *Fixture) ReadFile(path string) ([]byte, error) {
	return []byte(this.disk[path]), nil
}

func (this *Fixture) WriteFile(name string, data []byte, _ os.FileMode) error {
	this.disk[name] = string(data)
	return nil
}

func (this *Fixture) TestRenderPost() {
	this.disk["/input.md"] = inputFile

	this.renderer.RenderPost("/input.md", "/output/")

	this.So(this.disk["/output/hello-world/index.html"], should.Equal, expectedContent)
}
