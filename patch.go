package lupus

import (
	"reflect"
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
		0xFF, 0xe2,     // jmp rdx
	}
}

func getPage(p uintptr) []byte {
	return (*(*[0xFFFFFF]byte)(unsafe.Pointer(p & ^uintptr(syscall.Getpagesize()-1))))[:syscall.Getpagesize()]
}

func rawMemoryAccess(p uintptr, length int) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: p,
		Len:  length,
		Cap:  length,
	}))
}

func replace(orig uintptr, replacement []byte) {
	replacePtr := *(*uintptr)(unsafe.Pointer(&replacement))
	jumpData := assembleJump(replacePtr)
	// f := rawMemoryAccess(orig, len(jumpData))
	// original := make([]byte, len(f))
	// copy(original, f)

	copyToLocation(orig, jumpData)
	return
}

func copyToLocation(location uintptr, data []byte) {
	f := rawMemoryAccess(location, len(data))

	page := rawMemoryAccess(pageStart(location), syscall.Getpagesize())
	err := syscall.Mprotect(page, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC)
	if err != nil {
		panic(err)
	}
	copy(f, data[:])
}

func pageStart(ptr uintptr) uintptr {
	return ptr & ^(uintptr(syscall.Getpagesize() - 1))
}
