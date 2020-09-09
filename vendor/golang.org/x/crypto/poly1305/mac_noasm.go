// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !amd64,!ppc64le gccgo appengine
=======
// +build !amd64,!ppc64le gccgo purego
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375

package poly1305

type mac struct{ macGeneric }

func newMAC(key *[32]byte) mac { return mac{newMACGeneric(key)} }
