// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imgd

import (
	"image"
	"image/png"
	"os"
	"path"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

// SavePng saves png figure
func SavePng(outdir, fnameKey string, img image.Image) {

	// create file
	furl := path.Join(outdir, fnameKey+".png")
	file, err := os.Create(furl)
	if err != nil {
		chk.Panic("cannot create file at")
	}
	defer file.Close()

	// encode png
	err = png.Encode(file, img)
	if err != nil {
		chk.Panic("cannot encode image")
	}
	io.Pf("file <%s> written\n", furl)
}
