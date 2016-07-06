// Copyright (c) 2016, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

package sh

import (
	"bytes"
	"io/ioutil"

	"mvdan.cc/sh/syntax"
)

func Fuzz(data []byte) int {
	prog, err := syntax.NewParser(syntax.KeepComments).Parse(bytes.NewReader(data), "")
	if err != nil {
		return 0
	}
	syntax.Simplify(prog)
	syntax.NewPrinter().Print(ioutil.Discard, prog)
	return 1
}
