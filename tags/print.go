// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package tags

import (
	"fmt"
	"io"
	"sort"

	"github.com/google/tiff"
	"github.com/teris-io/geotiff/exec"
)

func Fprint(w io.Writer, tags []map[string]interface{}, prettify, verbose bool) {
	indent := ""
	if prettify {
		indent = "\t"
	}
	for page, data := range tags {
		if prettify {
			fmt.Fprintln(w, fmt.Sprintf("Page: %d", page))
		}
		var out []string
		for tag, value := range data {
			if arr, ok := value.([]interface{}); ok && len(arr) > 10 && !verbose {
				var repr []interface{}
				repr = append(repr, arr[:10]...)
				repr = append(repr, fmt.Sprintf("...(%d)", len(arr)))
				value = repr
			}
			out = append(out, fmt.Sprintf("%s%v=%v", indent, tag, value))
		}

		sort.Strings(out)
		for _, o := range out {
			fmt.Fprintln(w, o)
		}
	}
}

func FprintForFile(w io.Writer, filename string, page int, verbose bool) (err error) {
	var tags []map[string]interface{}
	err = exec.DoWithTiff(filename, func(tf tiff.TIFF) error {
		tags = Extract(tf)
		return nil
	})
	if err != nil {
		return err
	}
	prettify := true
	if page >= len(tags) {
		return fmt.Errorf("page %d does not exist in the file", page)
	} else if page >= 0 {
		tags = tags[page : page+1]
		prettify = false
	}
	Fprint(w, tags, prettify, verbose)
	return nil
}
