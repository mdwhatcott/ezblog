package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var Version = "dev"

func main() {
	flags := flag.NewFlagSet(fmt.Sprintf("%s @ %s", filepath.Base(os.Args[0]), Version), flag.ExitOnError)
	_ = flags.Parse(os.Args[1:])
}
