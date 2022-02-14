package main

import (
	"fmt"
	"github.com/masibw/bazel-playground/go/pkg/random"
	"github.com/masibw/bazel-playground/go/pkg/uuid"
)

func main() {
	fmt.Println("Hello bazel!", random.GenRandomInt())
	fmt.Println("UUID :", uuid.Gen())
}
