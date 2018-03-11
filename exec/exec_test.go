// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package exec_test

import (
	"testing"

	"github.com/google/tiff"
	"github.com/teris-io/geotiff/exec"
	"github.com/teris-io/geotiff/tags"
)

const filename = "../testdata/ch_zh_topo_lzw.tiff"

func TestDoWithTiff_Ok(t *testing.T) {
	var tgs []map[string]interface{}
	if err := exec.DoWithTiff(filename, func(tf tiff.TIFF) error {
		tgs = tags.Extract(tf)
		return nil
	}); err != nil {
		t.Fatal(err.Error())
	}
	if len(tgs) != 1 || len(tgs[0]) != 21 {
		t.Fatal("inconsistencies reading LZW TIFF")
	}
}

func TestDoWithTiffFileNotFound_Error(t *testing.T) {
	expected := "open some file somewhere: no such file or directory"
	if err := exec.DoWithTiff("some file somewhere", func(tf tiff.TIFF) error {
		return nil
	}); err == nil {
		t.Fatal("expected an error")
	} else if err.Error() != expected {
		t.Fatalf("expected '%s', found '%s'", expected, err.Error())
	}
}

func TestDoWithTiffNonTiff_Error(t *testing.T) {
	expected := "tiff: invalid byte order \"//\""
	if err := exec.DoWithTiff("exec_test.go", func(tf tiff.TIFF) error {
		return nil
	}); err == nil {
		t.Fatal("expected an error")
	} else if err.Error() != expected {
		t.Fatalf("expected '%s', found '%s'", expected, err.Error())
	}
}
