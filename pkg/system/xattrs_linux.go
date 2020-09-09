package system // import "github.com/docker/docker/pkg/system"

import "golang.org/x/sys/unix"

// Lgetxattr retrieves the value of the extended attribute identified by attr
// and associated with the given path in the file system.
// It will returns a nil slice and nil error if the xattr is not set.
func Lgetxattr(path string, attr string) ([]byte, error) {
	// Start with a 128 length byte array
	dest := make([]byte, 128)
	sz, errno := unix.Lgetxattr(path, attr, dest)

<<<<<<< HEAD
	switch {
	case errno == unix.ENODATA:
		return nil, nil
	case errno == unix.ERANGE:
		// 128 byte array might just not be good enough. A dummy buffer is used
		// to get the real size of the xattrs on disk
=======
	for errno == unix.ERANGE {
		// Buffer too small, use zero-sized buffer to get the actual size
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
		sz, errno = unix.Lgetxattr(path, attr, []byte{})
		if errno != nil {
			return nil, errno
		}
		dest = make([]byte, sz)
		sz, errno = unix.Lgetxattr(path, attr, dest)
<<<<<<< HEAD
		if errno != nil {
			return nil, errno
		}
=======
	}

	switch {
	case errno == unix.ENODATA:
		return nil, nil
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
	case errno != nil:
		return nil, errno
	}
	return dest[:sz], nil
}

// Lsetxattr sets the value of the extended attribute identified by attr
// and associated with the given path in the file system.
func Lsetxattr(path string, attr string, data []byte, flags int) error {
	return unix.Lsetxattr(path, attr, data, flags)
}
