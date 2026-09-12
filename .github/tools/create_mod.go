// Command create_mod produces a Go module source zip that is byte-identical in
// layout to the one served by proxy.golang.org, using Go's own module-zip
// implementation (golang.org/x/mod/zip). CreateFromDir applies the proxy's exact
// inclusion rules: it excludes .git/, vendor/ and nested modules.
//
// It lives under .github/ so the Go tool ignores it (directories beginning with
// a dot are invisible to ./...) and so the workflow can exclude the whole
// .github tree from the zipped copy of the source.
//
// usage: go run create_mod.go <module-path> <version> <source-dir> <output-zip>
package main

import (
	"log"
	"os"

	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatal("usage: create_mod <module-path> <version> <source-dir> <output-zip>")
	}
	m := module.Version{Path: os.Args[1], Version: os.Args[2]}
	f, err := os.Create(os.Args[4])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := zip.CreateFromDir(f, m, os.Args[3]); err != nil {
		log.Fatal(err)
	}
	log.Printf("created module zip: %s", os.Args[4])
}
