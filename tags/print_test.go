/*
 * Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.
 */

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

func TestFprint(t *testing.T) {
	var tgs []map[string]interface{}
	exec.DoWithTiff(filename, func(tf tiff.TIFF) error {
		tgs = tags.Extract(tf)
		return nil
	})

	w := &strwriter{}
	tags.Fprint(w, tgs, false, false)

	if len(w.str) != 617 || !strings.HasPrefix(w.str, "BitsPerSample=8") || !strings.HasSuffix(w.str, "YResolution=1/1") {
		t.Fatalf("expected 617 chars, found: %s", w.str)
	}
}
