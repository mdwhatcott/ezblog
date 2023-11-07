package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mdwhatcott/ezblog/blog"
)

var Version = "dev"

type Config struct {
	SourceFile string
	DestDir    string
}

func main() {
	var config Config
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	flags.StringVar(&config.SourceFile, "source", "", "The path to the source file.")
	flags.StringVar(&config.DestDir, "dest", "", "The path to the destination folder.")
	_ = flags.Parse(os.Args[1:])

	err := os.MkdirAll(config.DestDir, 0577)
	if err != nil {
		log.Fatal(err)
	}
	blog.RenderPost(fs{}, config.SourceFile, config.DestDir)
}

type fs struct{}

func (fs) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}
func (fs) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}
