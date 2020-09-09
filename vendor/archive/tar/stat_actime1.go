// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build linux dragonfly openbsd solaris
=======
// +build aix linux dragonfly openbsd solaris
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375

package tar

import (
	"syscall"
	"time"
)

func statAtime(st *syscall.Stat_t) time.Time {
	return time.Unix(st.Atim.Unix())
}

func statCtime(st *syscall.Stat_t) time.Time {
	return time.Unix(st.Ctim.Unix())
}
