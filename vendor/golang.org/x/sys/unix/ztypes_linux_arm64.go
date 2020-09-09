// cgo -godefs -- -Wall -Werror -static -I/tmp/include -fsigned-char linux/types.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build arm64,linux

package unix

const (
	SizeofPtr  = 0x8
	SizeofLong = 0x8
)

type (
	_C_long int64
)

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int64
}

type Timex struct {
	Modes     uint32
	Offset    int64
	Freq      int64
	Maxerror  int64
	Esterror  int64
	Status    int32
	Constant  int64
	Precision int64
	Tolerance int64
	Time      Timeval
	Tick      int64
	Ppsfreq   int64
	Jitter    int64
	Shift     int32
	Stabil    int64
	Jitcnt    int64
	Calcnt    int64
	Errcnt    int64
	Stbcnt    int64
	Tai       int32
	_         [44]byte
}

type Time_t int64

type Tms struct {
	Utime  int64
	Stime  int64
	Cutime int64
	Cstime int64
}

type Utimbuf struct {
	Actime  int64
	Modtime int64
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type Stat_t struct {
	Dev     uint64
	Ino     uint64
	Mode    uint32
	Nlink   uint32
	Uid     uint32
	Gid     uint32
	Rdev    uint64
	_       uint64
	Size    int64
	Blksize int32
	_       int32
	Blocks  int64
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	_       [2]int32
}

type Dirent struct {
	Ino    uint64
	Off    int64
	Reclen uint16
	Type   uint8
	Name   [256]int8
	_      [5]byte
}

type Flock_t struct {
	Type   int16
	Whence int16
	Start  int64
	Len    int64
	Pid    int32
	_      [4]byte
}

const (
	FADV_DONTNEED = 0x4
	FADV_NOREUSE  = 0x5
)

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [96]int8
}

type Iovec struct {
	Base *byte
	Len  uint64
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     uint64
	Control    *byte
	Controllen uint64
	Flags      int32
	_          [4]byte
}

type Cmsghdr struct {
	Len   uint64
	Level int32
	Type  int32
}

const (
	SizeofIovec   = 0x10
	SizeofMsghdr  = 0x38
	SizeofCmsghdr = 0x10
)

const (
<<<<<<< HEAD
	NDA_UNSPEC              = 0x0
	NDA_DST                 = 0x1
	NDA_LLADDR              = 0x2
	NDA_CACHEINFO           = 0x3
	NDA_PROBES              = 0x4
	NDA_VLAN                = 0x5
	NDA_PORT                = 0x6
	NDA_VNI                 = 0x7
	NDA_IFINDEX             = 0x8
	NDA_MASTER              = 0x9
	NDA_LINK_NETNSID        = 0xa
	NDA_SRC_VNI             = 0xb
	NTF_USE                 = 0x1
	NTF_SELF                = 0x2
	NTF_MASTER              = 0x4
	NTF_PROXY               = 0x8
	NTF_EXT_LEARNED         = 0x10
	NTF_OFFLOADED           = 0x20
	NTF_ROUTER              = 0x80
	NUD_INCOMPLETE          = 0x1
	NUD_REACHABLE           = 0x2
	NUD_STALE               = 0x4
	NUD_DELAY               = 0x8
	NUD_PROBE               = 0x10
	NUD_FAILED              = 0x20
	NUD_NOARP               = 0x40
	NUD_PERMANENT           = 0x80
	NUD_NONE                = 0x0
	IFA_UNSPEC              = 0x0
	IFA_ADDRESS             = 0x1
	IFA_LOCAL               = 0x2
	IFA_LABEL               = 0x3
	IFA_BROADCAST           = 0x4
	IFA_ANYCAST             = 0x5
	IFA_CACHEINFO           = 0x6
	IFA_MULTICAST           = 0x7
	IFA_FLAGS               = 0x8
	IFA_RT_PRIORITY         = 0x9
	IFA_TARGET_NETNSID      = 0xa
	IFLA_UNSPEC             = 0x0
	IFLA_ADDRESS            = 0x1
	IFLA_BROADCAST          = 0x2
	IFLA_IFNAME             = 0x3
	IFLA_MTU                = 0x4
	IFLA_LINK               = 0x5
	IFLA_QDISC              = 0x6
	IFLA_STATS              = 0x7
	IFLA_COST               = 0x8
	IFLA_PRIORITY           = 0x9
	IFLA_MASTER             = 0xa
	IFLA_WIRELESS           = 0xb
	IFLA_PROTINFO           = 0xc
	IFLA_TXQLEN             = 0xd
	IFLA_MAP                = 0xe
	IFLA_WEIGHT             = 0xf
	IFLA_OPERSTATE          = 0x10
	IFLA_LINKMODE           = 0x11
	IFLA_LINKINFO           = 0x12
	IFLA_NET_NS_PID         = 0x13
	IFLA_IFALIAS            = 0x14
	IFLA_NUM_VF             = 0x15
	IFLA_VFINFO_LIST        = 0x16
	IFLA_STATS64            = 0x17
	IFLA_VF_PORTS           = 0x18
	IFLA_PORT_SELF          = 0x19
	IFLA_AF_SPEC            = 0x1a
	IFLA_GROUP              = 0x1b
	IFLA_NET_NS_FD          = 0x1c
	IFLA_EXT_MASK           = 0x1d
	IFLA_PROMISCUITY        = 0x1e
	IFLA_NUM_TX_QUEUES      = 0x1f
	IFLA_NUM_RX_QUEUES      = 0x20
	IFLA_CARRIER            = 0x21
	IFLA_PHYS_PORT_ID       = 0x22
	IFLA_CARRIER_CHANGES    = 0x23
	IFLA_PHYS_SWITCH_ID     = 0x24
	IFLA_LINK_NETNSID       = 0x25
	IFLA_PHYS_PORT_NAME     = 0x26
	IFLA_PROTO_DOWN         = 0x27
	IFLA_GSO_MAX_SEGS       = 0x28
	IFLA_GSO_MAX_SIZE       = 0x29
	IFLA_PAD                = 0x2a
	IFLA_XDP                = 0x2b
	IFLA_EVENT              = 0x2c
	IFLA_NEW_NETNSID        = 0x2d
	IFLA_IF_NETNSID         = 0x2e
	IFLA_TARGET_NETNSID     = 0x2e
	IFLA_CARRIER_UP_COUNT   = 0x2f
	IFLA_CARRIER_DOWN_COUNT = 0x30
	IFLA_NEW_IFINDEX        = 0x31
	IFLA_MIN_MTU            = 0x32
	IFLA_MAX_MTU            = 0x33
	IFLA_MAX                = 0x33
	IFLA_INFO_KIND          = 0x1
	IFLA_INFO_DATA          = 0x2
	IFLA_INFO_XSTATS        = 0x3
	IFLA_INFO_SLAVE_KIND    = 0x4
	IFLA_INFO_SLAVE_DATA    = 0x5
	RT_SCOPE_UNIVERSE       = 0x0
	RT_SCOPE_SITE           = 0xc8
	RT_SCOPE_LINK           = 0xfd
	RT_SCOPE_HOST           = 0xfe
	RT_SCOPE_NOWHERE        = 0xff
	RT_TABLE_UNSPEC         = 0x0
	RT_TABLE_COMPAT         = 0xfc
	RT_TABLE_DEFAULT        = 0xfd
	RT_TABLE_MAIN           = 0xfe
	RT_TABLE_LOCAL          = 0xff
	RT_TABLE_MAX            = 0xffffffff
	RTA_UNSPEC              = 0x0
	RTA_DST                 = 0x1
	RTA_SRC                 = 0x2
	RTA_IIF                 = 0x3
	RTA_OIF                 = 0x4
	RTA_GATEWAY             = 0x5
	RTA_PRIORITY            = 0x6
	RTA_PREFSRC             = 0x7
	RTA_METRICS             = 0x8
	RTA_MULTIPATH           = 0x9
	RTA_FLOW                = 0xb
	RTA_CACHEINFO           = 0xc
	RTA_TABLE               = 0xf
	RTA_MARK                = 0x10
	RTA_MFC_STATS           = 0x11
	RTA_VIA                 = 0x12
	RTA_NEWDST              = 0x13
	RTA_PREF                = 0x14
	RTA_ENCAP_TYPE          = 0x15
	RTA_ENCAP               = 0x16
	RTA_EXPIRES             = 0x17
	RTA_PAD                 = 0x18
	RTA_UID                 = 0x19
	RTA_TTL_PROPAGATE       = 0x1a
	RTA_IP_PROTO            = 0x1b
	RTA_SPORT               = 0x1c
	RTA_DPORT               = 0x1d
	RTN_UNSPEC              = 0x0
	RTN_UNICAST             = 0x1
	RTN_LOCAL               = 0x2
	RTN_BROADCAST           = 0x3
	RTN_ANYCAST             = 0x4
	RTN_MULTICAST           = 0x5
	RTN_BLACKHOLE           = 0x6
	RTN_UNREACHABLE         = 0x7
	RTN_PROHIBIT            = 0x8
	RTN_THROW               = 0x9
	RTN_NAT                 = 0xa
	RTN_XRESOLVE            = 0xb
	RTNLGRP_NONE            = 0x0
	RTNLGRP_LINK            = 0x1
	RTNLGRP_NOTIFY          = 0x2
	RTNLGRP_NEIGH           = 0x3
	RTNLGRP_TC              = 0x4
	RTNLGRP_IPV4_IFADDR     = 0x5
	RTNLGRP_IPV4_MROUTE     = 0x6
	RTNLGRP_IPV4_ROUTE      = 0x7
	RTNLGRP_IPV4_RULE       = 0x8
	RTNLGRP_IPV6_IFADDR     = 0x9
	RTNLGRP_IPV6_MROUTE     = 0xa
	RTNLGRP_IPV6_ROUTE      = 0xb
	RTNLGRP_IPV6_IFINFO     = 0xc
	RTNLGRP_IPV6_PREFIX     = 0x12
	RTNLGRP_IPV6_RULE       = 0x13
	RTNLGRP_ND_USEROPT      = 0x14
	SizeofNlMsghdr          = 0x10
	SizeofNlMsgerr          = 0x14
	SizeofRtGenmsg          = 0x1
	SizeofNlAttr            = 0x4
	SizeofRtAttr            = 0x4
	SizeofIfInfomsg         = 0x10
	SizeofIfAddrmsg         = 0x8
	SizeofRtMsg             = 0xc
	SizeofRtNexthop         = 0x8
	SizeofNdUseroptmsg      = 0x10
	SizeofNdMsg             = 0xc
)

type NlMsghdr struct {
	Len   uint32
	Type  uint16
	Flags uint16
	Seq   uint32
	Pid   uint32
}

type NlMsgerr struct {
	Error int32
	Msg   NlMsghdr
}

type RtGenmsg struct {
	Family uint8
}

type NlAttr struct {
	Len  uint16
	Type uint16
}

type RtAttr struct {
	Len  uint16
	Type uint16
}

type IfInfomsg struct {
	Family uint8
	_      uint8
	Type   uint16
	Index  int32
	Flags  uint32
	Change uint32
}

type IfAddrmsg struct {
	Family    uint8
	Prefixlen uint8
	Flags     uint8
	Scope     uint8
	Index     uint32
}

type RtMsg struct {
	Family   uint8
	Dst_len  uint8
	Src_len  uint8
	Tos      uint8
	Table    uint8
	Protocol uint8
	Scope    uint8
	Type     uint8
	Flags    uint32
}

type RtNexthop struct {
	Len     uint16
	Flags   uint8
	Hops    uint8
	Ifindex int32
}

type NdUseroptmsg struct {
	Family    uint8
	Pad1      uint8
	Opts_len  uint16
	Ifindex   int32
	Icmp_type uint8
	Icmp_code uint8
	Pad2      uint16
	Pad3      uint32
}

type NdMsg struct {
	Family  uint8
	Pad1    uint8
	Pad2    uint16
	Ifindex int32
	State   uint16
	Flags   uint8
	Type    uint8
}

const (
	SizeofSockFilter = 0x8
	SizeofSockFprog  = 0x10
=======
	SizeofSockFprog = 0x10
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
)

type PtraceRegs struct {
	Regs   [31]uint64
	Sp     uint64
	Pc     uint64
	Pstate uint64
}

type FdSet struct {
	Bits [16]int64
}

type Sysinfo_t struct {
	Uptime    int64
	Loads     [3]uint64
	Totalram  uint64
	Freeram   uint64
	Sharedram uint64
	Bufferram uint64
	Totalswap uint64
	Freeswap  uint64
	Procs     uint16
	Pad       uint16
	Totalhigh uint64
	Freehigh  uint64
	Unit      uint32
	_         [0]int8
	_         [4]byte
}

type Ustat_t struct {
	Tfree  int32
	Tinode uint64
	Fname  [6]int8
	Fpack  [6]int8
	_      [4]byte
}

type EpollEvent struct {
	Events uint32
	PadFd  int32
	Fd     int32
	Pad    int32
}

const (
	POLLRDHUP = 0x2000
)

type Sigset_t struct {
	Val [16]uint64
}

const _C__NSIG = 0x41
<<<<<<< HEAD

type SignalfdSiginfo struct {
	Signo     uint32
	Errno     int32
	Code      int32
	Pid       uint32
	Uid       uint32
	Fd        int32
	Tid       uint32
	Band      uint32
	Overrun   uint32
	Trapno    uint32
	Status    int32
	Int       int32
	Ptr       uint64
	Utime     uint64
	Stime     uint64
	Addr      uint64
	Addr_lsb  uint16
	_         uint16
	Syscall   int32
	Call_addr uint64
	Arch      uint32
	_         [28]uint8
}

const PERF_IOC_FLAG_GROUP = 0x1
=======
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Line   uint8
	Cc     [19]uint8
	Ispeed uint32
	Ospeed uint32
}

type Taskstats struct {
	Version                   uint16
	Ac_exitcode               uint32
	Ac_flag                   uint8
	Ac_nice                   uint8
	Cpu_count                 uint64
	Cpu_delay_total           uint64
	Blkio_count               uint64
	Blkio_delay_total         uint64
	Swapin_count              uint64
	Swapin_delay_total        uint64
	Cpu_run_real_total        uint64
	Cpu_run_virtual_total     uint64
	Ac_comm                   [32]int8
	Ac_sched                  uint8
	Ac_pad                    [3]uint8
	_                         [4]byte
	Ac_uid                    uint32
	Ac_gid                    uint32
	Ac_pid                    uint32
	Ac_ppid                   uint32
	Ac_btime                  uint32
	Ac_etime                  uint64
	Ac_utime                  uint64
	Ac_stime                  uint64
	Ac_minflt                 uint64
	Ac_majflt                 uint64
	Coremem                   uint64
	Virtmem                   uint64
	Hiwater_rss               uint64
	Hiwater_vm                uint64
	Read_char                 uint64
	Write_char                uint64
	Read_syscalls             uint64
	Write_syscalls            uint64
	Read_bytes                uint64
	Write_bytes               uint64
	Cancelled_write_bytes     uint64
	Nvcsw                     uint64
	Nivcsw                    uint64
	Ac_utimescaled            uint64
	Ac_stimescaled            uint64
	Cpu_scaled_run_real_total uint64
	Freepages_count           uint64
	Freepages_delay_total     uint64
	Thrashing_count           uint64
	Thrashing_delay_total     uint64
}

type cpuMask uint64

const (
	_NCPUBITS = 0x40
)

const (
	CBitFieldMaskBit0  = 0x1
	CBitFieldMaskBit1  = 0x2
	CBitFieldMaskBit2  = 0x4
	CBitFieldMaskBit3  = 0x8
	CBitFieldMaskBit4  = 0x10
	CBitFieldMaskBit5  = 0x20
	CBitFieldMaskBit6  = 0x40
	CBitFieldMaskBit7  = 0x80
	CBitFieldMaskBit8  = 0x100
	CBitFieldMaskBit9  = 0x200
	CBitFieldMaskBit10 = 0x400
	CBitFieldMaskBit11 = 0x800
	CBitFieldMaskBit12 = 0x1000
	CBitFieldMaskBit13 = 0x2000
	CBitFieldMaskBit14 = 0x4000
	CBitFieldMaskBit15 = 0x8000
	CBitFieldMaskBit16 = 0x10000
	CBitFieldMaskBit17 = 0x20000
	CBitFieldMaskBit18 = 0x40000
	CBitFieldMaskBit19 = 0x80000
	CBitFieldMaskBit20 = 0x100000
	CBitFieldMaskBit21 = 0x200000
	CBitFieldMaskBit22 = 0x400000
	CBitFieldMaskBit23 = 0x800000
	CBitFieldMaskBit24 = 0x1000000
	CBitFieldMaskBit25 = 0x2000000
	CBitFieldMaskBit26 = 0x4000000
	CBitFieldMaskBit27 = 0x8000000
	CBitFieldMaskBit28 = 0x10000000
	CBitFieldMaskBit29 = 0x20000000
	CBitFieldMaskBit30 = 0x40000000
	CBitFieldMaskBit31 = 0x80000000
	CBitFieldMaskBit32 = 0x100000000
	CBitFieldMaskBit33 = 0x200000000
	CBitFieldMaskBit34 = 0x400000000
	CBitFieldMaskBit35 = 0x800000000
	CBitFieldMaskBit36 = 0x1000000000
	CBitFieldMaskBit37 = 0x2000000000
	CBitFieldMaskBit38 = 0x4000000000
	CBitFieldMaskBit39 = 0x8000000000
	CBitFieldMaskBit40 = 0x10000000000
	CBitFieldMaskBit41 = 0x20000000000
	CBitFieldMaskBit42 = 0x40000000000
	CBitFieldMaskBit43 = 0x80000000000
	CBitFieldMaskBit44 = 0x100000000000
	CBitFieldMaskBit45 = 0x200000000000
	CBitFieldMaskBit46 = 0x400000000000
	CBitFieldMaskBit47 = 0x800000000000
	CBitFieldMaskBit48 = 0x1000000000000
	CBitFieldMaskBit49 = 0x2000000000000
	CBitFieldMaskBit50 = 0x4000000000000
	CBitFieldMaskBit51 = 0x8000000000000
	CBitFieldMaskBit52 = 0x10000000000000
	CBitFieldMaskBit53 = 0x20000000000000
	CBitFieldMaskBit54 = 0x40000000000000
	CBitFieldMaskBit55 = 0x80000000000000
	CBitFieldMaskBit56 = 0x100000000000000
	CBitFieldMaskBit57 = 0x200000000000000
	CBitFieldMaskBit58 = 0x400000000000000
	CBitFieldMaskBit59 = 0x800000000000000
	CBitFieldMaskBit60 = 0x1000000000000000
	CBitFieldMaskBit61 = 0x2000000000000000
	CBitFieldMaskBit62 = 0x4000000000000000
	CBitFieldMaskBit63 = 0x8000000000000000
)

type SockaddrStorage struct {
	Family uint16
	_      [118]int8
	_      uint64
}

type HDGeometry struct {
	Heads     uint8
	Sectors   uint8
	Cylinders uint16
	Start     uint64
}

type Statfs_t struct {
	Type    int64
	Bsize   int64
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Fsid    Fsid
	Namelen int64
	Frsize  int64
	Flags   int64
	Spare   [4]int64
}

type TpacketHdr struct {
	Status  uint64
	Len     uint32
	Snaplen uint32
	Mac     uint16
	Net     uint16
	Sec     uint32
	Usec    uint32
	_       [4]byte
}

const (
	SizeofTpacketHdr = 0x20
)

type RTCPLLInfo struct {
	Ctrl    int32
	Value   int32
	Max     int32
	Min     int32
	Posmult int32
	Negmult int32
	Clock   int64
}

type BlkpgPartition struct {
	Start   int64
	Length  int64
	Pno     int32
	Devname [64]uint8
	Volname [64]uint8
	_       [4]byte
}

const (
	BLKPG = 0x1269
)

<<<<<<< HEAD
type TpacketBDTS struct {
	Sec  uint32
	Usec uint32
}

type TpacketHdrV1 struct {
	Block_status        uint32
	Num_pkts            uint32
	Offset_to_first_pkt uint32
	Blk_len             uint32
	Seq_num             uint64
	Ts_first_pkt        TpacketBDTS
	Ts_last_pkt         TpacketBDTS
}

type TpacketReq struct {
	Block_size uint32
	Block_nr   uint32
	Frame_size uint32
	Frame_nr   uint32
=======
type XDPUmemReg struct {
	Addr     uint64
	Len      uint64
	Size     uint32
	Headroom uint32
	Flags    uint32
	_        [4]byte
>>>>>>> 0906c7fae9345571e51d6103eb90774d5f408375
}

type CryptoUserAlg struct {
	Name        [64]int8
	Driver_name [64]int8
	Module_name [64]int8
	Type        uint32
	Mask        uint32
	Refcnt      uint32
	Flags       uint32
}

type CryptoStatAEAD struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatAKCipher struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Verify_cnt   uint64
	Sign_cnt     uint64
	Err_cnt      uint64
}

type CryptoStatCipher struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatCompress struct {
	Type            [64]int8
	Compress_cnt    uint64
	Compress_tlen   uint64
	Decompress_cnt  uint64
	Decompress_tlen uint64
	Err_cnt         uint64
}

type CryptoStatHash struct {
	Type      [64]int8
	Hash_cnt  uint64
	Hash_tlen uint64
	Err_cnt   uint64
}

type CryptoStatKPP struct {
	Type                      [64]int8
	Setsecret_cnt             uint64
	Generate_public_key_cnt   uint64
	Compute_shared_secret_cnt uint64
	Err_cnt                   uint64
}

type CryptoStatRNG struct {
	Type          [64]int8
	Generate_cnt  uint64
	Generate_tlen uint64
	Seed_cnt      uint64
	Err_cnt       uint64
}

type CryptoStatLarval struct {
	Type [64]int8
}

type CryptoReportLarval struct {
	Type [64]int8
}

type CryptoReportHash struct {
	Type       [64]int8
	Blocksize  uint32
	Digestsize uint32
}

type CryptoReportCipher struct {
	Type        [64]int8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
}

type CryptoReportBlkCipher struct {
	Type        [64]int8
	Geniv       [64]int8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
	Ivsize      uint32
}

type CryptoReportAEAD struct {
	Type        [64]int8
	Geniv       [64]int8
	Blocksize   uint32
	Maxauthsize uint32
	Ivsize      uint32
}

type CryptoReportComp struct {
	Type [64]int8
}

type CryptoReportRNG struct {
	Type     [64]int8
	Seedsize uint32
}

type CryptoReportAKCipher struct {
	Type [64]int8
}

type CryptoReportKPP struct {
	Type [64]int8
}

type CryptoReportAcomp struct {
	Type [64]int8
}

type LoopInfo struct {
	Number           int32
	Device           uint32
	Inode            uint64
	Rdevice          uint32
	Offset           int32
	Encrypt_type     int32
	Encrypt_key_size int32
	Flags            int32
	Name             [64]int8
	Encrypt_key      [32]uint8
	Init             [2]uint64
	Reserved         [4]int8
	_                [4]byte
}

type TIPCSubscr struct {
	Seq     TIPCServiceRange
	Timeout uint32
	Filter  uint32
	Handle  [8]int8
}

type TIPCSIOCLNReq struct {
	Peer     uint32
	Id       uint32
	Linkname [68]int8
}

type TIPCSIOCNodeIDReq struct {
	Peer uint32
	Id   [16]int8
}

const (
	CRYPTO_MSG_BASE      = 0x10
	CRYPTO_MSG_NEWALG    = 0x10
	CRYPTO_MSG_DELALG    = 0x11
	CRYPTO_MSG_UPDATEALG = 0x12
	CRYPTO_MSG_GETALG    = 0x13
	CRYPTO_MSG_DELRNG    = 0x14
	CRYPTO_MSG_GETSTAT   = 0x15
)

const (
	CRYPTOCFGA_UNSPEC           = 0x0
	CRYPTOCFGA_PRIORITY_VAL     = 0x1
	CRYPTOCFGA_REPORT_LARVAL    = 0x2
	CRYPTOCFGA_REPORT_HASH      = 0x3
	CRYPTOCFGA_REPORT_BLKCIPHER = 0x4
	CRYPTOCFGA_REPORT_AEAD      = 0x5
	CRYPTOCFGA_REPORT_COMPRESS  = 0x6
	CRYPTOCFGA_REPORT_RNG       = 0x7
	CRYPTOCFGA_REPORT_CIPHER    = 0x8
	CRYPTOCFGA_REPORT_AKCIPHER  = 0x9
	CRYPTOCFGA_REPORT_KPP       = 0xa
	CRYPTOCFGA_REPORT_ACOMP     = 0xb
	CRYPTOCFGA_STAT_LARVAL      = 0xc
	CRYPTOCFGA_STAT_HASH        = 0xd
	CRYPTOCFGA_STAT_BLKCIPHER   = 0xe
	CRYPTOCFGA_STAT_AEAD        = 0xf
	CRYPTOCFGA_STAT_COMPRESS    = 0x10
	CRYPTOCFGA_STAT_RNG         = 0x11
	CRYPTOCFGA_STAT_CIPHER      = 0x12
	CRYPTOCFGA_STAT_AKCIPHER    = 0x13
	CRYPTOCFGA_STAT_KPP         = 0x14
	CRYPTOCFGA_STAT_ACOMP       = 0x15
)

type CryptoUserAlg struct {
	Name        [64]int8
	Driver_name [64]int8
	Module_name [64]int8
	Type        uint32
	Mask        uint32
	Refcnt      uint32
	Flags       uint32
}

type CryptoStatAEAD struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatAKCipher struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Verify_cnt   uint64
	Sign_cnt     uint64
	Err_cnt      uint64
}

type CryptoStatCipher struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatCompress struct {
	Type            [64]int8
	Compress_cnt    uint64
	Compress_tlen   uint64
	Decompress_cnt  uint64
	Decompress_tlen uint64
	Err_cnt         uint64
}

type CryptoStatHash struct {
	Type      [64]int8
	Hash_cnt  uint64
	Hash_tlen uint64
	Err_cnt   uint64
}

type CryptoStatKPP struct {
	Type                      [64]int8
	Setsecret_cnt             uint64
	Generate_public_key_cnt   uint64
	Compute_shared_secret_cnt uint64
	Err_cnt                   uint64
}

type CryptoStatRNG struct {
	Type          [64]int8
	Generate_cnt  uint64
	Generate_tlen uint64
	Seed_cnt      uint64
	Err_cnt       uint64
}

type CryptoStatLarval struct {
	Type [64]int8
}

type CryptoReportLarval struct {
	Type [64]int8
}

type CryptoReportHash struct {
	Type       [64]int8
	Blocksize  uint32
	Digestsize uint32
}

type CryptoReportCipher struct {
	Type        [64]int8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
}

type CryptoReportBlkCipher struct {
	Type        [64]int8
	Geniv       [64]int8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
	Ivsize      uint32
}

type CryptoReportAEAD struct {
	Type        [64]int8
	Geniv       [64]int8
	Blocksize   uint32
	Maxauthsize uint32
	Ivsize      uint32
}

type CryptoReportComp struct {
	Type [64]int8
}

type CryptoReportRNG struct {
	Type     [64]int8
	Seedsize uint32
}

type CryptoReportAKCipher struct {
	Type [64]int8
}

type CryptoReportKPP struct {
	Type [64]int8
}

type CryptoReportAcomp struct {
	Type [64]int8
}

const (
	BPF_REG_0                           = 0x0
	BPF_REG_1                           = 0x1
	BPF_REG_2                           = 0x2
	BPF_REG_3                           = 0x3
	BPF_REG_4                           = 0x4
	BPF_REG_5                           = 0x5
	BPF_REG_6                           = 0x6
	BPF_REG_7                           = 0x7
	BPF_REG_8                           = 0x8
	BPF_REG_9                           = 0x9
	BPF_REG_10                          = 0xa
	BPF_MAP_CREATE                      = 0x0
	BPF_MAP_LOOKUP_ELEM                 = 0x1
	BPF_MAP_UPDATE_ELEM                 = 0x2
	BPF_MAP_DELETE_ELEM                 = 0x3
	BPF_MAP_GET_NEXT_KEY                = 0x4
	BPF_PROG_LOAD                       = 0x5
	BPF_OBJ_PIN                         = 0x6
	BPF_OBJ_GET                         = 0x7
	BPF_PROG_ATTACH                     = 0x8
	BPF_PROG_DETACH                     = 0x9
	BPF_PROG_TEST_RUN                   = 0xa
	BPF_PROG_GET_NEXT_ID                = 0xb
	BPF_MAP_GET_NEXT_ID                 = 0xc
	BPF_PROG_GET_FD_BY_ID               = 0xd
	BPF_MAP_GET_FD_BY_ID                = 0xe
	BPF_OBJ_GET_INFO_BY_FD              = 0xf
	BPF_PROG_QUERY                      = 0x10
	BPF_RAW_TRACEPOINT_OPEN             = 0x11
	BPF_BTF_LOAD                        = 0x12
	BPF_BTF_GET_FD_BY_ID                = 0x13
	BPF_TASK_FD_QUERY                   = 0x14
	BPF_MAP_LOOKUP_AND_DELETE_ELEM      = 0x15
	BPF_MAP_TYPE_UNSPEC                 = 0x0
	BPF_MAP_TYPE_HASH                   = 0x1
	BPF_MAP_TYPE_ARRAY                  = 0x2
	BPF_MAP_TYPE_PROG_ARRAY             = 0x3
	BPF_MAP_TYPE_PERF_EVENT_ARRAY       = 0x4
	BPF_MAP_TYPE_PERCPU_HASH            = 0x5
	BPF_MAP_TYPE_PERCPU_ARRAY           = 0x6
	BPF_MAP_TYPE_STACK_TRACE            = 0x7
	BPF_MAP_TYPE_CGROUP_ARRAY           = 0x8
	BPF_MAP_TYPE_LRU_HASH               = 0x9
	BPF_MAP_TYPE_LRU_PERCPU_HASH        = 0xa
	BPF_MAP_TYPE_LPM_TRIE               = 0xb
	BPF_MAP_TYPE_ARRAY_OF_MAPS          = 0xc
	BPF_MAP_TYPE_HASH_OF_MAPS           = 0xd
	BPF_MAP_TYPE_DEVMAP                 = 0xe
	BPF_MAP_TYPE_SOCKMAP                = 0xf
	BPF_MAP_TYPE_CPUMAP                 = 0x10
	BPF_MAP_TYPE_XSKMAP                 = 0x11
	BPF_MAP_TYPE_SOCKHASH               = 0x12
	BPF_MAP_TYPE_CGROUP_STORAGE         = 0x13
	BPF_MAP_TYPE_REUSEPORT_SOCKARRAY    = 0x14
	BPF_MAP_TYPE_PERCPU_CGROUP_STORAGE  = 0x15
	BPF_MAP_TYPE_QUEUE                  = 0x16
	BPF_MAP_TYPE_STACK                  = 0x17
	BPF_PROG_TYPE_UNSPEC                = 0x0
	BPF_PROG_TYPE_SOCKET_FILTER         = 0x1
	BPF_PROG_TYPE_KPROBE                = 0x2
	BPF_PROG_TYPE_SCHED_CLS             = 0x3
	BPF_PROG_TYPE_SCHED_ACT             = 0x4
	BPF_PROG_TYPE_TRACEPOINT            = 0x5
	BPF_PROG_TYPE_XDP                   = 0x6
	BPF_PROG_TYPE_PERF_EVENT            = 0x7
	BPF_PROG_TYPE_CGROUP_SKB            = 0x8
	BPF_PROG_TYPE_CGROUP_SOCK           = 0x9
	BPF_PROG_TYPE_LWT_IN                = 0xa
	BPF_PROG_TYPE_LWT_OUT               = 0xb
	BPF_PROG_TYPE_LWT_XMIT              = 0xc
	BPF_PROG_TYPE_SOCK_OPS              = 0xd
	BPF_PROG_TYPE_SK_SKB                = 0xe
	BPF_PROG_TYPE_CGROUP_DEVICE         = 0xf
	BPF_PROG_TYPE_SK_MSG                = 0x10
	BPF_PROG_TYPE_RAW_TRACEPOINT        = 0x11
	BPF_PROG_TYPE_CGROUP_SOCK_ADDR      = 0x12
	BPF_PROG_TYPE_LWT_SEG6LOCAL         = 0x13
	BPF_PROG_TYPE_LIRC_MODE2            = 0x14
	BPF_PROG_TYPE_SK_REUSEPORT          = 0x15
	BPF_PROG_TYPE_FLOW_DISSECTOR        = 0x16
	BPF_CGROUP_INET_INGRESS             = 0x0
	BPF_CGROUP_INET_EGRESS              = 0x1
	BPF_CGROUP_INET_SOCK_CREATE         = 0x2
	BPF_CGROUP_SOCK_OPS                 = 0x3
	BPF_SK_SKB_STREAM_PARSER            = 0x4
	BPF_SK_SKB_STREAM_VERDICT           = 0x5
	BPF_CGROUP_DEVICE                   = 0x6
	BPF_SK_MSG_VERDICT                  = 0x7
	BPF_CGROUP_INET4_BIND               = 0x8
	BPF_CGROUP_INET6_BIND               = 0x9
	BPF_CGROUP_INET4_CONNECT            = 0xa
	BPF_CGROUP_INET6_CONNECT            = 0xb
	BPF_CGROUP_INET4_POST_BIND          = 0xc
	BPF_CGROUP_INET6_POST_BIND          = 0xd
	BPF_CGROUP_UDP4_SENDMSG             = 0xe
	BPF_CGROUP_UDP6_SENDMSG             = 0xf
	BPF_LIRC_MODE2                      = 0x10
	BPF_FLOW_DISSECTOR                  = 0x11
	BPF_STACK_BUILD_ID_EMPTY            = 0x0
	BPF_STACK_BUILD_ID_VALID            = 0x1
	BPF_STACK_BUILD_ID_IP               = 0x2
	BPF_ADJ_ROOM_NET                    = 0x0
	BPF_HDR_START_MAC                   = 0x0
	BPF_HDR_START_NET                   = 0x1
	BPF_LWT_ENCAP_SEG6                  = 0x0
	BPF_LWT_ENCAP_SEG6_INLINE           = 0x1
	BPF_OK                              = 0x0
	BPF_DROP                            = 0x2
	BPF_REDIRECT                        = 0x7
	BPF_SOCK_OPS_VOID                   = 0x0
	BPF_SOCK_OPS_TIMEOUT_INIT           = 0x1
	BPF_SOCK_OPS_RWND_INIT              = 0x2
	BPF_SOCK_OPS_TCP_CONNECT_CB         = 0x3
	BPF_SOCK_OPS_ACTIVE_ESTABLISHED_CB  = 0x4
	BPF_SOCK_OPS_PASSIVE_ESTABLISHED_CB = 0x5
	BPF_SOCK_OPS_NEEDS_ECN              = 0x6
	BPF_SOCK_OPS_BASE_RTT               = 0x7
	BPF_SOCK_OPS_RTO_CB                 = 0x8
	BPF_SOCK_OPS_RETRANS_CB             = 0x9
	BPF_SOCK_OPS_STATE_CB               = 0xa
	BPF_SOCK_OPS_TCP_LISTEN_CB          = 0xb
	BPF_TCP_ESTABLISHED                 = 0x1
	BPF_TCP_SYN_SENT                    = 0x2
	BPF_TCP_SYN_RECV                    = 0x3
	BPF_TCP_FIN_WAIT1                   = 0x4
	BPF_TCP_FIN_WAIT2                   = 0x5
	BPF_TCP_TIME_WAIT                   = 0x6
	BPF_TCP_CLOSE                       = 0x7
	BPF_TCP_CLOSE_WAIT                  = 0x8
	BPF_TCP_LAST_ACK                    = 0x9
	BPF_TCP_LISTEN                      = 0xa
	BPF_TCP_CLOSING                     = 0xb
	BPF_TCP_NEW_SYN_RECV                = 0xc
	BPF_TCP_MAX_STATES                  = 0xd
	BPF_FIB_LKUP_RET_SUCCESS            = 0x0
	BPF_FIB_LKUP_RET_BLACKHOLE          = 0x1
	BPF_FIB_LKUP_RET_UNREACHABLE        = 0x2
	BPF_FIB_LKUP_RET_PROHIBIT           = 0x3
	BPF_FIB_LKUP_RET_NOT_FWDED          = 0x4
	BPF_FIB_LKUP_RET_FWD_DISABLED       = 0x5
	BPF_FIB_LKUP_RET_UNSUPP_LWT         = 0x6
	BPF_FIB_LKUP_RET_NO_NEIGH           = 0x7
	BPF_FIB_LKUP_RET_FRAG_NEEDED        = 0x8
	BPF_FD_TYPE_RAW_TRACEPOINT          = 0x0
	BPF_FD_TYPE_TRACEPOINT              = 0x1
	BPF_FD_TYPE_KPROBE                  = 0x2
	BPF_FD_TYPE_KRETPROBE               = 0x3
	BPF_FD_TYPE_UPROBE                  = 0x4
	BPF_FD_TYPE_URETPROBE               = 0x5
)
