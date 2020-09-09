// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build s390x,!go1.11 !amd64,!s390x,!ppc64le gccgo appengine nacl
=======
// +build s390x,!go1.11 !amd64,!s390x,!ppc64le gccgo purego
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375

package poly1305

func sum(out *[TagSize]byte, msg []byte, key *[32]byte) {
	h := newMAC(key)
	h.Write(msg)
	h.Sum(out)
}
