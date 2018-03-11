/*
 * Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.
 */

package main

import (
	"fmt"
	"os"

	"github.com/teris-io/cli"
)

func main() {

	tagsCmd := cli.NewCommand("tags", "Query image tags").
		WithShortcut("t").
		WithArg(cli.NewArg("file", "GeoTIFF file name")).
		WithOption(cli.NewOption("page", "Page to read tags for (default: all pages)").WithChar('p')).
		WithOption(cli.NewOption("quiet", "Quiet output dropping all auxiliary information").WithChar('q').WithType(cli.TypeBool)).
		WithAction(tags)

	versionCmd := cli.NewCommand("version", "Show version information").
		WithShortcut("v").
		WithAction(func(args []string, options map[string]string) int {
			fmt.Println(Version)
			return 0
		})

	app := cli.New("Command line utility to work with GeoTIFF files").
		WithCommand(tagsCmd).
		WithCommand(versionCmd)

	os.Exit(app.Run(os.Args, os.Stdout))
}
