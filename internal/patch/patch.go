package patch

import (
	"syscall"
	"unsafe"
)

func assembleJump(to uintptr) []byte {
	// pp.Println("assembleJump:", to)
	// pp.Println("assembleJump to >> 0~56", byte(to), byte(to>>8), byte(to>>16), byte(to>>24), byte(to>>32), byte(to>>40), byte(to>>48), byte(to>>56))
	return []byte{
		0x48, 0xBA,
		byte(to),
		byte(to >> 8),
		byte(to >> 16),
		byte(to >> 24),
		byte(to >> 32),
		byte(to >> 40),
		byte(to >> 48),
		byte(to >> 56), // movabs rdx,to
		0xFF, 0x22,     // jmp QWORD PTR [rdx]
	}
}

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

func rawMemoryAccess(b uintptr) []byte {
	return (*(*[0xFF]byte)(unsafe.Pointer(b)))[:]
}

func Replace(orig func() int, replacement []byte) {
	replacePtr := *(*uintptr)(unsafe.Pointer(&replacement))
	bytes := assembleJump(replacePtr)
	functionLocation := **(**uintptr)(unsafe.Pointer(&orig))
	window := rawMemoryAccess(functionLocation)

	page := getPage(functionLocation)
	syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	copy(window, bytes)
}
