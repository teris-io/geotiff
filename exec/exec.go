/*
 * Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.
 */

package exec

import (
	"os"

	"github.com/google/tiff"
	_ "github.com/google/tiff/geotiff"
)

func DoWithTiff(filename string, callback func(tf tiff.TIFF) error) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	r := tiff.NewReadAtReadSeeker(file)
	t, err := tiff.Parse(r, tiff.DefaultTagSpace, tiff.DefaultFieldTypeSpace)
	if err != nil {
		return err
	}
	return callback(t)
}
