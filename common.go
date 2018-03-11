/*
 * Copyright (c) Oleg Sklyar & teris.io, 2018. All rights reserved.
 */

package main

import "fmt"

const (
	Version = "0.1.0"
)

func done(err error) int {
	if err != nil {
		fmt.Println(fmt.Sprintf("fatal: %s", err.Error()))
		return 1
	}
	return 0
}
