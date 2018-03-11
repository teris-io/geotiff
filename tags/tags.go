/*
 * Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.
 */

package tags

import (
	"fmt"
	"io"

	"github.com/google/tiff"
	"github.com/teris-io/geotiff/exec"
)

func PrintTags(writer io.Writer, filename string, page int, verbose bool) error {

	return exec.DoWithTiff(filename, func(tf tiff.TIFF) error {
		singlepage := false
		indent := "\t"

		ifds := tf.IFDs()
		fpage := 0
		lpage := len(ifds) - 1
		if page >= 0 {
			if page > lpage {
				return fmt.Errorf("page %d does not exist", page)
			}
			fpage = page
			lpage = page
			singlepage = true
			indent = ""
		}

		for page = fpage; page <= lpage; page++ {
			if !singlepage {
				fmt.Fprintln(writer, fmt.Sprintf("Page: %d", page))
			}
			ifd := ifds[page]
			for _, f := range ifd.Fields() {
				value := f.ParsedValue()
				if arr, ok := value.([]interface{}); ok && len(arr) > 10 && !verbose {
					var repr []interface{}
					repr = append(repr, arr[:10]...)
					repr = append(repr, fmt.Sprintf("...(%d)", len(arr)))
					value = repr
				}
				fmt.Fprintln(writer, fmt.Sprintf("%s%v=%v", indent, f.Tag().Name(), value))
			}
		}
		return nil
	})
}
