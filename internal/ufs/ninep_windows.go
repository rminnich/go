// Copyright 2012 The Ninep Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// This code is imported from the old ninep repo,
// with some changes.

// +build windows

package ufs

import (
	"os"

	"harvey-os.org/pkg/ninep/protocol"
)

func fileInfoToQID(d os.FileInfo) protocol.QID {
	var qid protocol.QID

	qid.Path = uint64(d.ModTime().UnixNano())
	qid.Version = uint32(d.ModTime().UnixNano() / 1000000)
	qid.Type = dirToQIDType(d)

	return qid
}
