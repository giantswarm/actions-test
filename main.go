package main

import (
	"fmt"

	"github.com/giantswarm/actions-test/pkg/project"
)

func main() {
	fmt.Printf("Hello, world! %s", project.Version())
}
