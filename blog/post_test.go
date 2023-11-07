package blog

import (
	os2 "os"
	"testing"

	"github.com/mdwhatcott/ezblog/os"
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

func Test(t *testing.T) {
	disk := map[string]string{"/input.md": inputFile}
	os.ReadFile = func(path string) ([]byte, error) {
		return []byte(disk[path]), nil
	}
	os.WriteFile = func(name string, data []byte, perm os2.FileMode) error {
		disk[name] = string(data)
		return nil
	}

	RenderPost("/input.md", "/output/")

	actualContent := disk["/output/hello-world/index.html"]
	if actualContent != expectedContent {
		t.Errorf("\nwant: %s\ngot:  %s", expectedContent, actualContent)
	}
}
