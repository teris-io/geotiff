// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package tags_test

import (
	"strings"
	"testing"

	"github.com/google/tiff"
	"github.com/teris-io/geotiff/exec"
	"github.com/teris-io/geotiff/tags"
	"github.com/teris-io/geotiff/testdata"
)

func TestFprint_Ok(t *testing.T) {
	w := &testdata.StrWriter{}
	tags.Fprint(w, load(), false, false)
	if len(w.Str) != 618 || !strings.HasPrefix(w.Str, "BitsPerSample=8") || !strings.HasSuffix(w.Str, "YResolution=1/1\n") {
		t.Fatalf("expected 618 chars, found %d: %s", len(w.Str), w.Str)
	}
}

func TestFprintPretty_Ok(t *testing.T) {
	w := &testdata.StrWriter{}
	tags.Fprint(w, load(), true, false)
	if len(w.Str) != 647 || !strings.HasPrefix(w.Str, "Page: 0") || !strings.HasSuffix(w.Str, "YResolution=1/1\n") {
		t.Fatalf("expected 647 chars, found %d: %s", len(w.Str), w.Str)
	}
}

func load() (tgs []map[string]interface{}) {
	exec.DoWithTiff(filename, func(tf tiff.TIFF) error {
		tgs = tags.Extract(tf)
		return nil
	})
	return
}

func TestFprintForFile_Ok(t *testing.T) {
	w := &testdata.StrWriter{}
	if err := tags.FprintForFile(w, filename, 0, false); err != nil {
		t.Fatal(err.Error())
	}
	if len(w.Str) != 618 || !strings.HasPrefix(w.Str, "BitsPerSample=8") || !strings.HasSuffix(w.Str, "YResolution=1/1\n") {
		t.Fatalf("expected 618 chars, found %d: %s", len(w.Str), w.Str)
	}
}

func TestFprintForFilePretty_Ok(t *testing.T) {
	w := &testdata.StrWriter{}
	if err := tags.FprintForFile(w, filename, -1, false); err != nil {
		t.Fatal(err.Error())
	}
	if len(w.Str) != 647 || !strings.HasPrefix(w.Str, "Page: 0") || !strings.HasSuffix(w.Str, "YResolution=1/1\n") {
		t.Fatalf("expected 647 chars, found %d: %s", len(w.Str), w.Str)
	}
}

func TestFprintForFileNotFound_Error(t *testing.T) {
	expected := "open some file somewhere: no such file or directory"
	w := &testdata.StrWriter{}
	if err := tags.FprintForFile(w, "some file somewhere", -1, false); err == nil {
		t.Fatal("expected an error")
	} else if err.Error() != expected {
		t.Fatalf("expected '%s', found '%s'", expected, err.Error())
	}
}

func TestFprintForFilePageIndexOfBounds_Error(t *testing.T) {
	expected := "page 2 does not exist in the file"
	w := &testdata.StrWriter{}
	if err := tags.FprintForFile(w, filename, 2, false); err == nil {
		t.Fatal("expected an error")
	} else if err.Error() != expected {
		t.Fatalf("expected '%s', found '%s'", expected, err.Error())
	}
}
