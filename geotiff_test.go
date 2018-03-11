// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package main_test

import (
	"strings"
	"testing"

	geotiff "github.com/teris-io/geotiff"
	"github.com/teris-io/geotiff/testdata"
)

func TestRun_Version_Ok(t *testing.T) {
	w := &testdata.StrWriter{}
	if code := geotiff.Run([]string{"geotiff", "version"}, w); code != 0 {
		t.Fatalf("non zero exit code: %d", code)
	}
	if !strings.HasPrefix(w.Str, geotiff.Version) {
		t.Fatalf("expected %s, found %s", geotiff.Version, w.Str)
	}
}

func TestRun_Tags_AllOptions_Ok(t *testing.T) {
	w := &testdata.StrWriter{}
	if code := geotiff.Run([]string{"geotiff", "tags", "-p", "0", "-v", "testdata/ch_zh_topo_lzw.tiff"}, w); code != 0 {
		t.Fatalf("non zero exit code: %d", code)
	}
	if len(w.Str) != 5104 {
		t.Fatalf("expected 5104, found %d: %s", len(w.Str), w.Str)
	}
}

func TestRun_Tags_NoFileGiven_Error(t *testing.T) {
	w := &testdata.StrWriter{}
	if code := geotiff.Run([]string{"geotiff", "tags", "-p", "0"}, w); code == 0 {
		t.Fatal("zero exit code")
	}
	expected := "fatal: missing required argument filename"
	if !strings.HasPrefix(w.Str, expected) {
		t.Fatalf("expected: %s, found: %s", expected, w.Str)
	}
}

func TestRun_TagsFileNotFound_Error(t *testing.T) {
	w := &testdata.StrWriter{}
	if code := geotiff.Run([]string{"geotiff", "tags", "-p", "0", "some file somewhere"}, w); code == 0 {
		t.Fatal("zero exit code")
	}
	if len(w.Str) != 0 {
		t.Fatal(w.Str)
	}
}
