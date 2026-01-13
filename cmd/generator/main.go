/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin/v2"
	"github.com/crossplane/upjet/v2/pkg/pipeline"
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
	p, err := config.GetProvider(true)
	n, err := config.GetProviderNamespaced(true)
	kingpin.FatalIfError(err, "Cannot initialize the provider configuration")
	pipeline.Run(p, n, absRootDir)
}
