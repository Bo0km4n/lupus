package assemble

import (
	"syscall"
	"unsafe"
)

func assembleJump(replace uintptr) []byte {
	return []byte{
		0x48, 0xC7, 0xC2,
		byte(replace >> 0),
		byte(replace >> 8),
		byte(replace >> 16),
		byte(replace >> 24), // MOV rdx, replace
		0xFF, 0x22,          // JMP rdx
	}
}

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

func rawMemoryAccess(b uintptr) []byte {
	return (*(*[0xFF]byte)(unsafe.Pointer(b)))[:]
}

func Replace(orig, replacement func() int) {
	// bytes := assembleJump(replacement)
	// functionLocation := **(**uintptr)(unsafe.Pointer(&orig))
	// window := rawMemoryAccess(functionLocation)

	// page := getPage(functionLocation)
	// syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)

	// pp.Println("sizeof func", unsafe.Sizeof(functionLocation))
	// fmt.Printf("%x\n", bytes)
	// copy(window, bytes)
}
