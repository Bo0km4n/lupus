package lupus

import (
	"log"
	"reflect"

	"github.com/Bo0km4n/lupus/internal/elf"
	"github.com/Bo0km4n/lupus/internal/mmap"
	"github.com/Bo0km4n/lupus/internal/patch"
)

func PatchFunction(patchObjPath string, funcName string, target interface{}) error {
	elfFile, _ := elf.Open(patchObjPath)
	newFuncBytes := elf.GetFuncBytes(elfFile, funcName)
	// ptr := *(*uintptr)(unsafe.Pointer(&newFuncBytes))

	// debug
	// mov
	// newFuncBytes := []byte{0x48, 0xC7, 0xC2, 0x01, 0x00, 0x00, 0x00}

	newFuncBytes, err := mmap.WriteFuncVal(newFuncBytes)
	if err != nil {
		return err
	}
	patchValue(reflect.ValueOf(target), newFuncBytes)
	return nil
}

func patchValue(target reflect.Value, replace []byte) {
	if target.Kind() != reflect.Func {
		log.Fatal("Target has to be a Func")
	}
	patch.Replace(target.Pointer(), replace)
}
