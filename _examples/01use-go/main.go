package main

import (
	"log"
	"github.com/podhmo/multi-package-example/go/foo"
	"github.com/podhmo/multi-package-example/go/bar"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("!! %+v", err)
	}
}

func run() error {
	foo.Hello()
	bar.Hello()
	return nil
}
