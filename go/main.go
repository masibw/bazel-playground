package main

import (
	"fmt"
	"github.com/masibw/bazel-playground/go/pkg/random"
)

func main() {
	fmt.Println("Hello bazel!", random.GenRandomInt())
}
