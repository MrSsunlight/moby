// +build riscv64

package bbolt

// maxMapSize represents the largest mmap size supported by Bolt.
const maxMapSize = 0xFFFFFFFFFFFF // 256TB

// maxAllocSize is the size used when creating array pointers.
const maxAllocSize = 0x7FFFFFFF
<<<<<<< HEAD

// Are unaligned load/stores broken on this arch?
var brokenUnaligned = true
=======
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
