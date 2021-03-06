// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/teris-io/cli"
	"github.com/teris-io/geotiff/tags"
)

func main() {
	os.Exit(Run(os.Args, os.Stdout))
}

func Run(args []string, w io.Writer) int {
	tagsCmd := cli.NewCommand("tags", "Print image tags").
		WithShortcut("t").
		WithArg(cli.NewArg("filename", "GeoTIFF file name")).
		WithOption(cli.NewOption("page", "Page to read tags for (0-based, default: all pages)").WithChar('p').WithType(cli.TypeInt)).
		WithOption(cli.NewOption("verbose", "Verbose output of arrays").WithChar('v').WithType(cli.TypeBool)).
		WithAction(func(args []string, options map[string]string) int {
			page := -1
			if pagestr, singlepage := options["page"]; singlepage {
				var err error
				page, err = strconv.Atoi(pagestr)
				if err != nil {
					return done(err)
				}
			}
			_, verbose := options["verbose"]
			return done(tags.FprintForFile(w, args[0], page, verbose))
		})

	versionCmd := cli.NewCommand("version", "Show version information").
		WithShortcut("v").
		WithAction(func(args []string, options map[string]string) int {
			fmt.Fprintln(w, Version)
			return 0
		})

	app := cli.New("CLI to work with GeoTIFF files").
		WithCommand(tagsCmd).
		WithCommand(versionCmd)

	return app.Run(args, w)
}
