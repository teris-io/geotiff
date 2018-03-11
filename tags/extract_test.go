// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package tags_test

import (
	"strings"
	"testing"

	"github.com/google/tiff"
	"github.com/teris-io/geotiff/exec"
	"github.com/teris-io/geotiff/tags"
)

const filename = "../testdata/ch_zh_topo_lzw.tiff"

func TestExtract_Ok(t *testing.T) {
	var tgs []map[string]interface{}
	if err := exec.DoWithTiff(filename, func(tf tiff.TIFF) error {
		tgs = tags.Extract(tf)
		return nil
	}); err != nil {
		t.Fatal(err.Error())
	}
	if len(tgs) != 1 {
		t.Fatalf("expected 1, found %d", len(tgs))
	}
	if len(tgs[0]) != 21 {
		t.Fatalf("expected 21, found %d", len(tgs[0]))
	}
	if width, ok := tgs[0]["ImageWidth"].(uint32); !ok || width != 1105 {
		t.Fatalf("expected image width 1105, found %v: %v", ok, width)
	}
	if descr, ok := tgs[0]["ImageDescription"].(string); !ok || !strings.HasPrefix(descr, "Generated by MultiSpecIntel") {
		t.Fatalf("expected 'Generated by MultiSpecIntel...', found '%s'", descr)
	}
}
