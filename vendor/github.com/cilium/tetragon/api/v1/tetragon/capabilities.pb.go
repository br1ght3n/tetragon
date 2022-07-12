// Copyright 2020 Authors of Hubble
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: tetragon/capabilities.proto

package tetragon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CapabilitiesType int32

const (
	// In a system with the [_POSIX_CHOWN_RESTRICTED] option defined, this
	//overrides the restriction of changing file ownership and group
	//ownership.
	CapabilitiesType_CAP_CHOWN CapabilitiesType = 0
	// Override all DAC access, including ACL execute access if
	//[_POSIX_ACL] is defined. Excluding DAC access covered by
	//CAP_LINUX_IMMUTABLE.
	CapabilitiesType_DAC_OVERRIDE CapabilitiesType = 1
	// Overrides all DAC restrictions regarding read and search on files
	//and directories, including ACL restrictions if [_POSIX_ACL] is
	//defined. Excluding DAC access covered by "$1"_LINUX_IMMUTABLE.
	CapabilitiesType_CAP_DAC_READ_SEARCH CapabilitiesType = 2
	// Overrides all restrictions about allowed operations on files, where
	//file owner ID must be equal to the user ID, except where CAP_FSETID
	//is applicable. It doesn't override MAC and DAC restrictions.
	CapabilitiesType_CAP_FOWNER CapabilitiesType = 3
	// Overrides the following restrictions that the effective user ID
	//shall match the file owner ID when setting the S_ISUID and S_ISGID
	//bits on that file; that the effective group ID (or one of the
	//supplementary group IDs) shall match the file owner ID when setting
	//the S_ISGID bit on that file; that the S_ISUID and S_ISGID bits are
	//cleared on successful return from chown(2) (not implemented).
	CapabilitiesType_CAP_FSETID CapabilitiesType = 4
	// Overrides the restriction that the real or effective user ID of a
	//process sending a signal must match the real or effective user ID
	//of the process receiving the signal.
	CapabilitiesType_CAP_KILL CapabilitiesType = 5
	// Allows forged gids on socket credentials passing.
	CapabilitiesType_CAP_SETGID CapabilitiesType = 6
	// Allows forged pids on socket credentials passing.
	CapabilitiesType_CAP_SETUID CapabilitiesType = 7
	// Without VFS support for capabilities:
	//   Transfer any capability in your permitted set to any pid,
	//   remove any capability in your permitted set from any pid
	// With VFS support for capabilities (neither of above, but)
	//   Add any capability from current's capability bounding set
	//       to the current process' inheritable set
	//   Allow taking bits out of capability bounding set
	//   Allow modification of the securebits for a process
	CapabilitiesType_CAP_SETPCAP CapabilitiesType = 8
	// Allow modification of S_IMMUTABLE and S_APPEND file attributes
	CapabilitiesType_CAP_LINUX_IMMUTABLE CapabilitiesType = 9
	// Allows binding to ATM VCIs below 32
	CapabilitiesType_CAP_NET_BIND_SERVICE CapabilitiesType = 10
	// Allow broadcasting, listen to multicast
	CapabilitiesType_CAP_NET_BROADCAST CapabilitiesType = 11
	// Allow activation of ATM control sockets
	CapabilitiesType_CAP_NET_ADMIN CapabilitiesType = 12
	// Allow binding to any address for transparent proxying (also via NET_ADMIN)
	CapabilitiesType_CAP_NET_RAW CapabilitiesType = 13
	// Allow mlock and mlockall (which doesn't really have anything to do
	//with IPC)
	CapabilitiesType_CAP_IPC_LOCK CapabilitiesType = 14
	// Override IPC ownership checks
	CapabilitiesType_CAP_IPC_OWNER CapabilitiesType = 15
	// Insert and remove kernel modules - modify kernel without limit
	CapabilitiesType_CAP_SYS_MODULE CapabilitiesType = 16
	// Allow sending USB messages to any device via /dev/bus/usb
	CapabilitiesType_CAP_SYS_RAWIO CapabilitiesType = 17
	// Allow use of chroot()
	CapabilitiesType_CAP_SYS_CHROOT CapabilitiesType = 18
	// Allow ptrace() of any process
	CapabilitiesType_CAP_SYS_PTRACE CapabilitiesType = 19
	// Allow configuration of process accounting
	CapabilitiesType_CAP_SYS_PACCT CapabilitiesType = 20
	// Allow everything under CAP_BPF and CAP_PERFMON for backward compatibility
	CapabilitiesType_CAP_SYS_ADMIN CapabilitiesType = 21
	// Allow use of reboot()
	CapabilitiesType_CAP_SYS_BOOT CapabilitiesType = 22
	// Allow setting cpu affinity on other processes
	CapabilitiesType_CAP_SYS_NICE CapabilitiesType = 23
	// Control memory reclaim behavior
	CapabilitiesType_CAP_SYS_RESOURCE CapabilitiesType = 24
	// Allow setting the real-time clock
	CapabilitiesType_CAP_SYS_TIME CapabilitiesType = 25
	// Allow vhangup() of tty
	CapabilitiesType_CAP_SYS_TTY_CONFIG CapabilitiesType = 26
	// Allow the privileged aspects of mknod()
	CapabilitiesType_CAP_MKNOD CapabilitiesType = 27
	// Allow taking of leases on files
	CapabilitiesType_CAP_LEASE CapabilitiesType = 28
	// Allow writing the audit log via unicast netlink socket
	CapabilitiesType_CAP_AUDIT_WRITE CapabilitiesType = 29
	// Allow configuration of audit via unicast netlink socket
	CapabilitiesType_CAP_AUDIT_CONTROL CapabilitiesType = 30
	// Set or remove capabilities on files
	CapabilitiesType_CAP_SETFCAP CapabilitiesType = 31
	// Override MAC access.
	//The base kernel enforces no MAC policy.
	//An LSM may enforce a MAC policy, and if it does and it chooses
	//to implement capability based overrides of that policy, this is
	//the capability it should use to do so.
	CapabilitiesType_CAP_MAC_OVERRIDE CapabilitiesType = 32
	// Allow MAC configuration or state changes.
	//The base kernel requires no MAC configuration.
	//An LSM may enforce a MAC policy, and if it does and it chooses
	//to implement capability based checks on modifications to that
	//policy or the data required to maintain it, this is the
	//capability it should use to do so.
	CapabilitiesType_CAP_MAC_ADMIN CapabilitiesType = 33
	// Allow configuring the kernel's syslog (printk behaviour)
	CapabilitiesType_CAP_SYSLOG CapabilitiesType = 34
	// Allow triggering something that will wake the system
	CapabilitiesType_CAP_WAKE_ALARM CapabilitiesType = 35
	// Allow preventing system suspends
	CapabilitiesType_CAP_BLOCK_SUSPEND CapabilitiesType = 36
	// Allow reading the audit log via multicast netlink socket
	CapabilitiesType_CAP_AUDIT_READ CapabilitiesType = 37
	//
	// Allow system performance and observability privileged operations
	// using perf_events, i915_perf and other kernel subsystems
	CapabilitiesType_CAP_PERFMON CapabilitiesType = 38
	//
	// CAP_BPF allows the following BPF operations:
	// - Creating all types of BPF maps
	// - Advanced verifier features
	//   - Indirect variable access
	//   - Bounded loops
	//   - BPF to BPF function calls
	//   - Scalar precision tracking
	//   - Larger complexity limits
	//   - Dead code elimination
	//   - And potentially other features
	// - Loading BPF Type Format (BTF) data
	// - Retrieve xlated and JITed code of BPF programs
	// - Use bpf_spin_lock() helper
	// CAP_PERFMON relaxes the verifier checks further:
	// - BPF progs can use of pointer-to-integer conversions
	// - speculation attack hardening measures are bypassed
	// - bpf_probe_read to read arbitrary kernel memory is allowed
	// - bpf_trace_printk to print kernel memory is allowed
	// CAP_SYS_ADMIN is required to use bpf_probe_write_user.
	// CAP_SYS_ADMIN is required to iterate system wide loaded
	// programs, maps, links, BTFs and convert their IDs to file descriptors.
	// CAP_PERFMON and CAP_BPF are required to load tracing programs.
	// CAP_NET_ADMIN and CAP_BPF are required to load networking programs.
	CapabilitiesType_CAP_BPF CapabilitiesType = 39
	// Allow writing to ns_last_pid
	CapabilitiesType_CAP_CHECKPOINT_RESTORE CapabilitiesType = 40
)

// Enum value maps for CapabilitiesType.
var (
	CapabilitiesType_name = map[int32]string{
		0:  "CAP_CHOWN",
		1:  "DAC_OVERRIDE",
		2:  "CAP_DAC_READ_SEARCH",
		3:  "CAP_FOWNER",
		4:  "CAP_FSETID",
		5:  "CAP_KILL",
		6:  "CAP_SETGID",
		7:  "CAP_SETUID",
		8:  "CAP_SETPCAP",
		9:  "CAP_LINUX_IMMUTABLE",
		10: "CAP_NET_BIND_SERVICE",
		11: "CAP_NET_BROADCAST",
		12: "CAP_NET_ADMIN",
		13: "CAP_NET_RAW",
		14: "CAP_IPC_LOCK",
		15: "CAP_IPC_OWNER",
		16: "CAP_SYS_MODULE",
		17: "CAP_SYS_RAWIO",
		18: "CAP_SYS_CHROOT",
		19: "CAP_SYS_PTRACE",
		20: "CAP_SYS_PACCT",
		21: "CAP_SYS_ADMIN",
		22: "CAP_SYS_BOOT",
		23: "CAP_SYS_NICE",
		24: "CAP_SYS_RESOURCE",
		25: "CAP_SYS_TIME",
		26: "CAP_SYS_TTY_CONFIG",
		27: "CAP_MKNOD",
		28: "CAP_LEASE",
		29: "CAP_AUDIT_WRITE",
		30: "CAP_AUDIT_CONTROL",
		31: "CAP_SETFCAP",
		32: "CAP_MAC_OVERRIDE",
		33: "CAP_MAC_ADMIN",
		34: "CAP_SYSLOG",
		35: "CAP_WAKE_ALARM",
		36: "CAP_BLOCK_SUSPEND",
		37: "CAP_AUDIT_READ",
		38: "CAP_PERFMON",
		39: "CAP_BPF",
		40: "CAP_CHECKPOINT_RESTORE",
	}
	CapabilitiesType_value = map[string]int32{
		"CAP_CHOWN":              0,
		"DAC_OVERRIDE":           1,
		"CAP_DAC_READ_SEARCH":    2,
		"CAP_FOWNER":             3,
		"CAP_FSETID":             4,
		"CAP_KILL":               5,
		"CAP_SETGID":             6,
		"CAP_SETUID":             7,
		"CAP_SETPCAP":            8,
		"CAP_LINUX_IMMUTABLE":    9,
		"CAP_NET_BIND_SERVICE":   10,
		"CAP_NET_BROADCAST":      11,
		"CAP_NET_ADMIN":          12,
		"CAP_NET_RAW":            13,
		"CAP_IPC_LOCK":           14,
		"CAP_IPC_OWNER":          15,
		"CAP_SYS_MODULE":         16,
		"CAP_SYS_RAWIO":          17,
		"CAP_SYS_CHROOT":         18,
		"CAP_SYS_PTRACE":         19,
		"CAP_SYS_PACCT":          20,
		"CAP_SYS_ADMIN":          21,
		"CAP_SYS_BOOT":           22,
		"CAP_SYS_NICE":           23,
		"CAP_SYS_RESOURCE":       24,
		"CAP_SYS_TIME":           25,
		"CAP_SYS_TTY_CONFIG":     26,
		"CAP_MKNOD":              27,
		"CAP_LEASE":              28,
		"CAP_AUDIT_WRITE":        29,
		"CAP_AUDIT_CONTROL":      30,
		"CAP_SETFCAP":            31,
		"CAP_MAC_OVERRIDE":       32,
		"CAP_MAC_ADMIN":          33,
		"CAP_SYSLOG":             34,
		"CAP_WAKE_ALARM":         35,
		"CAP_BLOCK_SUSPEND":      36,
		"CAP_AUDIT_READ":         37,
		"CAP_PERFMON":            38,
		"CAP_BPF":                39,
		"CAP_CHECKPOINT_RESTORE": 40,
	}
)

func (x CapabilitiesType) Enum() *CapabilitiesType {
	p := new(CapabilitiesType)
	*p = x
	return p
}

func (x CapabilitiesType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CapabilitiesType) Descriptor() protoreflect.EnumDescriptor {
	return file_tetragon_capabilities_proto_enumTypes[0].Descriptor()
}

func (CapabilitiesType) Type() protoreflect.EnumType {
	return &file_tetragon_capabilities_proto_enumTypes[0]
}

func (x CapabilitiesType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CapabilitiesType.Descriptor instead.
func (CapabilitiesType) EnumDescriptor() ([]byte, []int) {
	return file_tetragon_capabilities_proto_rawDescGZIP(), []int{0}
}

var File_tetragon_capabilities_proto protoreflect.FileDescriptor

var file_tetragon_capabilities_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74,
	0x65, 0x74, 0x72, 0x61, 0x67, 0x6f, 0x6e, 0x2a, 0xa4, 0x06, 0x0a, 0x10, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x41, 0x50, 0x5f, 0x43, 0x48, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x44,
	0x41, 0x43, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x52, 0x49, 0x44, 0x45, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x43, 0x41, 0x50, 0x5f, 0x44, 0x41, 0x43, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x5f, 0x53, 0x45,
	0x41, 0x52, 0x43, 0x48, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x50, 0x5f, 0x46, 0x4f,
	0x57, 0x4e, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x50, 0x5f, 0x46, 0x53,
	0x45, 0x54, 0x49, 0x44, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x50, 0x5f, 0x4b, 0x49,
	0x4c, 0x4c, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x45, 0x54, 0x47,
	0x49, 0x44, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x45, 0x54, 0x55,
	0x49, 0x44, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x45, 0x54, 0x50,
	0x43, 0x41, 0x50, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x50, 0x5f, 0x4c, 0x49, 0x4e,
	0x55, 0x58, 0x5f, 0x49, 0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x09, 0x12, 0x18,
	0x0a, 0x14, 0x43, 0x41, 0x50, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x42, 0x49, 0x4e, 0x44, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x50, 0x5f,
	0x4e, 0x45, 0x54, 0x5f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x10, 0x0b, 0x12,
	0x11, 0x0a, 0x0d, 0x43, 0x41, 0x50, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x0c, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x50, 0x5f, 0x4e, 0x45, 0x54, 0x5f, 0x52, 0x41,
	0x57, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x50, 0x5f, 0x49, 0x50, 0x43, 0x5f, 0x4c,
	0x4f, 0x43, 0x4b, 0x10, 0x0e, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x50, 0x5f, 0x49, 0x50, 0x43,
	0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0x0f, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x5f,
	0x53, 0x59, 0x53, 0x5f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d,
	0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x52, 0x41, 0x57, 0x49, 0x4f, 0x10, 0x11, 0x12,
	0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x43, 0x48, 0x52, 0x4f, 0x4f,
	0x54, 0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x50,
	0x54, 0x52, 0x41, 0x43, 0x45, 0x10, 0x13, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x50, 0x5f, 0x53,
	0x59, 0x53, 0x5f, 0x50, 0x41, 0x43, 0x43, 0x54, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41,
	0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x15, 0x12, 0x10, 0x0a,
	0x0c, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x42, 0x4f, 0x4f, 0x54, 0x10, 0x16, 0x12,
	0x10, 0x0a, 0x0c, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x4e, 0x49, 0x43, 0x45, 0x10,
	0x17, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53, 0x5f, 0x52, 0x45, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x18, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x41, 0x50, 0x5f, 0x53,
	0x59, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x19, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x50,
	0x5f, 0x53, 0x59, 0x53, 0x5f, 0x54, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10,
	0x1a, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x50, 0x5f, 0x4d, 0x4b, 0x4e, 0x4f, 0x44, 0x10, 0x1b,
	0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x50, 0x5f, 0x4c, 0x45, 0x41, 0x53, 0x45, 0x10, 0x1c, 0x12,
	0x13, 0x0a, 0x0f, 0x43, 0x41, 0x50, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x1d, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x50, 0x5f, 0x41, 0x55, 0x44, 0x49,
	0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c, 0x10, 0x1e, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x41, 0x50, 0x5f, 0x53, 0x45, 0x54, 0x46, 0x43, 0x41, 0x50, 0x10, 0x1f, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x52, 0x49, 0x44, 0x45,
	0x10, 0x20, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x50, 0x5f, 0x4d, 0x41, 0x43, 0x5f, 0x41, 0x44,
	0x4d, 0x49, 0x4e, 0x10, 0x21, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x41, 0x50, 0x5f, 0x53, 0x59, 0x53,
	0x4c, 0x4f, 0x47, 0x10, 0x22, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x5f, 0x57, 0x41, 0x4b,
	0x45, 0x5f, 0x41, 0x4c, 0x41, 0x52, 0x4d, 0x10, 0x23, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x50,
	0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x24,
	0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x50, 0x5f, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x25, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x41, 0x50, 0x5f, 0x50, 0x45, 0x52, 0x46,
	0x4d, 0x4f, 0x4e, 0x10, 0x26, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x50, 0x5f, 0x42, 0x50, 0x46,
	0x10, 0x27, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x50, 0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x50,
	0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x28, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tetragon_capabilities_proto_rawDescOnce sync.Once
	file_tetragon_capabilities_proto_rawDescData = file_tetragon_capabilities_proto_rawDesc
)

func file_tetragon_capabilities_proto_rawDescGZIP() []byte {
	file_tetragon_capabilities_proto_rawDescOnce.Do(func() {
		file_tetragon_capabilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_tetragon_capabilities_proto_rawDescData)
	})
	return file_tetragon_capabilities_proto_rawDescData
}

var file_tetragon_capabilities_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tetragon_capabilities_proto_goTypes = []interface{}{
	(CapabilitiesType)(0), // 0: tetragon.CapabilitiesType
}
var file_tetragon_capabilities_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tetragon_capabilities_proto_init() }
func file_tetragon_capabilities_proto_init() {
	if File_tetragon_capabilities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tetragon_capabilities_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tetragon_capabilities_proto_goTypes,
		DependencyIndexes: file_tetragon_capabilities_proto_depIdxs,
		EnumInfos:         file_tetragon_capabilities_proto_enumTypes,
	}.Build()
	File_tetragon_capabilities_proto = out.File
	file_tetragon_capabilities_proto_rawDesc = nil
	file_tetragon_capabilities_proto_goTypes = nil
	file_tetragon_capabilities_proto_depIdxs = nil
}
