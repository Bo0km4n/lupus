package lupus

import (
	"log"
	"reflect"
)

func PatchFunction(patchObjPath string, funcName string, target interface{}) error {
	elfFile, _ := open(patchObjPath)
	newFuncBytes := getFuncBytes(elfFile, funcName)
	// ptr := *(*uintptr)(unsafe.Pointer(&newFuncBytes))

	// debug
	// mov
	// newFuncBytes := []byte{0x48, 0xC7, 0xC2, 0x01, 0x00, 0x00, 0x00}

	newFuncBytes, err := writeFuncVal(newFuncBytes)
	if err != nil {
		return err
	}
	patchValue(reflect.ValueOf(target), newFuncBytes)
	return nil
}

func patchValue(target reflect.Value, replaceValue []byte) {
	if target.Kind() != reflect.Func {
		log.Fatal("Target has to be a Func")
	}
	replace(target.Pointer(), replaceValue)
}
