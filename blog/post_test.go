package blog

import (
	"os"
	"testing"
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

type FakeFS struct {
	disk map[string]string
}

func (this *FakeFS) ReadFile(path string) ([]byte, error) {
	return []byte(this.disk[path]), nil
}

func (this *FakeFS) WriteFile(name string, data []byte, perm os.FileMode) error {
	this.disk[name] = string(data)
	return nil
}

func Test(t *testing.T) {
	disk := map[string]string{"/input.md": inputFile}

	RenderPost(&FakeFS{disk}, "/input.md", "/output/")

	actualContent := disk["/output/hello-world/index.html"]
	if actualContent != expectedContent {
		t.Errorf("\nwant: %s\ngot:  %s", expectedContent, actualContent)
	}
}
