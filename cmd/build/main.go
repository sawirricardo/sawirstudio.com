// Command build writes the English pages of sawirstudio.com from content/
// (settings in content/site.yaml).
//
//	go tool templ generate && go run ./cmd/build
//
// Then run the scripts that work on the built pages (see README): i18n.py for the
// translations and markdown_pages.py for the Markdown copies.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sawirricardo/sawirstudio.com/internal/site"
)

func main() {
	root := flag.String("root", ".", "the site's root folder (with content/ and assets/)")
	check := flag.Bool("check", false, "write nothing; fail if any page would change")
	flag.Parse()
	if *check {
		stale, err := site.Check(*root)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range stale {
			fmt.Println("out of date:", f)
		}
		if len(stale) > 0 {
			os.Exit(1)
		}
		return
	}
	written, err := site.Build(*root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range written {
		fmt.Println("wrote", f)
	}
	fmt.Printf("%d pages changed\n", len(written))
}
