/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/tagesjump/provider-upjet-yc/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	p, err := config.GetProvider(context.Background(), true)
	kingpin.FatalIfError(err, "Cannot initialize the provider configuration")
	pipeline.Run(p, absRootDir)
}
