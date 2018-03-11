// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package tags

import "github.com/google/tiff"

// Extract collects tag name to value maps for all pages in the TIFF
func Extract(tf tiff.TIFF) (tags []map[string]interface{}) {
	for _, ifd := range tf.IFDs() {
		data := make(map[string]interface{})
		for _, f := range ifd.Fields() {
			data[f.Tag().Name()] = f.ParsedValue()
		}
		tags = append(tags, data)
	}
	return
}
