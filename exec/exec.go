// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package exec

import (
	"os"

	"github.com/google/tiff"
	// provides geotiff extensions for the standard tiff library
	_ "github.com/google/tiff/geotiff"
)

// DoWithTiff executes a function with TIFF file handler for filename
func DoWithTiff(filename string, callback func(tiff.TIFF) error) (err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return err
	}
	defer file.Close()

	r := tiff.NewReadAtReadSeeker(file)
	var tf tiff.TIFF
	if tf, err = tiff.Parse(r, tiff.DefaultTagSpace, tiff.DefaultFieldTypeSpace); err != nil {
		return err
	}
	return callback(tf)
}
