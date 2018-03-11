// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package tags_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/tiff"
	"github.com/teris-io/geotiff/exec"
	"github.com/teris-io/geotiff/tags"
)

type strwriter struct {
	str string
}

func (w *strwriter) Write(p []byte) (n int, err error) {
	w.str = fmt.Sprintf("%s%s", w.str, string(p))
	return len(p), nil
}

func TestFprint_Ok(t *testing.T) {
	w := &strwriter{}
	tags.Fprint(w, load(), false, false)
	if len(w.str) != 617 || !strings.HasPrefix(w.str, "BitsPerSample=8") || !strings.HasSuffix(w.str, "YResolution=1/1") {
		t.Fatalf("expected 617 chars, found %d: %s", len(w.str), w.str)
	}
}

func TestFprintPretty_Ok(t *testing.T) {
	w := &strwriter{}
	tags.Fprint(w, load(), true, false)
	if len(w.str) != 646 || !strings.HasPrefix(w.str, "Page: 0") || !strings.HasSuffix(w.str, "YResolution=1/1") {
		t.Fatalf("expected 646 chars, found %d: %s", len(w.str), w.str)
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
	w := &strwriter{}
	if err := tags.FprintForFile(w, filename, 0, false); err != nil {
		t.Fatal(err.Error())
	}
	if len(w.str) != 617 || !strings.HasPrefix(w.str, "BitsPerSample=8") || !strings.HasSuffix(w.str, "YResolution=1/1") {
		t.Fatalf("expected 617 chars, found %d: %s", len(w.str), w.str)
	}
}

func TestFprintForFilePretty_Ok(t *testing.T) {
	w := &strwriter{}
	if err := tags.FprintForFile(w, filename, -1, false); err != nil {
		t.Fatal(err.Error())
	}
	if len(w.str) != 646 || !strings.HasPrefix(w.str, "Page: 0") || !strings.HasSuffix(w.str, "YResolution=1/1") {
		t.Fatalf("expected 617 chars, found %d: %s", len(w.str), w.str)
	}
}

func TestFprintForFileNotFound_Error(t *testing.T) {
	expected := "open some file somewhere: no such file or directory"
	w := &strwriter{}
	if err := tags.FprintForFile(w, "some file somewhere", -1, false); err == nil {
		t.Fatal("expected an error")
	} else if err.Error() != expected {
		t.Fatalf("expected '%s', found '%s'", expected, err.Error())
	}
}

func TestFprintForFilePageIndexOfBounds_Error(t *testing.T) {
	expected := "page 2 does not exist in the file"
	w := &strwriter{}
	if err := tags.FprintForFile(w, filename, 2, false); err == nil {
		t.Fatal("expected an error")
	} else if err.Error() != expected {
		t.Fatalf("expected '%s', found '%s'", expected, err.Error())
	}
}
