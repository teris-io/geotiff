/*
 * Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.
 */

package main

import (
	"fmt"
	"strconv"

	"github.com/google/tiff"
)

func tags(args []string, options map[string]string) int {

	var err error
	singlepage := false
	indent := "\t"

	page := -1
	if pagestr := options["page"]; pagestr != "" {
		if page, err = strconv.Atoi(pagestr); err != nil {
			return done(err)
		}
		singlepage = true
		indent = ""
	}

	_, verbose := options["verbose"]

	err = do(args[0], func(tf tiff.TIFF) error {
		ifds := tf.IFDs()
		fpage := 0
		lpage := len(ifds) - 1
		if page >= 0 {
			if page > lpage {
				return fmt.Errorf("page %d does not exist", page)
			}
			fpage = page
			lpage = page
		}

		for page = fpage; page <= lpage; page++ {
			if !singlepage {
				fmt.Println(fmt.Sprintf("Page: %d", page))
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
				fmt.Println(fmt.Sprintf("%s%v=%v", indent, f.Tag().Name(), value))
			}
		}

		return nil
	})
	return done(err)
}
