// Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.

package testdata

import "fmt"

type StrWriter struct {
	Str string
}

func (w *StrWriter) Write(p []byte) (n int, err error) {
	w.Str = fmt.Sprintf("%s%s", w.Str, string(p))
	return len(p), nil
}
