package mmap

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"syscall"
	"unsafe"
)

// WriteFuncVal allocates func value as byte array to heap.
func WriteFuncVal(funcVal []byte) (uintptr, error) {
	f, err := writeFuncValueToTempFile(funcVal)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	data, err := syscall.Mmap(int(f.Fd()), 0, (len(funcVal)+0xfff)&^0xfff, syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_SHARED)
	if err != nil {
		return 0, err
	}

	ptr := *(*uintptr)(unsafe.Pointer(&data))
	return ptr, nil
}

func writeFuncValueToTempFile(funcVal []byte) (*os.File, error) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "tempfile-test-")
	if err != nil {
		fmt.Println("Error, can not create temp file.")
		log.Fatal(err)
	}

	fmt.Println("File : " + tempFile.Name())

	_, err = tempFile.Write(funcVal)
	return tempFile, err
}
